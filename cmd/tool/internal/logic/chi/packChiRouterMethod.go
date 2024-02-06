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
	"go/token"

	"github.com/hertz-contrib/migrate/cmd/tool/internal/utils"
	"golang.org/x/tools/go/ast/astutil"
)

func PackChiRouterMethod(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok || len(callExpr.Args) < 2 {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok || selExpr.Sel == nil {
		return
	}

	if len(callExpr.Args) == 2 {
		switch selExpr.Sel.Name {
		case "Get":
			selExpr.Sel.Name = "GET"
		case "Post":
			selExpr.Sel.Name = "POST"
		case "Put":
			selExpr.Sel.Name = "PUT"
		case "Delete":
			selExpr.Sel.Name = "DELETE"
		case "Patch":
			selExpr.Sel.Name = "PATCH"
		case "Head":
			selExpr.Sel.Name = "HEAD"
		case "Options":
			selExpr.Sel.Name = "OPTIONS"
		}
		basicLit, ok := callExpr.Args[0].(*BasicLit)
		if !ok || basicLit.Kind != token.STRING {
			return
		}

		basicLit.Value = utils.ReplaceParamsInStr(basicLit.Value)
	}

	if selExpr.Sel.Name != "Method" {
		return
	}
	selExpr.Sel.Name = "Handle"
	basicLit, ok := callExpr.Args[1].(*BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		return
	}

	basicLit.Value = utils.ReplaceParamsInStr(basicLit.Value)
}
