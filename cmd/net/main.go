package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/hertz-contrib/migrate/cmd/net/internal/config"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"

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
	options := config.NewHertzOption()

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		utils.GetOptionsFromHttpServer(c, options)
		return true
	}, nil)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		utils.PackNewServeMux(c, fset, file, options)
		utils.PackHandleFunc(c, fset, file)
		utils.PackFprintf(c)
		return true
	}, nil)

	if _args.PrintMode == "console" {
		printer.Fprint(os.Stdout, fset, file)
		return
	}
	ast.Print(fset, file)
}
