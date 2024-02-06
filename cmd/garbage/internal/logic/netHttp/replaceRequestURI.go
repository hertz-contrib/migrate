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

func ReplaceRequestURI(cur *astutil.Cursor) {
	replaceAssignStmtRequestURI(cur)
	replaceIfStmtRequestURI(cur)
}

func replaceAssignStmtRequestURI(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok {
		return
	}
	if len(assignStmt.Rhs) == 1 {
		selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "RequestURI" {
			if utils.CheckPtrStructName(selExpr, "Request") {
				newExpr := &CallExpr{
					Fun: &Ident{Name: "string"},
					Args: []Expr{
						&CallExpr{
							Fun: &SelectorExpr{
								X: &SelectorExpr{
									X:   &Ident{Name: "c"},
									Sel: &Ident{Name: "Request"},
								},
								Sel: &Ident{Name: "RequestURI"},
							},
						},
					},
				}
				assignStmt.Rhs[0] = newExpr
			}
		}
	}
}

func replaceIfStmtRequestURI(cur *astutil.Cursor) {
	ifStmt, ok := cur.Node().(*IfStmt)
	if !ok {
		return
	}
	binaryExpr, ok := ifStmt.Cond.(*BinaryExpr)
	if !ok {
		return
	}
	se, ok := binaryExpr.X.(*SelectorExpr)
	if !ok {
		return
	}
	if utils.CheckPtrStructName(se, "Request") && se.Sel.Name == "RequestURI" {
		be := &BinaryExpr{
			X: &CallExpr{
				Fun: &Ident{Name: "string"},
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X: &SelectorExpr{
								X:   &Ident{Name: "c"},
								Sel: &Ident{Name: "Request"},
							},
							Sel: &Ident{Name: "RequestURI"},
						},
					},
				},
			},
			Op: binaryExpr.Op,
			Y:  binaryExpr.Y,
		}
		ifStmt.Cond = be
	}
}
