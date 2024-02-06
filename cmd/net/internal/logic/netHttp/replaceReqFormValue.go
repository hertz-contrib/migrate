package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
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
			args := callExpr.Args
			cur.Replace(&CallExpr{
				Fun: NewIdent("string"),
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("FormValue"),
						},
						Args: args,
					},
				},
			})
		}
	}
}
