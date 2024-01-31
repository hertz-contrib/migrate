package main

import (
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"

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
	utils.AliasMap = utils.GetAllAliasForPackage(fset, file)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		utils.PackHandleFunc(c, fset, file)
		utils.PackFprintf(c, fset, file)
		return true
	}, nil)
	printer.Fprint(os.Stdout, fset, file)
	//ast.Print(fset, file)
}
