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

func ReplaceRequestOp(se *SelectorExpr, cur *astutil.Cursor) {
	if _se, ok := se.X.(*SelectorExpr); ok {
		if ident, ok := _se.X.(*Ident); ok {
			if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
				if _se.Sel.Name == "Request" {
					switch se.Sel.Name {
					case "RequestURI":
						cur.Replace(types.ExportRequestURI(ident.Name))
					case "Method":
						cur.Replace(types.ExportReqMethod(ident.Name))
					case "Host":
						cur.Replace(types.ExportReqHost(ident.Name))
					}
				}
			}
		}

		if __se, ok := _se.X.(*SelectorExpr); ok {
			if ident, ok := __se.X.(*Ident); ok {
				if utils.CheckObjStarExpr(ident.Obj, "hzapp", "RequestContext") {
					if __se.Sel.Name == "Request" {
						switch _se.Sel.Name {
						case "URL":
							switch se.Sel.Name {
							case "Path":
								cur.Replace(types.ExportURIPath(ident.Name))
							case "String":
								cur.Replace(types.ExportURIString(ident.Name))
							case "RawQuery":
								cur.Replace(types.ExportURIQueryString(ident.Name))
							}
						case "Form":
							if se.Sel.Name == "Get" {
								cur.Replace(types.ExportCtxOp(ident.Name, "FormValue"))
							}
						case "Header":
							if se.Sel.Name == "Values" {
								cur.Replace(types.ExportReqHeaderGetAll(ident.Name))
							}
						}
					}
				}
			}
		}
	}
}
