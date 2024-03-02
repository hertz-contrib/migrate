// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/logic"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/logic/gin"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/logs"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/types"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/logic/chi"
	nethttp "github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/logic/netHttp"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"
	"github.com/urfave/cli/v2"

	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	globalMap = make(map[string]interface{})
	webCtxSet = mapset.NewSet[string]()
}

var (
	globalArgs = &Args{}
	globalMap  map[string]any
	webCtxSet  mapset.Set[string]
	fset       *token.FileSet
	wg         sync.WaitGroup

	printerConfig = printer.Config{
		Mode:     printer.UseSpaces,
		Tabwidth: 4,
	}
)

type Args struct {
	TargetDir  string
	Filepath   string
	HzRepo     string
	IgnoreDirs []string
	Debug      bool
	UseGin     bool
	UseNetHTTP bool
	UseChi     bool
}

const ignoreDirsText = `
Fill in the folders to be ignored, separating the folders with ",".
Example:
    hertz_migrate -target-dir ./project -ignore-dirs=kitex_gen,hz_gen
`

func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "hertz_migrate"
	app.Usage = "A tool for migrating to hertz from other go web frameworks"
	app.Version = internal.Version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "hz-repo",
			Aliases:     []string{"r"},
			Value:       "github.com/cloudwego/hertz",
			Usage:       "Specify the url of the hertz repository you want to bring in.",
			Destination: &globalArgs.HzRepo,
		},
		&cli.StringFlag{
			Name:        "target-dir",
			Aliases:     []string{"d"},
			Usage:       "project directory you wants to migrate",
			Destination: &globalArgs.TargetDir,
		},
		&cli.StringSliceFlag{
			Name:    "ignore-dirs",
			Aliases: []string{"I"},
			Usage:   ignoreDirsText,
		},
		&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"D"},
			Destination: &globalArgs.Debug,
			Value:       false,
		},
		&cli.BoolFlag{
			Name:        "use-gin",
			Aliases:     []string{"g"},
			Usage:       "migrate to hertz with gin as the web framework",
			Destination: &globalArgs.UseGin,
		},
		&cli.BoolFlag{
			Name:        "use-net-http",
			Aliases:     []string{"n"},
			Usage:       "migrate to hertz with net/http as the web framework",
			Destination: &globalArgs.UseNetHTTP,
		},
		&cli.BoolFlag{
			Name:        "use-chi",
			Aliases:     []string{"c"},
			Usage:       "migrate to hertz with chi as the web framework",
			Destination: &globalArgs.UseNetHTTP,
		},
	}
	app.Action = Run
	return app
}

func Run(c *cli.Context) error {
	fset = token.NewFileSet()
	globalArgs.IgnoreDirs = c.StringSlice("ignore-dirs")
	if globalArgs.UseChi {
		globalArgs.UseNetHTTP = true
	}

	if globalArgs.Debug {
		logs.SetLevel(logs.LevelDebug)
	}

	if globalArgs.TargetDir != "" {
		gofiles, err := utils.CollectGoFiles(globalArgs.TargetDir, globalArgs.IgnoreDirs)
		if err != nil {
			return err
		}

		goModDirs, err := utils.SearchAllDirHasGoMod(globalArgs.TargetDir)
		if err != nil {
			return err
		}

		for _, dir := range goModDirs {
			wg.Add(1)
			dir := dir
			go func() {
				defer wg.Done()
				utils.RunGoGet(dir, globalArgs.HzRepo)
			}()
		}
		wg.Wait()

		for _, path := range gofiles {
			file, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				logs.Debugf("Parse file fail, error: %v", err)
				return internal.ErrParseFile
			}

			astutil.Apply(file, func(c *astutil.Cursor) bool {
				logic.GetHttpServerProps(c)
				if globalArgs.UseGin {
					gin.GetFuncNameHasGinCtx(c)
				}
				if globalArgs.UseNetHTTP {
					nethttp.FindHandlerFuncName(c, webCtxSet)
				}
				return true
			}, nil)
		}

		if err = processFiles(gofiles); err != nil {
			return err
		}

		for _, dir := range goModDirs {
			utils.RunGoImports(dir)
			utils.RunGoModTidy(dir)
		}
	}
	logs.Info("everything are ok!")
	return nil
}

