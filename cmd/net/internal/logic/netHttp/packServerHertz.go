package netHttp

import (
	. "go/ast"
	"go/token"

	"github.com/hertz-contrib/migrate/cmd/net/internal/global"

	"golang.org/x/tools/go/ast/astutil"
)

func PackServerHertz(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	assign, ok := cur.Node().(*AssignStmt)
	if ok {
		if len(assign.Lhs) == 1 && len(assign.Rhs) == 1 {
			if callExpr, ok := assign.Rhs[0].(*CallExpr); ok {
				if fun, ok := callExpr.Fun.(*SelectorExpr); ok {
					ident, ok := fun.X.(*Ident)
					if !ok {
						return
					}
					if ident.Name == "http" && fun.Sel.Name == "NewServeMux" {
						callExpr.Fun.(*SelectorExpr).X.(*Ident).Name = "server"
						callExpr.Fun.(*SelectorExpr).Sel.Name = "Default"
						global.Map["server"] = assign.Lhs[0].(*Ident).Name
						AddOptionsForServer(callExpr, global.Map)
					}
				}
			}
		}
	}

	funcType, ok := cur.Node().(*FuncType)
	if ok {
		if funcType.Results == nil {
			return
		}
		if len(funcType.Results.List) == 1 {
			starExpr, ok := funcType.Results.List[0].Type.(*StarExpr)
			if !ok {
				return
			}
			selExpr, ok := starExpr.X.(*SelectorExpr)
			if !ok {
				return
			}
			if selExpr.Sel.Name == "ServeMux" || selExpr.Sel.Name == "Mux" {
				selExpr.X.(*Ident).Name = "server"
				selExpr.Sel.Name = "Hertz"
			}
		}
	}

	fieldList, ok := cur.Node().(*FieldList)
	if ok {
		for _, field := range fieldList.List {
			starExpr, ok := field.Type.(*StarExpr)
			if !ok {
				continue
			}
			selExpr, ok := starExpr.X.(*SelectorExpr)
			if !ok {
				continue
			}
			if selExpr.Sel.Name == "ServeMux" || selExpr.Sel.Name == "Mux" {
				selExpr.X.(*Ident).Name = "server"
				selExpr.Sel.Name = "Hertz"
			}
		}
	}
}
