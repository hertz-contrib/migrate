package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceRespHeader(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}

	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}
	if selExpr.Sel.Name == "Header" {
		ident, ok := selExpr.X.(*Ident)
		if !ok {
			return
		}
		if field, ok := ident.Obj.Decl.(*Field); ok {
			_selExpr, ok := field.Type.(*SelectorExpr)
			if !ok || _selExpr.Sel.Name != "ResponseWriter" {
				return
			}

			// Create a new expression for the replacement
			callExpr := &SelectorExpr{
				X: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("Response"),
				},
				Sel: NewIdent("Header"),
			}
			// Replace the right-hand side of the assignment statement
			cur.Replace(callExpr)
		}
	}
}
