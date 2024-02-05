package netHttp

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceFuncBodyHttpHandlerParam(cur *astutil.Cursor, funcSet map[string][2]int) {
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	funcInBodyList := blockStmt.List
	for _, stmt := range funcInBodyList {
		exprStmt, ok := stmt.(*ExprStmt)
		if ok {
			switch t := exprStmt.X.(type) {
			case *CallExpr:
				ident, ok := t.Fun.(*Ident)
				if ok {
					replaceCallExprParams(funcSet, t, ident.Name)
					continue
				}
				selExpr, ok := t.Fun.(*SelectorExpr)
				if ok {
					replaceCallExprParams(funcSet, t, selExpr.Sel.Name)
					continue
				}
			}
		}
		ifStmt, ok := stmt.(*IfStmt)
		if ok {
			assignStmt, ok := ifStmt.Init.(*AssignStmt)
			if !ok {
				continue
			}
			if len(assignStmt.Rhs) == 1 {
				callExpr, ok := assignStmt.Rhs[0].(*CallExpr)
				if !ok {
					continue
				}
				selExpr, ok := callExpr.Fun.(*SelectorExpr)
				if ok {
					replaceCallExprParams(funcSet, callExpr, selExpr.Sel.Name)
					continue
				}
			}
		}
	}
}

func replaceCallExprParams(funcSet map[string][2]int, callExpr *CallExpr, funcName string) {
	_, ok := funcSet[funcName]
	var (
		rwIndex = -1
		rIndex  = -1
	)
	if ok {
		for i, arg := range callExpr.Args {
			indent, ok := arg.(*Ident)
			if !ok {
				continue
			}
			if utils.CheckStarProp(indent, "Request") {
				rIndex = i
			}

			if utils.CheckProps(indent, "ResponseWriter") {
				rwIndex = i
			}
		}
		// Only w http.ResponseWriter
		if rIndex == -1 && rwIndex != -1 {
			callExpr.Args[rwIndex].(*Ident).Name = "c"
		}

		// Only r *http.Request
		if rwIndex == -1 && rIndex != -1 {
			callExpr.Args[rIndex].(*Ident).Name = "c"
		}

		// r *http.Request and w http.ResponseWriter
		if rIndex != -1 && rwIndex != -1 && len(callExpr.Args) == 2 {
			callExpr.Args[rIndex].(*Ident).Name = "c"
			callExpr.Args[rwIndex].(*Ident).Name = "ctx"
		}

		// r *http.Request and w http.ResponseWriter
		if rIndex != -1 && rwIndex != -1 && len(callExpr.Args) > 2 {
			callExpr.Args = append(callExpr.Args[:rwIndex], callExpr.Args[rwIndex+1:]...)
			callExpr.Args[rwIndex].(*Ident).Name = "c"
		}
	}
}
