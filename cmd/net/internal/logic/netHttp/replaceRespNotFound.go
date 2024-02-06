package netHttp

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceRespNotFound(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}

	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}

	ident, ok := selExpr.X.(*Ident)
	if !ok || ident.Name != "http" || selExpr.Sel.Name != "NotFound" {
		return
	}

	callExpr.Fun = &SelectorExpr{
		X:   NewIdent("c"),
		Sel: NewIdent("NotFound"),
	}
	callExpr.Args = []Expr{}
}
