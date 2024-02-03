package chi

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func PackChiRouterMethod(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok || len(callExpr.Args) < 2 {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}

	switch selExpr.Sel.Name {
	case "Get":
		selExpr.Sel.Name = "GET"
	case "Post":
		selExpr.Sel.Name = "POST"
	case "Put":
		selExpr.Sel.Name = "PUT"
	case "Delete":
		selExpr.Sel.Name = "DELETE"
	case "Patch":
		selExpr.Sel.Name = "PATCH"
	case "Head":
		selExpr.Sel.Name = "HEAD"
	case "Options":
		selExpr.Sel.Name = "OPTIONS"
	}

	if selExpr.Sel.Name != "Method" {
		return
	}
	selExpr.Sel.Name = "Handle"
}
