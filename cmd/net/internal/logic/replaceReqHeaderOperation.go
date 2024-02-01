package logic

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqHeaderOperation(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	if len(callExpr.Args) != 2 {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}
	replaceSetOrAdd(selExpr, callExpr)
	replaceDelOrGet(selExpr, callExpr)
}

func replaceSetOrAdd(selExpr *SelectorExpr, callExpr *CallExpr) {
	if selExpr.Sel.Name == "Set" || selExpr.Sel.Name == "Add" {
		_selExpr, ok := selExpr.X.(*SelectorExpr)
		if !ok {
			return
		}
		if _selExpr.Sel.Name == "Header" {
			fun := &SelectorExpr{
				X: &SelectorExpr{
					X: &SelectorExpr{
						X:   NewIdent("c"),
						Sel: NewIdent("Request"),
					},
					Sel: NewIdent("Header"),
				},
				Sel: NewIdent(selExpr.Sel.Name),
			}
			callExpr.Fun = fun
		}
	}
}

func replaceDelOrGet(selExpr *SelectorExpr, callExpr *CallExpr) {
	if selExpr.Sel.Name == "Del" || selExpr.Sel.Name == "Get" {
		_selExpr, ok := selExpr.X.(*SelectorExpr)
		if !ok {
			return
		}
		if _selExpr.Sel.Name == "Header" {
			fun := &SelectorExpr{
				X: &SelectorExpr{
					X: &SelectorExpr{
						X:   NewIdent("c"),
						Sel: NewIdent("Request"),
					},
					Sel: NewIdent("Header"),
				},
				Sel: NewIdent("Del"),
			}
			callExpr.Fun = fun
		}
	}
}
