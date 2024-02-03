package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func Replace2ReqURLString(cur *astutil.Cursor) {
	selExpr, ok := cur.Node().(*SelectorExpr)
	if !ok || selExpr.Sel == nil || selExpr.Sel.Name != "String" {
		return
	}
	ce, ok := selExpr.X.(*SelectorExpr)
	if !ok || ce.Sel.Name != "URL" {
		return
	}
	selExpr.X = &CallExpr{
		Fun: &SelectorExpr{
			X:   NewIdent("c"),
			Sel: NewIdent("URI"),
		},
	}
}
