package logic

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func Replace2RequestURI(cur *astutil.Cursor) {
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
			if ident, ok := selExpr.X.(*Ident); ok {
				if field, ok := ident.Obj.Decl.(*Field); ok {
					starExpr, ok := field.Type.(*StarExpr)
					if !ok {
						return
					}
					selExpr, ok := starExpr.X.(*SelectorExpr)
					if !ok {
						return
					}
					if selExpr.Sel.Name == "Request" {
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
	}
}
