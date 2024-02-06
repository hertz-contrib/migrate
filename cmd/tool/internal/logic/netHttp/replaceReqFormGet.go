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

func ReplaceReqFormGet(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}
	if selExpr.Sel.Name == "Get" {
		se, ok := selExpr.X.(*SelectorExpr)
		if !ok {
			return
		}
		if se.Sel.Name == "Form" || se.Sel.Name == "PostForm" {
			callExpr.Fun = NewIdent("string")
			callExpr.Args = []Expr{
				&CallExpr{
					Fun: &SelectorExpr{
						X:   NewIdent("c"),
						Sel: NewIdent("FormValue"),
					},
					Args: callExpr.Args,
				},
			}
		}
	}
}
