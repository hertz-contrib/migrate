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

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceRespWrite(cur *astutil.Cursor) {
	stmt, ok := cur.Node().(*BlockStmt)
	if !ok || len(stmt.List) == 0 {
		return
	}
	var setStatusCodeInserted bool

	for i, s := range stmt.List {
		es, ok := s.(*ExprStmt)
		if !ok {
			continue
		}
		ce, ok := es.X.(*CallExpr)
		if !ok {
			continue
		}
		selExpr, ok := ce.Fun.(*SelectorExpr)
		if !ok || selExpr.Sel == nil {
			continue
		}

		// 检查是否已经插入了 c.SetStatusCode
		if selExpr.Sel.Name == "SetStatusCode" {
			setStatusCodeInserted = true
			continue
		}

		if selExpr.Sel.Name == "Write" {
			var _es *ExprStmt
			if !setStatusCodeInserted {
				_es = &ExprStmt{
					X: &CallExpr{
						Fun: &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("SetStatusCode"),
						},
						Args: []Expr{NewIdent("200")},
					},
				}
				stmt.List = append(stmt.List[:i], append([]Stmt{_es}, stmt.List[i:]...)...)
			}

			ce.Fun = &SelectorExpr{
				X: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("Response"),
				},
				Sel: NewIdent("SetBody"),
			}
		}
	}
}
