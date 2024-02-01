package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	"github.com/hertz-contrib/migrate/cmd/net/internal/logic"

	"golang.org/x/tools/go/ast/astutil"

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

	logic.AliasMap = logic.GetAllAliasForPackage(fset, file)
	cfg := config.NewConfig()
	funcSet := mapset.NewSet[string]()

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		logic.GetOptionsFromHttpServer(c, cfg)
		return true
	}, nil)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		logic.PackNewServeMux(c, fset, file, cfg)
		logic.PackHandleFunc(c)
		logic.PackSetStatusCode(c)
		logic.PackFprintf(c)
		logic.PackListenAndServe(c, cfg)
		logic.ReplaceNetHttpHandler(c, funcSet)
		logic.ReplaceHttpError(c)
		logic.ReplaceRequestURI(c)
		logic.ReplaceReqMethod(c)
		return true
	}, nil)

	if _args.PrintMode == "console" {
		printer.Fprint(os.Stdout, fset, file)
		return
	} else {
		ast.Print(fset, file)
	}
}
