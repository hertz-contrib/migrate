package logic

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	. "go/ast"
	"go/token"
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
						astutil.AddImport(fset, file, "github.com/cloudwego/hertz/pkg/app/server")
						callExpr.Fun.(*SelectorExpr).X.(*Ident).Name = "server"
						callExpr.Fun.(*SelectorExpr).Sel.Name = "Default"
						config.Map["server"] = assign.Lhs[0].(*Ident).Name
						newOptions(callExpr, config.Map)
					}
				}
			}
		}
	}

	funcType, ok := cur.Node().(*FuncType)
	if !ok || funcType.Results == nil {
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
		if selExpr.Sel.Name == "ServeMux" {
			selExpr.X.(*Ident).Name = "server"
			selExpr.Sel.Name = "Hertz"
		}
	}
}
