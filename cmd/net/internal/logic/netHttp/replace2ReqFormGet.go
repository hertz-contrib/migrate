package netHttp

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqFormGet(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}
	if selExpr.Sel.Name == "Get" {
		selectExpr, ok := selExpr.X.(*SelectorExpr)
		if !ok || selectExpr.Sel.Name != "Form" {
			return
		}

		callExpr.Fun = NewIdent("string")
		callExpr.Args = []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("FormValue"),
				},
				Args: callExpr.Args,
			},
		}
	}
}
