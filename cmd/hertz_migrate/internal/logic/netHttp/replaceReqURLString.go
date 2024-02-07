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

func ReplaceReqURLString(cur *astutil.Cursor) {
	selExpr, ok := cur.Node().(*SelectorExpr)
	if !ok || selExpr.Sel == nil || selExpr.Sel.Name != "String" {
		return
	}
	ce, ok := selExpr.X.(*SelectorExpr)
	if !ok || ce.Sel.Name != "URL" {
		return
	}
	selExpr.X = &CallExpr{
		Fun: &SelectorExpr{
			X:   NewIdent("c"),
			Sel: NewIdent("URI"),
		},
	}
}
