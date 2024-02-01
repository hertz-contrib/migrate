package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func PackSetStatusCode(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}
	if selExpr.Sel.Name == "WriteHeader" {
		if ident, ok := selExpr.X.(*Ident); ok {
			if field, ok := ident.Obj.Decl.(*Field); ok {
				ident, ok := field.Type.(*Ident)
				if ok {
					if ident.Name == "ResponseWriter" {
						se := &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("SetStatusCode"),
						}
						callExpr.Fun = se
					}
				}
			}
		}
	}
}
