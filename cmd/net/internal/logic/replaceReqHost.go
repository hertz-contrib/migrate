package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

// ReplaceReqHost replaces r.Host with string(c.Host)
func ReplaceReqHost(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(assignStmt.Rhs) != 1 {
		return
	}

	selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "Host" {
		return
	}

	ident, ok := selExpr.X.(*Ident)
	if !ok {
		return
	}

	if field, ok := ident.Obj.Decl.(*Field); ok {
		starExpr, ok := field.Type.(*StarExpr)
		if !ok {
			return
		}

		selExpr, ok := starExpr.X.(*SelectorExpr)
		if !ok || selExpr.Sel.Name != "Request" {
			return
		}

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
