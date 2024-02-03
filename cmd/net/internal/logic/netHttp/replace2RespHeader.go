package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func Replace2RespHeader(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}

	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}

	if selExpr.Sel.Name == "Header" {
		if utils.CheckStructName(selExpr, "ResponseWriter") {
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
