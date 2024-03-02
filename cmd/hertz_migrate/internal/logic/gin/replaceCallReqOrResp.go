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

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/types"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceCallReqOrResp(call *CallExpr, cur *astutil.Cursor) {
	if se, ok := call.Fun.(*SelectorExpr); ok {
		if _se, ok := se.X.(*SelectorExpr); ok {
			if ident, ok := _se.X.(*Ident); ok {
				if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
					switch _se.Sel.Name {
					case "Request":
						switch se.Sel.Name {
						case "FormValue":
							cur.Replace(&CallExpr{
								Fun: &Ident{Name: "string"},
								Args: []Expr{
									&CallExpr{
										Fun:  types.ExportCtxOp(ident.Name, "FormValue"),
										Args: call.Args,
									},
								},
							})
						case "FormFile":
							if as, ok := cur.Parent().(*AssignStmt); ok {
								as.Lhs = as.Lhs[1:]
							}
						case "UserAgent":
							cur.Replace(types.ExportUserAgent(ident.Name))
						}
					case "Writer":
						switch se.Sel.Name {
						case "Header":
							cur.Replace(types.ExportRespHeader(ident.Name))
						}
					case "FormFile":
						if as, ok := cur.Parent().(*AssignStmt); ok {
							as.Lhs = as.Lhs[1:]
						}
					case "GetRawData":
						if as, ok := cur.Parent().(*AssignStmt); ok {
							as.Lhs = as.Lhs[:1]
						}
					}
				}
			}
		}
	}
}
