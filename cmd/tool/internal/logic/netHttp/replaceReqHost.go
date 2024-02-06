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

	"github.com/hertz-contrib/migrate/cmd/tool/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

// ReplaceReqHost replaces r.Host with string(c.Host)
func ReplaceReqHost(cur *astutil.Cursor) {
	replaceAssignStmtReqHost(cur)
	replaceIfStmtReqHost(cur)
}

// replaceAssignStmtReqHost replaces r.Host with string(c.Host) in AssignStmt
func replaceAssignStmtReqHost(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(assignStmt.Rhs) != 1 {
		return
	}

	selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "Host" {
		return
	}

	if utils.CheckPtrStructName(selExpr, "Request") {
		// Create a new expression
		newExpr := &CallExpr{
			Fun: &Ident{Name: "string"},
			Args: []Expr{
				&CallExpr{
					Fun: &SelectorExpr{
						X:   &SelectorExpr{X: &Ident{Name: "c"}, Sel: &Ident{Name: "Request"}},
						Sel: &Ident{Name: "Host"},
					},
				},
			},
		}

		// Replace the right-hand side of the assignment statement
		assignStmt.Rhs = []Expr{newExpr}
	}
}

// replaceIfStmtReqHost replaces r.Host with string(c.Host) in IfStmt
func replaceIfStmtReqHost(cur *astutil.Cursor) {
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

	if utils.CheckPtrStructName(se, "Request") && se.Sel.Name == "Host" {
		be := &BinaryExpr{
			X: &CallExpr{
				Fun: &Ident{Name: "string"},
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X:   &SelectorExpr{X: &Ident{Name: "c"}, Sel: &Ident{Name: "Request"}},
							Sel: &Ident{Name: "Host"},
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
