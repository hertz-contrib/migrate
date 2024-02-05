package netHttp

import (
	mapset "github.com/deckarep/golang-set/v2"
	. "go/ast"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
)

func CollectMsg(cur *astutil.Cursor, funcSet map[string][2]int, fset *token.FileSet, file *File) {
	var (
		paramList []*Field
	)
	paramSet := mapset.NewSet[string]()
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
				paramSet.Add(t.Sel.Name)
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					paramSet.Add(selExpr.Sel.Name)
				}
			}
		}
	}

	if paramSet.Contains("ResponseWriter") {
		funcSet[funcDecl.Name.Name] = collectParamIndex(funcType.Params.List)
	}

	if paramSet.Contains("Request") {
		funcSet[funcDecl.Name.Name] = collectParamIndex(funcType.Params.List)
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
