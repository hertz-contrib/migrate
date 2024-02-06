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
		switch st := stmt.(type) {
		case *ExprStmt:
			switch t := st.X.(type) {
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
		case *IfStmt:
			assignStmt, ok := st.Init.(*AssignStmt)
			if ok {
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
			callExpr, ok := st.Cond.(*CallExpr)
			if ok {
				ident, ok := callExpr.Fun.(*Ident)
				if ok {
					replaceCallExprParams(funcSet, callExpr, ident.Name)
					continue
				}
				selExpr, ok := callExpr.Fun.(*SelectorExpr)
				if ok {
					replaceCallExprParams(funcSet, callExpr, selExpr.Sel.Name)
					continue
				}
			}

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
							replaceCallExprParams(funcSet, callExpr, callExpr.Fun.(*Ident).Name)
							continue
						}
						ident, ok := callExpr.Fun.(*Ident)
						if ok {
							replaceCallExprParams(funcSet, callExpr, ident.Name)
							continue
						}
						selExpr, ok := callExpr.Fun.(*SelectorExpr)
						if ok {
							replaceCallExprParams(funcSet, callExpr, selExpr.Sel.Name)
							continue
						}
						continue
					}

					ifStmt, ok := _case.(*IfStmt)
					if ok {
						assignStmt, ok := ifStmt.Init.(*AssignStmt)
						if ok {
							if len(assignStmt.Rhs) == 1 {
								callExpr, ok := assignStmt.Rhs[0].(*CallExpr)
								if !ok {
									continue
								}
								switch cc := callExpr.Fun.(type) {
								case *SelectorExpr:
									replaceCallExprParams(funcSet, callExpr, cc.Sel.Name)
									continue
								case *Ident:
									replaceCallExprParams(funcSet, callExpr, callExpr.Fun.(*Ident).Name)
									continue
								}
							}
						}

						callExpr, ok := ifStmt.Cond.(*CallExpr)
						if ok {
							ident, ok := callExpr.Fun.(*Ident)
							if ok {
								replaceCallExprParams(funcSet, callExpr, ident.Name)
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
		case *AssignStmt:
			for _, rh := range st.Rhs {
				ce, ok := rh.(*CallExpr)
				if !ok {
					continue
				}
				ident, ok := ce.Fun.(*Ident)
				if ok {
					replaceCallExprParams(funcSet, ce, ident.Name)
					continue
				}
				selExpr, ok := ce.Fun.(*SelectorExpr)
				if ok {
					replaceCallExprParams(funcSet, ce, selExpr.Sel.Name)
					continue
				}
			}
		}
	}
}

func replaceCallExprParams(funcSet map[string][2]int, callExpr *CallExpr, funcName string) {
	var (
		rwIndex = -1
		rIndex  = -1
	)
	_, ok := funcSet[funcName]
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
