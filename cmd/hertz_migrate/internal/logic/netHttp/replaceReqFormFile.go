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

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqFormFile(cur *astutil.Cursor) {
	stmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(stmt.Lhs) != 3 || len(stmt.Rhs) != 1 {
		return
	}

	ce, ok := stmt.Rhs[0].(*CallExpr)
	if !ok || len(ce.Args) != 1 {
		return
	}

	selExpr, ok := ce.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "FormFile" {
		return
	}

	if utils.CheckPtrStructName(selExpr, "Request") {
		se := &SelectorExpr{
			X: &SelectorExpr{
				X:   NewIdent("c"),
				Sel: NewIdent("Request"),
			},
			Sel: NewIdent("FormFile"),
		}
		ce.Fun = se
		stmt.Lhs = stmt.Lhs[1:]
	}
}
