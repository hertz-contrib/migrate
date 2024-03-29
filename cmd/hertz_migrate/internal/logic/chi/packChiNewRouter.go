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

package chi

import (
	. "go/ast"

	nethttp "github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/logic/netHttp"

	"golang.org/x/tools/go/ast/astutil"
)

func PackChiNewRouter(cur *astutil.Cursor, globalMap map[string]any) {
	stmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(stmt.Lhs) != 1 || len(stmt.Rhs) != 1 {
		return
	}
	callExpr, ok := stmt.Rhs[0].(*CallExpr)
	if !ok {
		return
	}
	selExpr, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return
	}
	if selExpr.Sel.Name == "NewRouter" {
		callExpr.Fun = &SelectorExpr{
			X:   NewIdent("hzserver"),
			Sel: NewIdent("Default"),
		}
		globalMap["serverName"] = stmt.Lhs[0].(*Ident).Name
		nethttp.AddOptionsForServer(callExpr, globalMap)
	}
}
