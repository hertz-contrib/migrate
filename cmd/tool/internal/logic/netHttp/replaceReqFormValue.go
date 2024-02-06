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

func ReplaceReqFormValue(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}
	if selExpr.Sel.Name == "FormValue" {
		if utils.CheckPtrStructName(selExpr, "Request") {
			args := callExpr.Args
			cur.Replace(&CallExpr{
				Fun: NewIdent("string"),
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("FormValue"),
						},
						Args: args,
					},
				},
			})
		}
	}
}
