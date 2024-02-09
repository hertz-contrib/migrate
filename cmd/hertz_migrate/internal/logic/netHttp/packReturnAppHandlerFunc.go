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

func PackType2AppHandlerFunc(cur *astutil.Cursor) {
	packReturnStmt2AppHandlerFunc(cur)
}

func packReturnStmt2AppHandlerFunc(cur *astutil.Cursor) {
	funcDecl, ok := cur.Node().(*FuncDecl)
	if !ok {
		return
	}
	funcType := funcDecl.Type
	if funcType.Results == nil {
		return
	}
	fields := funcType.Results.List
	if len(fields) == 1 {
		ft, ok := fields[0].Type.(*FuncType)
		if !ok {
			return
		}
		if len(ft.Params.List) != 2 {
			return
		}
		if ft.Params.List[0].Names[0].Name == "ctx" && ft.Params.List[1].Names[0].Name == "c" {
			funcType.Results.List = []*Field{
				{
					Type: &SelectorExpr{
						X:   NewIdent("hzapp"),
						Sel: NewIdent("HandlerFunc"),
					},
				},
			}
		}
	}
}
