// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gin

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceCtxParamList(cur *astutil.Cursor) {
	ifStmtReplaceLogic := func(st *IfStmt) {
		if assignStmt, ok := st.Init.(*AssignStmt); ok {
			if len(assignStmt.Rhs) == 1 {
				if callExpr, ok := assignStmt.Rhs[0].(*CallExpr); ok {
					replaceParamsWithFuncName(callExpr)
					return
				}
			}
		}

		switch node := st.Cond.(type) {
		case *CallExpr:
			replaceParamsWithFuncName(node)
		case *BinaryExpr:
			inspectBinary(node)
		}
	}

	assignStmtReplaceLogic := func(as *AssignStmt) {
		for _, rhs := range as.Rhs {
			switch node := rhs.(type) {
			case *CallExpr:
				replaceParamsWithFuncName(node)
				continue
			case *TypeAssertExpr:
				if callExpr, ok := node.X.(*CallExpr); ok {
					replaceParamsWithFuncName(callExpr)
					continue
				}
			}
		}
	}

	// Remove the function name from the set if it is found in the current file
	internal.CtxSet.Each(func(s string) bool {
		if internal.WebCtxSet.Contains(s) {
			internal.WebCtxSet.Remove(s)
		}
		return true
	})

	if blockStmt, ok := cur.Node().(*BlockStmt); ok {
		for _, stmt := range blockStmt.List {
			switch _stmt := stmt.(type) {
			case *ExprStmt:
				switch t := _stmt.X.(type) {
				case *CallExpr:
					replaceParamsWithFuncName(t)
					continue
				}
			case *IfStmt:
				ifStmtReplaceLogic(_stmt)
				continue
			case *SwitchStmt:
				for _, s := range _stmt.Body.List {
					if caseClause, ok := s.(*CaseClause); ok {
						for _, _case := range caseClause.Body {
							switch node := _case.(type) {
							case *ExprStmt:
								if callExpr, ok := node.X.(*CallExpr); ok {
									replaceParamsWithFuncName(callExpr)
									continue
								}
							case *IfStmt:
								ifStmtReplaceLogic(node)
								continue
							case *AssignStmt:
								assignStmtReplaceLogic(node)
								continue
							case *ReturnStmt:
								for _, field := range node.Results {
									if ce, ok := field.(*CallExpr); ok {
										replaceParamsWithFuncName(ce)
										continue
									}
								}
							}
						}
					}
				}
			case *AssignStmt:
				assignStmtReplaceLogic(_stmt)
				continue
			case *ReturnStmt:
				for _, field := range _stmt.Results {
					if ce, ok := field.(*CallExpr); ok {
						replaceParamsWithFuncName(ce)
						continue
					}
				}
			}
		}
	}

	if call, ok := cur.Node().(*CallExpr); ok {
		for _, elt := range call.Args {
			if elt, ok := elt.(*CallExpr); ok {
				replaceParamsWithFuncName(elt)
			}
		}
	}
}

func replaceParamsWithFuncName(callExpr *CallExpr) {
	switch node := callExpr.Fun.(type) {
	case *Ident:
		replaceCallExprParams(callExpr, node.Name)
	case *SelectorExpr:
		replaceCallExprParams(callExpr, node.Sel.Name)
	}
}

func replaceCallExprParams(callExpr *CallExpr, funcName string) {
	if internal.WebCtxSet.Contains(funcName) {
		for index, arg := range callExpr.Args {
			if ident, ok := arg.(*Ident); ok {
				if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
					var fieldList []Expr
					fieldList = append(fieldList, callExpr.Args[:index]...)
					fieldList = append(fieldList, NewIdent("_ctx"))
					fieldList = append(fieldList, callExpr.Args[index:]...)
					callExpr.Args = fieldList
					return
				}
			}
		}
	}

	if internal.CtxSet.Contains(funcName) {
		for index, arg := range callExpr.Args {
			if ident, ok := arg.(*Ident); ok {
				if utils.CheckObjSelExpr(ident.Obj, "context", "Context") ||
					utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
					callExpr.Args[index] = NewIdent("_ctx")
					break
				}
			}
		}
	}
}

func inspectBinary(binary *BinaryExpr) {
	switch node := binary.X.(type) {
	case *CallExpr:
		replaceParamsWithFuncName(node)
	case *BinaryExpr:
		inspectBinary(node)
	}

	switch node := binary.Y.(type) {
	case *CallExpr:
		replaceParamsWithFuncName(node)
	case *BinaryExpr:
		inspectBinary(node)
	case *UnaryExpr:
		if callExpr, ok := node.X.(*CallExpr); ok {
			replaceParamsWithFuncName(callExpr)
		}
	}
}
