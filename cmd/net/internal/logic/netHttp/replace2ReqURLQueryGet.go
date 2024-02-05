package netHttp

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqURLQuery(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok || len(callExpr.Args) != 1 {
		return
	}

	_selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}

	if _selExpr.Sel.Name == "Get" {
		ce, ok := _selExpr.X.(*CallExpr)
		if !ok {
			return
		}
		se, ok := ce.Fun.(*SelectorExpr)
		if !ok || se.Sel.Name != "Query" {
			return
		}
		callExpr.Fun = &SelectorExpr{
			X:   NewIdent("c"),
			Sel: NewIdent("Query"),
		}
	}
}
