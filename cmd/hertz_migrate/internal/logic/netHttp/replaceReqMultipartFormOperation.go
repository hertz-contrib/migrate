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

package netHttp

import (
	. "go/ast"
	"go/token"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/global"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqMultipartFormOperation(cur *astutil.Cursor) {
	var (
		fIndex     int
		opFuncName string
		lhsName    string
	)
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	for i, stmt := range blockStmt.List {
		assignStmt, ok := stmt.(*AssignStmt)
		if !ok {
			continue
		}
		if len(assignStmt.Rhs) != 1 {
			continue
		}
		selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
		if !ok || selExpr.Sel == nil {
			continue
		}
		success := func() {
			if opFuncName != "" && lhsName != "" {
				fIndexBlock := blockStmt.List[fIndex]
				var blockStmtList []Stmt
				blockStmt.List = append(blockStmt.List[:fIndex], blockStmt.List[fIndex+1:]...)

				blockStmtList = append(blockStmtList, blockStmt.List[:fIndex]...)
				blockStmtList = append(blockStmtList, fIndexBlock)
				blockStmtList = append(blockStmtList, &AssignStmt{
					Lhs: []Expr{
						&Ident{
							Name: lhsName,
							Obj:  NewObj(Var, lhsName),
						},
					},
					Tok: token.DEFINE,
					Rhs: []Expr{
						&SelectorExpr{
							X:   NewIdent("_form"),
							Sel: NewIdent(opFuncName),
						},
					},
				})
				blockStmtList = append(blockStmtList, blockStmt.List[fIndex:]...)
				blockStmt.List = blockStmtList
			}
		}
		if selExpr.Sel.Name == "Value" || selExpr.Sel.Name == "File" {
			se, ok := selExpr.X.(*SelectorExpr)
			if !ok || se.Sel.Name != "MultipartForm" {
				continue
			}
			if utils.CheckPtrStructName(se, "Request") {
				opFuncName = selExpr.Sel.Name
				lhsName = assignStmt.Lhs[0].(*Ident).Name
				if global.Map["hasMultipartForm"] == true {
					blockStmt.List = append(blockStmt.List[:i+1], blockStmt.List[i+2:]...)
					success()
					continue
				}
				fIndex = i
				assignStmt.Rhs = []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("MultiPartForm"),
						},
					},
				}
				assignStmt.Lhs[0] = NewIdent("_form")
				assignStmt.Lhs = append(assignStmt.Lhs, NewIdent("err"))
				global.Map["hasMultipartForm"] = true
				success()
			}
		}
	}
}
