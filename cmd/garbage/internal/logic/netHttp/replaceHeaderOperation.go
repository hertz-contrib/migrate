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

func ReplaceReqHeaderOperation(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	if len(callExpr.Args) == 1 || len(callExpr.Args) == 2 {
		selExpr, ok := callExpr.Fun.(*SelectorExpr)
		if !ok {
			return
		}
		replaceSetOrAdd(selExpr, callExpr)
		replaceDelOrGet(selExpr, callExpr)
	}
}

func replaceSetOrAdd(selExpr *SelectorExpr, callExpr *CallExpr) {
	if selExpr.Sel.Name == "Set" || selExpr.Sel.Name == "Add" {
		_selExpr, ok := selExpr.X.(*SelectorExpr)
		if !ok {
			return
		}
		if _selExpr.Sel.Name == "Header" {
			fun := &SelectorExpr{
				X: &SelectorExpr{
					X: &SelectorExpr{
						X:   NewIdent("c"),
						Sel: NewIdent("Request"),
					},
					Sel: NewIdent("Header"),
				},
				Sel: NewIdent(selExpr.Sel.Name),
			}
			callExpr.Fun = fun
		}
	}
}

func replaceDelOrGet(selExpr *SelectorExpr, callExpr *CallExpr) {
	if selExpr.Sel.Name == "Del" || selExpr.Sel.Name == "Get" {
		_selExpr, ok := selExpr.X.(*SelectorExpr)
		if !ok {
			return
		}
		if _selExpr.Sel.Name == "Header" {
			fun := &SelectorExpr{
				X: &SelectorExpr{
					X: &SelectorExpr{
						X:   NewIdent("c"),
						Sel: NewIdent("Request"),
					},
					Sel: NewIdent("Header"),
				},
				Sel: NewIdent("Del"),
			}
			callExpr.Fun = fun
		}
	}
}
