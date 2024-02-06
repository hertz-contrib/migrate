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

package logic

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

	"golang.org/x/tools/go/ast/astutil"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/migrate/cmd/tool/internal/args"
	"github.com/hertz-contrib/migrate/cmd/tool/internal/global"
	"github.com/hertz-contrib/migrate/cmd/tool/internal/logic/chi"
	nethttp "github.com/hertz-contrib/migrate/cmd/tool/internal/logic/netHttp"
	"github.com/hertz-contrib/migrate/cmd/tool/internal/utils"
)

var (
	funcSet   mapset.Set[string]
	goModDirs []string
	wg        sync.WaitGroup
)

func init() {
	global.Map = make(map[string]interface{})
	funcSet = mapset.NewSet[string]()

}

func Run(opt args.Args) {
	global.HzRepo = opt.HzRepo
	if opt.TargetDir != "" {
		gofiles, err := utils.CollectGoFiles(opt.TargetDir)
		if err != nil {
			log.Fatal("Error collecting go files:", err)
		}
		goModDirs = utils.SearchAllDirHasGoMod(opt.TargetDir)
		for _, dir := range goModDirs {
			wg.Add(1)
			dir := dir
			go func() {
				defer wg.Done()
				utils.RunGoGet(dir, global.HzRepo)
			}()
		}
		wg.Wait()

		beforeProcessFiles(gofiles)
		processFiles(gofiles, opt.Debug)
		for _, dir := range goModDirs {
			utils.RunGoImports(dir)
		}
	}
}

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
	astutil.AddNamedImport(fset, file, "hzserver", global.HzRepo+"/pkg/server")
	astutil.AddNamedImport(fset, file, "hzapp", global.HzRepo+"/pkg/app")

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		nethttp.GetOptionsFromHttpServer(c)
		nethttp.PackServerHertz(c, fset, file)
		chi.PackChiMux(c)
		nethttp.ReplaceNetHttpHandler(c, fset, file)
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
	chi.PackChiNewRouter(c)
	chi.PackChiMux(c)
}

func netHttpGroup(c *astutil.Cursor, funcSet mapset.Set[string]) {
	funcsToProcess := []func(*astutil.Cursor){
		nethttp.PackHandleFunc,
		nethttp.PackFprintf,
		nethttp.PackListenAndServe,
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
		nethttp.ReplaceReqMultipartFormOperation,
		func(c *astutil.Cursor) {
			nethttp.PackType2AppHandlerFunc(c)
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
