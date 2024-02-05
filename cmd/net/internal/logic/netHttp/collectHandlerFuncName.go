package netHttp

import (
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func CollectHandlerFuncName(cur *astutil.Cursor, funcSet map[string][2]int) {
	collectTmpFuncName(cur, funcSet)
	collectCommonFuncName(cur, funcSet)
}

func collectTmpFuncName(cur *astutil.Cursor, funcSet map[string][2]int) {
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
						funcSet[funcName] = collectParamIndex(paramList)
					}
				case *StarExpr:
					selExpr, ok := t.X.(*SelectorExpr)
					if ok {
						if selExpr.Sel.Name == "Request" {
							funcSet[funcName] = collectParamIndex(paramList)
						}
					}
				}
			}
		}
	}
}

func collectCommonFuncName(cur *astutil.Cursor, funcSet map[string][2]int) {
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
				funcSet[funcDecl.Name.Name] = collectParamIndex(funcType.Params.List)
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					funcSet[funcDecl.Name.Name] = collectParamIndex(funcType.Params.List)
				}
			}
		}
	}
}

func collectParamIndex(fields []*Field) [2]int {
	var paramList [2]int
	for index, f := range fields {
		switch t := f.Type.(type) {
		case *SelectorExpr:
			if t.Sel.Name == "ResponseWriter" {
				paramList[0] = index
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					paramList[1] = index
				}
			}
		}
	}
	return paramList
}
