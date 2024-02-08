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

package chi

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func PackChiMux(cur *astutil.Cursor) {
	funcType, ok := cur.Node().(*FuncType)
	if !ok || funcType.Results == nil {
		return
	}

	if len(funcType.Results.List) == 1 {
		starExpr, ok := funcType.Results.List[0].Type.(*StarExpr)
		if !ok {
			return
		}
		selExpr, ok := starExpr.X.(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "Mux" && selExpr.X.(*Ident).Name == "chi" {
			selExpr.X.(*Ident).Name = "hzserver"
			selExpr.Sel.Name = "Hertz"
		}
	}
}
