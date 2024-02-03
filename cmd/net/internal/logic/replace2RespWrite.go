package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func Replace2RespWrite(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}
	if selExpr.Sel.Name == "Write" {
		callExpr.Fun = &SelectorExpr{
			X: &SelectorExpr{
				X:   NewIdent("c"),
				Sel: NewIdent("Response"),
			},
			Sel: NewIdent("Write"),
		}
	}
}
