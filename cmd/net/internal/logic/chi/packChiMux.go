package chi

import (
	. "go/ast"
	"go/token"

	"github.com/hertz-contrib/migrate/cmd/net/internal/global"
	"golang.org/x/tools/go/ast/astutil"
)

func PackChiMux(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	funcType, ok := cur.Node().(*FuncType)
	if !ok || funcType.Results == nil {
		return
	}

	if len(funcType.Results.List) == 1 {
		starExpr, ok := funcType.Results.List[0].Type.(*StarExpr)
		if !ok {
			return
		}
		selExpr, ok := starExpr.X.(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "Mux" && selExpr.X.(*Ident).Name == "chi" {
			selExpr.X.(*Ident).Name = "server"
			selExpr.Sel.Name = "Hertz"
			astutil.AddImport(fset, file, global.HzRepo+"/pkg/app/server")
		}
	}
}
