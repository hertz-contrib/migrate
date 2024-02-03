package netHttp

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func PackHandleFunc(cur *astutil.Cursor) {
	if selExpr, ok := cur.Node().(*SelectorExpr); ok {
		if selExpr.Sel == nil {
			return
		}
		if selExpr.Sel.Name == "HandleFunc" {
			selExpr.Sel.Name = "Any"
		}
	}
}
