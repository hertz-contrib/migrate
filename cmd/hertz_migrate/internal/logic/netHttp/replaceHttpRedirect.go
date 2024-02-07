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

func ReplaceHttpRedirect(cur *astutil.Cursor) {
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	for _, stmt := range blockStmt.List {
		estmt, ok := stmt.(*ExprStmt)
		if !ok {
			continue
		}
		callExpr, ok := estmt.X.(*CallExpr)
		if !ok {
			continue
		}
		selExpr, ok := callExpr.Fun.(*SelectorExpr)
		if ok && selExpr.Sel.Name == "Redirect" && selExpr.X.(*Ident).Name == "http" {
			uriExpr := callExpr.Args[2]
			statusExpr := callExpr.Args[3]
			estmt.X = &CallExpr{
				Fun: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("Redirect"),
				},
				Args: []Expr{
					statusExpr,
					&CallExpr{
						Fun: &ArrayType{
							Elt: NewIdent("byte"),
						},
						Args: []Expr{uriExpr},
					},
				},
			}
		}
	}
}
