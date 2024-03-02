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

func GetFuncNameHasGinCtx(c *astutil.Cursor) {
	findInnerFuncName(c)
	findFuncName(c)
	findTypePropName(c)
	findExprStmtFuncName(c)
}

func findTypePropName(cur *astutil.Cursor) {
	var paramList []*Field
	if field, ok := cur.Node().(*Field); ok {
		if funcType, ok := field.Type.(*FuncType); ok {
			paramList = funcType.Params.List
			for _, _field := range paramList {
				switch t := _field.Type.(type) {
				case *StarExpr:
					if selExpr, ok := t.X.(*SelectorExpr); ok {
						if selExpr.Sel.Name == "Context" {
							internal.WebCtxSet.Add(field.Names[0].Name)
						}
					}
				case *SelectorExpr:
					if utils.CheckSelPkgAndStruct(t, "context", "Context") {
						internal.CtxSet.Add(field.Names[0].Name)
					}
				}
			}
		}
	}
}

func findExprStmtFuncName(cur *astutil.Cursor) {
	if exprStmt, ok := cur.Node().(*ExprStmt); ok {
		if callExpr, ok := exprStmt.X.(*CallExpr); ok {
			if ident, ok := callExpr.Fun.(*Ident); ok {
				funcName := ident.Name
				for _, arg := range callExpr.Args {
					switch node := arg.(type) {
					case *Ident:
						if utils.CheckObjStarExpr(node.Obj, "gin", "Context") {
							internal.WebCtxSet.Add(funcName)
						}
					case *SelectorExpr:
						if utils.CheckSelPkgAndStruct(node, "gin", "Context") {
							internal.CtxSet.Add(funcName)
						}
					}
				}
			}
		}
	}
}

func findInnerFuncName(cur *astutil.Cursor) {
	var (
		funcName  string
		paramList []*Field
	)
	if blockStmt, ok := cur.Node().(*BlockStmt); ok {
		for _, stmt := range blockStmt.List {
			if as, ok := stmt.(*AssignStmt); ok {
				if len(as.Lhs) == 1 {
					if ident, ok := as.Lhs[0].(*Ident); ok {
						funcName = ident.Name
					}
				}
				if len(as.Rhs) == 1 {
					if funcLit, ok := as.Rhs[0].(*FuncLit); ok {
						paramList = funcLit.Type.Params.List
						for _, field := range paramList {
							switch t := field.Type.(type) {
							case *StarExpr:
								if selExpr, ok := t.X.(*SelectorExpr); ok {
									if utils.CheckSelPkgAndStruct(selExpr, "gin", "Context") {
										internal.WebCtxSet.Add(funcName)
									}
								}
							case *SelectorExpr:
								if utils.CheckSelPkgAndStruct(t, "context", "Context") {
									internal.CtxSet.Add(funcName)
								}
							}
						}
					}
				}
			}
		}
	}
}

func findFuncName(cur *astutil.Cursor) {
	var paramList []*Field
	if funcDecl, ok := cur.Node().(*FuncDecl); ok {
		funcType := funcDecl.Type
		paramList = funcType.Params.List

		for _, field := range paramList {
			switch t := field.Type.(type) {
			case *StarExpr:
				if selExpr, ok := t.X.(*SelectorExpr); ok {
					if selExpr.Sel.Name == "Context" {
						internal.WebCtxSet.Add(funcDecl.Name.Name)
					}
				}
			case *SelectorExpr:
				if utils.CheckSelPkgAndStruct(t, "context", "Context") {
					internal.CtxSet.Add(funcDecl.Name.Name)
				}
			}
		}
	}
}
