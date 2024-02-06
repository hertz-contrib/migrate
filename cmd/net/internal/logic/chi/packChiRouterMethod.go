package chi

import (
	. "go/ast"
	"go/token"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
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

	if len(callExpr.Args) == 2 {
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
		basicLit, ok := callExpr.Args[0].(*BasicLit)
		if !ok || basicLit.Kind != token.STRING {
			return
		}

		basicLit.Value = utils.ReplaceParamsInStr(basicLit.Value)
	}

	if selExpr.Sel.Name != "Method" {
		return
	}
	selExpr.Sel.Name = "Handle"
	basicLit, ok := callExpr.Args[1].(*BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		return
	}

	basicLit.Value = utils.ReplaceParamsInStr(basicLit.Value)
}
