package logic

import (
	. "go/ast"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

func Replace2HttpError(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}

	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}

	ident, ok := selExpr.X.(*Ident)
	if !ok || ident.Name != "http" || selExpr.Sel.Name != "Error" {
		return
	}

	callExpr.Args = callExpr.Args[1:]
	lit, ok := callExpr.Args[0].(*BasicLit)
	if ok {
		if lit.Kind == token.STRING {
			if lit.Value == "\"\"" {
				callExpr.Args = callExpr.Args[1:]
				callExpr.Fun = &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("AbortWithStatus"),
				}
				return
			}
		}
	}

	callExpr.Fun = &SelectorExpr{
		X:   NewIdent("c"),
		Sel: NewIdent("AbortWithMsg"),
	}
}
