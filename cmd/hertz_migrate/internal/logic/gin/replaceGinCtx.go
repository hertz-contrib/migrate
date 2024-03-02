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

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/types"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"
)

func ReplaceGinCtx(fields *ast.FieldList) {
	ctxIndex := -1
	if fields.List == nil {
		return
	}

	for index, field := range fields.List {
		if star, ok := field.Type.(*ast.StarExpr); ok {
			if sel, ok := star.X.(*ast.SelectorExpr); ok {
				if utils.CheckSelPkgAndStruct(sel, "gin", "Context") {
					ctxIndex = index
					field.Type = types.StarCtx
					break
				}
			}
		}
	}

	if ctxIndex >= 0 {
		slice := make([]*ast.Field, len(fields.List)+1)
		slice = append(slice, fields.List[:ctxIndex]...)
		slice = append(slice, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent("_ctx"),
			},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("context"),
				Sel: ast.NewIdent("Context"),
			},
		})

		slice = append(slice, fields.List[ctxIndex:]...)
		var fieldList []*ast.Field
		for _, field := range slice {
			if field != nil {
				fieldList = append(fieldList, field)
			}
		}
		fields.List = fieldList
	}
}
