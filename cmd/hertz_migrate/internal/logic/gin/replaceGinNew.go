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

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceGinNew(call *CallExpr, c *astutil.Cursor) {
	if sel, ok := call.Fun.(*SelectorExpr); ok {
		if utils.CheckSelPkgAndStruct(sel, "gin", "New") ||
			utils.CheckSelPkgAndStruct(sel, "gin", "Default") {
			args := internal.Options
			if as, ok := c.Parent().(*AssignStmt); ok {
				if ident, ok := as.Lhs[0].(*Ident); ok {
					internal.ServerName = ident.Name
				}
			}

			var initTypeIdent *Ident
			if utils.CheckSelPkgAndStruct(sel, "gin", "New") {
				initTypeIdent = NewIdent("New")
			} else {
				initTypeIdent = NewIdent("Default")
			}

			c.Replace(&CallExpr{
				Fun: &SelectorExpr{
					X:   NewIdent("hzserver"),
					Sel: initTypeIdent,
				},
				Args: args,
			})
		}
	}
}
