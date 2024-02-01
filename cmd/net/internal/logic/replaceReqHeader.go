package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqHeader(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(assignStmt.Rhs) != 1 {
		return
	}

	selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "Header" {
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

		// Create a new expression for the replacement
		callExpr := &SelectorExpr{
			X: &SelectorExpr{
				X:   NewIdent("c"),
				Sel: NewIdent("Request"),
			},
			Sel: NewIdent("Header"),
		}
		// Replace the right-hand side of the assignment statement
		assignStmt.Rhs[0] = callExpr
	}
}
