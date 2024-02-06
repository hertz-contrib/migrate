package netHttp

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqFormValue(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}
	if selExpr.Sel.Name == "FormValue" {
		if utils.CheckPtrStructName(selExpr, "Request") {
			selExpr.X = NewIdent("c")
		}
	}
}
