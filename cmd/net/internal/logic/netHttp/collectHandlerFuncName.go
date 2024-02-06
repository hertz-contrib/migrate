package netHttp

import (
	mapset "github.com/deckarep/golang-set/v2"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func CollectHandlerFuncName(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	collectTmpFuncName(cur, funcSet)
	collectCommonFuncName(cur, funcSet)
}

func collectTmpFuncName(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	var (
		funcName  string
		paramList []*Field
	)
	funcLit, ok := cur.Node().(*FuncLit)
	if !ok {
		return
	}
	blockStmt := funcLit.Body

	for _, stmt := range blockStmt.List {
		as, ok := stmt.(*AssignStmt)
		if !ok {
			return
		}
		if len(as.Lhs) == 1 {
			funcName = as.Lhs[0].(*Ident).Name
		}
		if len(as.Rhs) == 1 {
			funcLit, ok := as.Rhs[0].(*FuncLit)
			if !ok {
				return
			}
			paramList = funcLit.Type.Params.List
			for _, field := range paramList {
				switch t := field.Type.(type) {
				case *SelectorExpr:
					if t.Sel.Name == "ResponseWriter" {
						funcSet.Add(funcName)
					}
				case *StarExpr:
					selExpr, ok := t.X.(*SelectorExpr)
					if ok {
						if selExpr.Sel.Name == "Request" {
							funcSet.Add(funcName)
						}
					}
				}
			}
		}
	}
}

func collectCommonFuncName(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	var (
		paramList []*Field
	)
	funcDecl, ok := cur.Node().(*FuncDecl)
	if !ok {
		return
	}
	funcType := funcDecl.Type
	paramList = funcType.Params.List

	for _, field := range paramList {
		switch t := field.Type.(type) {
		case *SelectorExpr:
			if t.Sel.Name == "ResponseWriter" {
				funcSet.Add(funcDecl.Name.Name)
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					funcSet.Add(funcDecl.Name.Name)

				}
			}
		}
	}
}
