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
	"sync"

	"github.com/hertz-contrib/migrate/cmd/tool/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

var stringMethodExpr *CallExpr

func ReplaceReqMethod(cur *astutil.Cursor) {
	var once sync.Once
	once.Do(func() {
		stringMethodExpr = &CallExpr{
			Fun: &Ident{Name: "string"},
			Args: []Expr{
				&CallExpr{
					Fun: &SelectorExpr{
						X:   &Ident{Name: "c"},
						Sel: &Ident{Name: "Method"},
					},
				},
			},
		}
	})
	replaceAssignStmtReqMethod(cur)
	replaceIfStmtReqMethod(cur)
	replaceSwitchStmtReqMethod(cur)
}

func replaceSwitchStmtReqMethod(cur *astutil.Cursor) {
	switchStmt, ok := cur.Node().(*SwitchStmt)
	if !ok {
		return
	}
	selExpr, ok := switchStmt.Tag.(*SelectorExpr)
	if !ok {
		return
	}
	if utils.CheckPtrStructName(selExpr, "Request") && selExpr.Sel.Name == "Method" {
		switchStmt.Tag = stringMethodExpr
	}
}

func replaceAssignStmtReqMethod(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok {
		return
	}
	if len(assignStmt.Rhs) == 1 {
		selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "Method" {
			if utils.CheckPtrStructName(selExpr, "Request") {
				assignStmt.Rhs[0] = stringMethodExpr
			}
		}
	}
}
func replaceIfStmtReqMethod(cur *astutil.Cursor) {
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
	if utils.CheckPtrStructName(se, "Request") {
		if se.Sel.Name == "Method" {
			be := &BinaryExpr{
				X:  stringMethodExpr,
				Y:  binaryExpr.Y,
				Op: binaryExpr.Op,
			}
			ifStmt.Cond = be
		}
	}
}
