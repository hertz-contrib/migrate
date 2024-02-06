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

func ReplaceReqMultipartForm(cur *astutil.Cursor) {
	stmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(stmt.Rhs) != 1 {
		return
	}

	se, ok := stmt.Rhs[0].(*SelectorExpr)
	if !ok || se.Sel.Name != "MultipartForm" {
		return
	}

	// Ensure that stmt.Lhs is not nil and has at least one identifier
	if stmt.Lhs == nil || len(stmt.Lhs) == 0 {
		return
	}

	// Assuming you have the utils.CheckPtrStructName function implemented
	if utils.CheckPtrStructName(se, "Request") {
		// Create a new CallExpr with c.MultiPartForm()
		newCallExpr := &CallExpr{
			Fun: &SelectorExpr{
				X:   NewIdent("c"),
				Sel: NewIdent("MultiPartForm"),
			},
		}

		// Replace the original stmt.Rhs with the newCallExpr
		stmt.Rhs = []Expr{newCallExpr}

		// Assuming c.MultiPartForm returns two values (result, error)
		// Create a new identifier "err" and append it to stmt.Lhs
		stmt.Lhs = append(stmt.Lhs, NewIdent("err"))
	}
}
