package netHttp

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceHttpRedirect(cur *astutil.Cursor) {
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	for _, stmt := range blockStmt.List {
		estmt, ok := stmt.(*ExprStmt)
		if !ok {
			continue
		}
		callExpr, ok := estmt.X.(*CallExpr)
		if !ok {
			continue
		}
		selExpr, ok := callExpr.Fun.(*SelectorExpr)
		if ok && selExpr.Sel.Name == "Redirect" && selExpr.X.(*Ident).Name == "http" {
			uriExpr := callExpr.Args[2]
			statusExpr := callExpr.Args[3]
			estmt.X = &CallExpr{
				Fun: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("Redirect"),
				},
				Args: []Expr{
					statusExpr,
					&CallExpr{
						Fun: &ArrayType{
							Elt: NewIdent("byte"),
						},
						Args: []Expr{uriExpr},
					},
				},
			}
		}
	}
}
