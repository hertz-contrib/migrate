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

package gin

import (
	. "go/ast"
	"go/token"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/types"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceGinCtxOp(call *CallExpr, cur *astutil.Cursor) {
	if se, ok := call.Fun.(*SelectorExpr); ok {
		if ident, ok := se.X.(*Ident); ok {
			if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
				switch se.Sel.Name {
				case "Next":
					cur.Replace(types.ExportCtxNext(ident.Name))
				case "GetHeader":
					cur.Replace(types.ExportCtxGetHeader(ident.Name, call.Args))
				case "Cookie":
					as := cur.Parent().(*AssignStmt)
					// Remove the second assignment
					as.Lhs = as.Lhs[:1]
					cur.Replace(types.ExportCtxCookie(ident.Name, call.Args))
				case "SetCookie":
					var callArgs []Expr
					for index, elt := range call.Args {
						if index == 5 {
							callArgs = append(callArgs, &BasicLit{Value: "0", Kind: token.INT})
							callArgs = append(callArgs, elt)
							continue
						}
						callArgs = append(callArgs, elt)
					}
					call.Args = callArgs
				case "Redirect":
					redirectURIString := call.Args[1]
					cur.Replace(types.ExportCallRedirect(ident.Name, call.Args[0], redirectURIString))
				case "GetRawData":
					if as, ok := cur.Parent().(*AssignStmt); ok {
						as.Lhs = as.Lhs[:1]
					}
				}
			}
		}
	}
}
