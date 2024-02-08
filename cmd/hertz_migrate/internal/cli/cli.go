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
	"path/filepath"
	"sync"

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
	funcSet = mapset.NewSet[string]()
}

var (
	globalArgs = &Args{}
	globalMap  map[string]any
)

type Args struct {
	TargetDir  string
	Filepath   string
	HzRepo     string
	IgnoreDirs []string
	Debug      bool
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
			Usage:       "Specify the url of the hertz repository you want to bring in.",
			DefaultText: "github.com/cloudwego/hertz",
			Destination: &globalArgs.HzRepo,
		},
		&cli.StringFlag{
			Name:        "target-dir",
			Usage:       "project directory you wants to migrate",
			Destination: &globalArgs.TargetDir,
		},
		&cli.StringSliceFlag{
			Name:  "ignore-dirs",
			Usage: ignoreDirsText,
		},
	}
	app.Action = Run
	return app
}

func Run(c *cli.Context) error {
	globalArgs.IgnoreDirs = c.StringSlice("ignore-dirs")
	if globalArgs.TargetDir != "" {
		gofiles, err := utils.CollectGoFiles(globalArgs.TargetDir, globalArgs.IgnoreDirs)
		for _, f := range gofiles {
			log.Println(f)
		}
		if err != nil {
			log.Fatal("Error collecting go files:", err)
		}
		goModDirs = utils.SearchAllDirHasGoMod(globalArgs.TargetDir)
		for _, dir := range goModDirs {
			wg.Add(1)
			dir := dir
			go func() {
				defer wg.Done()
				utils.RunGoGet(dir, globalArgs.HzRepo)
			}()
		}
		wg.Wait()

		beforeProcessFiles(gofiles)
		processFiles(gofiles, globalArgs.Debug)

		for _, dir := range goModDirs {
			utils.RunGoImports(dir)
		}
	}
	return nil
}

var (
	funcSet   mapset.Set[string]
	goModDirs []string
	wg        sync.WaitGroup
)

func processFiles(gofiles []string, debug bool) {
	for _, path := range gofiles {
		processFile(path, "", debug)
	}
}

func beforeProcessFiles(gofiles []string) {
	for _, path := range gofiles {
		beforeProcessFile(path)
	}
}

func beforeProcessFile(path string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		nethttp.CollectHandlerFuncName(c, funcSet)
		return true
	}, nil)
}

func processFile(path, printMode string, debug bool) {
	var mode parser.Mode
	fset := token.NewFileSet()
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		mode = 0
	} else {
		mode = parser.ParseComments
	}

	file, err := parser.ParseFile(fset, path, nil, mode)
	if err != nil {
		log.Fatal(err)
	}

	processAST(file, fset)

	if debug {
		if printMode == "console" {
			printer.Fprint(os.Stdout, fset, file)
		} else {
			ast.Print(fset, file)
		}
		return
	}

	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, file); err != nil {
		log.Println(err)
	}

	replace := formatCodeAfterReplace(fset, buf)
	outputPath := path

	if err := os.WriteFile(outputPath, replace.Bytes(), os.ModePerm); err == nil {
		log.Println("File updated:", outputPath)
	} else {
		log.Println(err)
	}
}

func processAST(file *ast.File, fset *token.FileSet) {
	astutil.AddNamedImport(fset, file, "hzserver", globalArgs.HzRepo+"/pkg/server")
	astutil.AddNamedImport(fset, file, "hzapp", globalArgs.HzRepo+"/pkg/app")

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		nethttp.GetOptionsFromHttpServer(c, globalMap)
		nethttp.PackServerHertz(c, globalMap)
		chi.PackChiMux(c)
		nethttp.ReplaceNetHttpHandler(c)
		nethttp.PackSetStatusCode(c)
		return true
	}, nil)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		chiGroup(c)
		netHttpGroup(c, funcSet)
		return true
	}, nil)
}

func chiGroup(c *astutil.Cursor) {
	chi.PackChiRouterMethod(c)
	chi.PackChiNewRouter(c, globalMap)
	chi.PackChiMux(c)
}

func netHttpGroup(c *astutil.Cursor, funcSet mapset.Set[string]) {
	funcsToProcess := []func(c *astutil.Cursor){
		nethttp.PackHandleFunc,
		nethttp.PackFprintf,
		nethttp.ReplaceHttpError,
		nethttp.ReplaceHttpRedirect,
		nethttp.ReplaceRequestURI,
		nethttp.ReplaceReqMethod,
		nethttp.ReplaceReqHost,
		nethttp.ReplaceReqHeader,
		nethttp.ReplaceReqHeaderOperation,
		nethttp.ReplaceRespHeader,
		nethttp.ReplaceRespWrite,
		nethttp.ReplaceRespNotFound,
		nethttp.ReplaceReqURLQuery,
		nethttp.ReplaceReqURLString,
		nethttp.ReplaceReqURLPath,
		nethttp.ReplaceReqCookie,
		nethttp.ReplaceReqFormFile,
		nethttp.ReplaceReqFormGet,
		nethttp.ReplaceReqFormValue,
		nethttp.ReplaceReqMultipartForm,
		func(c *astutil.Cursor) {
			nethttp.PackType2AppHandlerFunc(c)
			nethttp.ReplaceReqMultipartFormOperation(c, globalMap)
			nethttp.ReplaceFuncBodyHttpHandlerParam(c, funcSet)
		},
	}

	for _, fn := range funcsToProcess {
		fn(c)
	}
}

func formatCodeAfterReplace(fset *token.FileSet, buf bytes.Buffer) *bytes.Buffer {
	file, _ := parser.ParseFile(fset, "", buf.String(), parser.ParseComments)

	var output bytes.Buffer
	cfg := printer.Config{
		Mode:     printer.UseSpaces,
		Tabwidth: 4,
	}
	err := cfg.Fprint(&output, fset, file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &output
}
