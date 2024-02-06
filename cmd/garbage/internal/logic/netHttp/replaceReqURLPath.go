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

	"github.com/hertz-contrib/migrate/cmd/garbage/internal/utils"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqURLPath(cur *astutil.Cursor) {
	replaceBlockStmtReqURLPath(cur)
	replaceExprStmtReqURLPath(cur)
}

func replaceExprStmtReqURLPath(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	for i, arg := range callExpr.Args {
		arg, ok := arg.(*SelectorExpr)
		if !ok || arg.Sel.Name != "Path" {
			continue
		}
		selExpr, ok := arg.X.(*SelectorExpr)
		if !ok {
			continue
		}
		if utils.CheckPtrStructName(selExpr, "Request") {
			callExpr.Args[i] = &CallExpr{
				Fun: &Ident{Name: "string"},
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X: &CallExpr{
								Fun: &SelectorExpr{
									X:   &Ident{Name: "c"},
									Sel: &Ident{Name: "URI"},
								},
							},
							Sel: &Ident{Name: "Path"},
						},
					},
				},
			}
		}
	}
}

func replaceBlockStmtReqURLPath(cur *astutil.Cursor) {
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	for i, stmt := range blockStmt.List {
		assignStmt, ok := stmt.(*AssignStmt)
		if !ok {
			continue
		}
		if len(assignStmt.Rhs) == 1 {
			selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
			if !ok {
				continue
			}
			if selExpr.Sel.Name == "Path" {
				assignStmt.Rhs[0] = &CallExpr{
					Fun: &Ident{Name: "string"},
					Args: []Expr{
						&CallExpr{
							Fun: &SelectorExpr{
								X: &CallExpr{
									Fun: &SelectorExpr{
										X:   &Ident{Name: "c"},
										Sel: &Ident{Name: "URI"},
									},
								},
								Sel: &Ident{Name: "Path"},
							},
						},
					},
				}
				blockStmt.List[i] = assignStmt
			}
		}
	}
}
