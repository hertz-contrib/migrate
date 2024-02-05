package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqMethod(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok {
		return
	}
	if len(assignStmt.Rhs) == 1 {
		selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "Method" {
			if utils.CheckPtrStructName(selExpr, "Request") {
				newExpr := &CallExpr{
					Fun: &Ident{Name: "string"},
					Args: []Expr{
						&CallExpr{
							Fun: &SelectorExpr{
								X: &SelectorExpr{
									X:   &Ident{Name: "c"},
									Sel: &Ident{Name: "Request"},
								},
								Sel: &Ident{Name: "Method"},
							},
						},
					},
				}
				assignStmt.Rhs[0] = newExpr
			}
		}
	}
}
