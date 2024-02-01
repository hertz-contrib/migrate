package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func replaceReqMethod(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}
}
