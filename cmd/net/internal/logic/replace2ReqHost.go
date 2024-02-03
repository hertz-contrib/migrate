package logic

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

// Replace2ReqHost replaces r.Host with string(c.Host)
func Replace2ReqHost(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(assignStmt.Rhs) != 1 {
		return
	}

	selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "Host" {
		return
	}

	if utils.CheckPtrStructName(selExpr, "Request") {
		// Create a new expression
		newExpr := &CallExpr{
			Fun: &Ident{Name: "string"},
			Args: []Expr{
				&CallExpr{
					Fun: &SelectorExpr{
						X:   &SelectorExpr{X: &Ident{Name: "c"}, Sel: &Ident{Name: "Request"}},
						Sel: &Ident{Name: "Host"},
					},
				},
			},
		}

		// Replace the right-hand side of the assignment statement
		assignStmt.Rhs = []Expr{newExpr}
	}
}
