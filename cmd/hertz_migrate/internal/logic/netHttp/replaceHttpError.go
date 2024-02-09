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
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceHttpError(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}

	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}

	ident, ok := selExpr.X.(*Ident)
	if !ok || ident.Name != "http" || selExpr.Sel.Name != "Error" {
		return
	}

	callExpr.Args = callExpr.Args[1:]
	lit, ok := callExpr.Args[0].(*BasicLit)
	if ok {
		if lit.Kind == token.STRING {
			if lit.Value == "\"\"" {
				callExpr.Args = callExpr.Args[1:]
				callExpr.Fun = &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("AbortWithStatus"),
				}
				return
			}
		}
	}

	callExpr.Fun = &SelectorExpr{
		X:   NewIdent("c"),
		Sel: NewIdent("AbortWithMsg"),
	}
}
