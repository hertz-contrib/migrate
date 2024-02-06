package logic

import (
	"bytes"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/migrate/cmd/net/internal/args"
	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	"github.com/hertz-contrib/migrate/cmd/net/internal/logic/chi"
	nethttp "github.com/hertz-contrib/migrate/cmd/net/internal/logic/netHttp"
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"log"
	"os"
	"path/filepath"
)

var funcSet mapset.Set[string]

func init() {
	config.Map = make(map[string]interface{})
	funcSet = mapset.NewSet[string]()
}

func Run(opt args.Args) {
	if opt.Debug {
		beforeProcessFile(opt.Filepath)
		processFile(opt.Filepath, opt.PrintMode, opt.Debug)
		return
	}

	if opt.TargetDir != "" {
		gofiles, err := utils.CollectGoFiles(opt.TargetDir)
		if err != nil {
			log.Fatal(err)
		}
		beforeProcessFiles(gofiles)
		processFiles(gofiles, opt.Debug)
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
	path, _ = filepath.Abs(path)
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
	astutil.Apply(file, func(c *astutil.Cursor) bool {
		nethttp.GetOptionsFromHttpServer(c)
		nethttp.PackServerHertz(c, fset, file)
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
		nethttp.ReplaceReqURLQuery,
		nethttp.ReplaceReqURLString,
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
