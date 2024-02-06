package netHttp

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func PackType2AppHandlerFunc(cur *astutil.Cursor) {
	packReturnStmt2AppHandlerFunc(cur)
}

func packReturnStmt2AppHandlerFunc(cur *astutil.Cursor) {
	funcDecl, ok := cur.Node().(*FuncDecl)
	if !ok {
		return
	}
	funcType := funcDecl.Type
	if funcType.Results == nil {
		return
	}
	fields := funcType.Results.List
	if len(fields) == 1 {
		ft, ok := fields[0].Type.(*FuncType)
		if !ok {
			return
		}
		if len(ft.Params.List) != 2 {
			return
		}
		if ft.Params.List[0].Names[0].Name == "ctx" && ft.Params.List[1].Names[0].Name == "c" {
			funcType.Results.List = []*Field{
				{
					Type: &SelectorExpr{
						X:   NewIdent("app"),
						Sel: NewIdent("HandlerFunc"),
					},
				},
			}
		}
	}
}
