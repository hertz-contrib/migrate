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

func ReplaceRespOp(se *SelectorExpr, cur *astutil.Cursor) {
	if ident, ok := se.X.(*Ident); ok {
		if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
			switch se.Sel.Name {
			case "Writer":
				cur.Replace(&CallExpr{
					Fun: &SelectorExpr{
						X: &SelectorExpr{
							X:   NewIdent(ident.Name),
							Sel: NewIdent("Response"),
						},
						Sel: NewIdent("BodyWriter"),
					},
				})
			}
		}
	}

	if _se, ok := se.X.(*SelectorExpr); ok {
		if ident, ok := _se.X.(*Ident); ok {
			if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
				if _se.Sel.Name == "Writer" {
					switch se.Sel.Name {
					case "Write":
						cur.Replace(types.ExportCtxOp(ident.Name, "Write"))
						return
					case "WriteString":
						cur.Replace(types.ExportCtxOp(ident.Name, "SetBodyString"))
					case "WriteHeader":
						cur.Replace(types.ExportCtxOp(ident.Name, "SetStatusCode"))
						return
					case "Status":
						cur.Replace(types.ExportStatusCode(ident.Name))
					}
				}
			}
		}
	}

	if call, ok := se.X.(*CallExpr); ok {
		if _se, ok := call.Fun.(*SelectorExpr); ok {
			if __se, ok := _se.X.(*SelectorExpr); ok {
				if ident, ok := __se.X.(*Ident); ok {
					if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
						if __se.Sel.Name == "Writer" && se.Sel.Name == "Values" {
							se.Sel.Name = "GetAll"
						}
					}
				}
			}
		}
	}
}
