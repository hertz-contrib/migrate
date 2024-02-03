package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/migrate/cmd/net/internal/logic"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hertz-contrib/migrate/cmd/net/internal/args"
)

var _args args.Args

func main() {
	_args.Parse()
	fset := token.NewFileSet() // positions are relative to fset
	path, _ := filepath.Abs(_args.Filepath)

	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	funcSet := mapset.NewSet[string]()
	logic.AliasMap = logic.GetAllAliasForPackage(fset, file)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		logic.GetOptionsFromHttpServer(c)
		logic.PackServerHertz(c, fset, file)
		return true
	}, nil)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		logic.PackHandleFunc(c)
		logic.PackSetStatusCode(c)
		logic.PackFprintf(c)
		logic.PackListenAndServe(c)
		logic.Replace2NetHttpHandler(c, funcSet)
		logic.Replace2HttpError(c)
		logic.Replace2RequestURI(c)
		logic.Replace2ReqMethod(c)
		logic.Replace2ReqHost(c)
		logic.Replace2ReqHeader(c)
		logic.Replace2ReqHeaderOperation(c)
		logic.Replace2RespHeader(c)
		logic.Replace2RespWrite(c)
		logic.Replace2ReqURLQuery(c)
		logic.Replace2ReqURLString(c)
		logic.Replace2ReqFormFile(c)
		logic.Replace2ReqFormGet(c)
		logic.Replace2ReqMultipartForm(c)
		logic.Replace2ReqMultipartFormOperation(c)
		return true
	}, nil)

	if _args.PrintMode == "console" {
		log.Fatal(printer.Fprint(os.Stdout, fset, file))
		return
	} else {
		ast.Print(fset, file)
	}
}
