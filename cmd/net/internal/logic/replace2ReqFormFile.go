package logic

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func Replace2ReqFormFile(cur *astutil.Cursor) {
	stmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(stmt.Lhs) != 3 || len(stmt.Rhs) != 1 {
		return
	}

	ce, ok := stmt.Rhs[0].(*CallExpr)
	if !ok || len(ce.Args) != 1 {
		return
	}

	selExpr, ok := ce.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "FormFile" {
		return
	}

	if utils.CheckPtrStructName(selExpr, "Request") {
		se := &SelectorExpr{
			X: &SelectorExpr{
				X:   NewIdent("c"),
				Sel: NewIdent("Request"),
			},
			Sel: NewIdent("FormFile"),
		}
		ce.Fun = se
		stmt.Lhs = stmt.Lhs[1:]
	}
}
