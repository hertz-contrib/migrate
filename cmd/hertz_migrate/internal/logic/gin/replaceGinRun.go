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
	"go/ast"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"
)

func ReplaceGinRun(node *ast.CallExpr) {
	if sel, ok := node.Fun.(*ast.SelectorExpr); ok {
		if ident, ok := sel.X.(*ast.Ident); ok {
			if utils.CheckObjSelExpr(ident.Obj, "hzserver", "Default") ||
				utils.CheckObjStarExpr(ident.Obj, "hzserver", "Hertz") {
				if sel.Sel.Name == "Run" {
					sel.Sel.Name = "Spin"
					node.Args = []ast.Expr{}
				}
			}
			if utils.CheckObjSelExpr(ident.Obj, "http", "Server") {
				if sel.Sel.Name == "ListenAndServe" {
					ident.Name = internal.ServerName
					sel.Sel.Name = "Run"
				}
			}
		}
	}
}
