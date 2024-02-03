package main

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	chi "github.com/hertz-contrib/migrate/cmd/net/internal/logic/chi"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	nethttp "github.com/hertz-contrib/migrate/cmd/net/internal/logic/netHttp"

	"github.com/hertz-contrib/migrate/cmd/net/internal/args"
	"golang.org/x/tools/go/ast/astutil"
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
	config.Map = make(map[string]any)
	nethttp.AliasMap = nethttp.GetAllAliasForPackage(fset, file)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		nethttp.GetOptionsFromHttpServer(c)
		nethttp.PackServerHertz(c, fset, file)
		chi.PackChiMux(c, fset, file)
		chi.PackChiNewRouter(c, fset, file)
		return true
	}, nil)

	astutil.Apply(file, func(c *astutil.Cursor) bool {
		chiGroup(c)
		netHttpGroup(c, funcSet)
		return true
	}, nil)

	if _args.PrintMode == "console" {
		log.Fatal(printer.Fprint(os.Stdout, fset, file))
		return
	} else {
		ast.Print(fset, file)
	}
}

func chiGroup(c *astutil.Cursor) {
	chi.PackChiRouterMethod(c)
}

func netHttpGroup(c *astutil.Cursor, funcSet mapset.Set[string]) {
	nethttp.PackHandleFunc(c)
	nethttp.PackSetStatusCode(c)
	nethttp.PackFprintf(c)
	nethttp.PackListenAndServe(c)
	nethttp.Replace2NetHttpHandler(c, funcSet)
	nethttp.Replace2HttpError(c)
	nethttp.Replace2RequestURI(c)
	nethttp.Replace2ReqMethod(c)
	nethttp.Replace2ReqHost(c)
	nethttp.Replace2ReqHeader(c)
	nethttp.Replace2ReqHeaderOperation(c)
	nethttp.Replace2RespHeader(c)
	nethttp.Replace2RespWrite(c)
	nethttp.Replace2ReqURLQuery(c)
	nethttp.Replace2ReqURLString(c)
	nethttp.Replace2ReqFormFile(c)
	nethttp.Replace2ReqFormGet(c)
	nethttp.Replace2ReqMultipartForm(c)
	nethttp.Replace2ReqMultipartFormOperation(c)
}
