package main

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
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

	utils.AliasMap = utils.GetAllAliasForPackage(fset, file)
	cfg := config.NewConfig()

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		utils.GetOptionsFromHttpServer(c, cfg)
		return true
	}, nil)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		utils.PackNewServeMux(c, fset, file, cfg)
		utils.PackHandleFunc(c, fset, file)
		utils.PackFprintf(c)
		utils.PackListenAndServe(c, cfg)
		return true
	}, nil)

	if _args.PrintMode == "console" {
		printer.Fprint(os.Stdout, fset, file)
		return
	}
	ast.Print(fset, file)
}
