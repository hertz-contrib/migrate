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
)

func ReplaceStatisFS(call *CallExpr) {
	if se, ok := call.Fun.(*SelectorExpr); ok {
		switch se.Sel.Name {
		case "StaticFS":
			httpSystemExpr := call.Args[1]
			if _call, ok := httpSystemExpr.(*CallExpr); ok {
				if _se, ok := _call.Fun.(*SelectorExpr); ok {
					if utils.CheckSelPkgAndStruct(_se, "gin", "Dir") {
						root := _call.Args[0]
						listDir := _call.Args[1]
						call.Args[1] = types.ExportedAppFSPtr(root, listDir)
					}
				}
			}
		}
	}
}