func processFiles(gofiles []string) error {
	for _, path := range gofiles {
		var containsGin bool

		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			logs.Debugf("Parse file fail, error: %v", err)
			return internal.ErrParseFile
		}

		astutil.AddNamedImport(fset, file, "hzserver", globalArgs.HzRepo+"/pkg/app/server")
		astutil.AddNamedImport(fset, file, "hzapp", globalArgs.HzRepo+"/pkg/app")
		astutil.AddNamedImport(fset, file, "hzroute", globalArgs.HzRepo+"/pkg/route")
		astutil.AddNamedImport(fset, file, "hzerrors", globalArgs.HzRepo+"/pkg/common/errors")
		astutil.AddNamedImport(fset, file, "hzutils", globalArgs.HzRepo+"/pkg/common/utils")

		for _, importSpec := range file.Imports {
			importStr := importSpec.Path.Value

			if globalArgs.UseGin {
				if strings.Contains(importStr, `github.com/gin-gonic/gin`) {
					containsGin = true
				}

				if strings.Contains(importStr, `"github.com/gin-contrib/cors"`) {
					importSpec.Path.Value = `"github.com/hertz-contrib/cors"`
					containsGin = true
				}

				if strings.Contains(importStr, `github.com/swaggo/gin-swagger`) {
					importSpec.Path.Value = `"github.com/hertz-contrib/swagger"`
					containsGin = true
				}
			}
		}

		if !containsGin && globalArgs.UseGin {
			continue
		}

		astutil.Apply(file, func(c *astutil.Cursor) bool {
			switch node := c.Node().(type) {
			case *ast.StarExpr:
				if globalArgs.UseChi {
					if utils.CheckPtrPkgAndStructName(node, "chi", "Mux") {
						c.Replace(types.StarServerHertz)
					}
				}
				if globalArgs.UseGin {
					if sel, ok := node.X.(*ast.SelectorExpr); ok {
						if utils.CheckSelPkgAndStruct(sel, "gin", "Engine") {
							c.Replace(types.StarServerHertz)
						}

						if utils.CheckSelPkgAndStruct(sel, "gin", "RouterGroup") {
							c.Replace(types.StarRouteGroup)
						}
					}
				}
			case *ast.FieldList:
				if globalArgs.UseGin {
					gin.ReplaceGinCtx(node)
				}
			case *ast.SelectorExpr:
				if globalArgs.UseGin {
					if utils.CheckSelPkgAndStruct(node, "route", "IRoutes") {
						c.Replace(types.SelIRoutes)
					}
				}
			}
			if globalArgs.UseNetHTTP {
				nethttp.GetOptionsFromHttpServer(c, globalMap)
				nethttp.PackServerHertz(c, globalMap)
				nethttp.ReplaceNetHttpHandler(c)
			}
			return true
		}, nil)

		astutil.Apply(file, func(c *astutil.Cursor) bool {
			netHttpGroup(c, webCtxSet)
			switch node := c.Node().(type) {
			case *ast.SelectorExpr:
				if globalArgs.UseNetHTTP {
					if utils.CheckSelObj(node, "http", "ResponseWriter") {
						switch node.Sel.Name {
						case "WriteHeader":
							c.Replace(types.SelSetStatusCode)
						case "Write":
							c.Replace(types.SelWrite)
						case "Header":
							c.Replace(types.SelRespHeader)
						}
					}

					if node.Sel.Name == "HandleFunc" {
						node.Sel.Name = "Any"
					}
					nethttp.ReplaceRequestOp(node, c)
				}

				if globalArgs.UseGin {
					if utils.CheckSelPkgAndStruct(node, "gin", "HandlerFunc") {
						c.Replace(types.SelAppHandlerFunc)
					}

					if utils.CheckSelPkgAndStruct(node, "gin", "H") {
						node.X.(*ast.Ident).Name = "hzutils"
					}

					gin.ReplaceBinding(node, c)
					gin.ReplaceRequestOp(node, c)
					gin.ReplaceRespOp(node, c)
					gin.ReplaceErrorType(node)
				}
			case *ast.CallExpr:
				if globalArgs.UseChi {
					chi.PackChiRouterMethod(node)
				}
				if globalArgs.UseNetHTTP {
					nethttp.ReplaceHttpOp(node, c)
					nethttp.ReplaceReqOrRespOp(node, c)
					nethttp.ReplaceReqURLQuery(node)
					if utils.CheckCallPkgAndMethodName(node, "http", "NotFound") {
						c.Replace(types.CallNotFound)
					}
				}

				if globalArgs.UseGin {
					gin.ReplaceGinNew(node, c)
					gin.ReplaceGinRun(node)
					gin.ReplaceGinCtxOp(node, c)
					gin.ReplaceCallReqOrResp(node, c)
					gin.ReplaceStatisFS(node)
				}
			}
			if globalArgs.UseGin {
				gin.ReplaceCtxParamList(c)
			}
			if globalArgs.UseChi {
				chi.PackChiNewRouter(c, globalMap)
			}
			return true
		}, nil)

		var buf bytes.Buffer

		if err = printerConfig.Fprint(&buf, fset, file); err != nil {
			logs.Debugf("Fprint fail, error: %v", err)
			return internal.ErrSaveChanges
		}

		if err := os.WriteFile(path, buf.Bytes(), os.ModePerm); err == nil {
			log.Println("File updated:", path)
		}
	}
	return nil
}

func netHttpGroup(c *astutil.Cursor, funcSet mapset.Set[string]) {
	if globalArgs.UseNetHTTP {
		nethttp.PackFprintf(c)
		nethttp.ReplaceReqHeader(c)
		nethttp.ReplaceReqHeaderOperation(c)
		nethttp.ReplaceRespWrite(c)
		nethttp.ReplaceReqFormGet(c)
		nethttp.ReplaceReqFormValue(c)
		nethttp.ReplaceReqMultipartForm(c)
		nethttp.PackType2AppHandlerFunc(c)
		nethttp.ReplaceReqMultipartFormOperation(c, globalMap)
		nethttp.ReplaceFuncBodyHttpHandlerParam(c, funcSet)
	}
}
