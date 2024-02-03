package chi

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	nethttp "github.com/hertz-contrib/migrate/cmd/net/internal/logic/netHttp"
	. "go/ast"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
)

func PackChiNewRouter(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	stmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(stmt.Lhs) != 1 || len(stmt.Rhs) != 1 {
		return
	}
	callExpr, ok := stmt.Rhs[0].(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}
	if selExpr.Sel.Name == "NewRouter" {
		callExpr.Fun = &SelectorExpr{
			X:   NewIdent("server"),
			Sel: NewIdent("Default"),
		}
		config.Map["server"] = stmt.Lhs[0].(*Ident).Name
		astutil.AddImport(fset, file, "github.com/cloudwego/hertz/pkg/app/server")
		nethttp.AddOptionsForServer(callExpr, config.Map)
	}
}
