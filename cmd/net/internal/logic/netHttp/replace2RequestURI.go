package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceRequestURI(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok {
		return
	}
	if len(assignStmt.Rhs) == 1 {
		selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "RequestURI" {
			if utils.CheckPtrStructName(selExpr, "Request") {
				callExpr := &CallExpr{
					Fun: &SelectorExpr{
						X: &CallExpr{
							Fun: &SelectorExpr{
								X: &SelectorExpr{
									X:   NewIdent("c"),
									Sel: NewIdent("Request"),
								},
								Sel: NewIdent("URI"),
							},
						},
						Sel: NewIdent("String"),
					},
				}
				assignStmt.Rhs[0] = callExpr
			}
		}
	}
}
