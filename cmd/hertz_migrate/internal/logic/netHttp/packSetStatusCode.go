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

func PackSetStatusCode(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}
	if selExpr.Sel == nil {
		return
	}
	if selExpr.Sel.Name == "WriteHeader" {
		if ident, ok := selExpr.X.(*Ident); ok {
			if field, ok := ident.Obj.Decl.(*Field); ok {
				expr, ok := field.Type.(*SelectorExpr)
				if ok {
					if expr.Sel.Name == "ResponseWriter" {
						se := &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("SetStatusCode"),
						}
						callExpr.Fun = se
					}
				}
			}
		}
	}
}
