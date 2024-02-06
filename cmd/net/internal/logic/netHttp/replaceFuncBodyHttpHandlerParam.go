package netHttp

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceFuncBodyHttpHandlerParam(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	funcInBodyList := blockStmt.List

	ifStmtReplaceLogic := func(st *IfStmt) {
		assignStmt, ok := st.Init.(*AssignStmt)
		if ok {
			if len(assignStmt.Rhs) == 1 {
				callExpr, ok := assignStmt.Rhs[0].(*CallExpr)
				if ok {
					replaceCallExprParamsWithFuncName(funcSet, callExpr)
					return
				}

			}
		}
		callExpr, ok := st.Cond.(*CallExpr)
		if ok {
			replaceCallExprParamsWithFuncName(funcSet, callExpr)
			return
		}
	}

	for _, stmt := range funcInBodyList {
		switch st := stmt.(type) {
		case *ExprStmt:
			switch t := st.X.(type) {
			case *CallExpr:
				replaceCallExprParamsWithFuncName(funcSet, t)
				continue
			}
		case *IfStmt:
			ifStmtReplaceLogic(st)
			continue
		case *SwitchStmt:
			for _, s := range st.Body.List {
				caseClause, ok := s.(*CaseClause)
				if !ok {
					continue
				}
				for _, _case := range caseClause.Body {
					exprStmt, ok := _case.(*ExprStmt)
					if ok {
						callExpr, ok := exprStmt.X.(*CallExpr)
						if ok {
							replaceCallExprParamsWithFuncName(funcSet, callExpr)
							continue
						}
					}

					ifStmt, ok := _case.(*IfStmt)
					if ok {
						ifStmtReplaceLogic(ifStmt)
						continue
					}
				}
			}
		case *AssignStmt:
			for _, rh := range st.Rhs {
				ce, ok := rh.(*CallExpr)
				if ok {
					replaceCallExprParamsWithFuncName(funcSet, ce)
					continue
				}
			}
		}
	}
}

func replaceCallExprParamsWithFuncName(funcSet mapset.Set[string], callExpr *CallExpr) {
	ident, ok := callExpr.Fun.(*Ident)
	if ok {
		replaceCallExprParams(funcSet, callExpr, ident.Name)
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if ok {
		replaceCallExprParams(funcSet, callExpr, selExpr.Sel.Name)
		return
	}
}

func replaceCallExprParams(funcSet mapset.Set[string], callExpr *CallExpr, funcName string) {
	var (
		rwIndex = -1
		rIndex  = -1
	)
	ok := funcSet.Contains(funcName)
	if ok {
		for i, arg := range callExpr.Args {
			indent, ok := arg.(*Ident)
			if !ok {
				continue
			}
			if utils.CheckStarProp(indent, "Request") {
				rIndex = i
			}

			if utils.CheckProp(indent, "ResponseWriter") {
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
			callExpr.Args[rwIndex].(*Ident).Name = "c"
			callExpr.Args[rIndex].(*Ident).Name = "ctx"
		}
	}
}
