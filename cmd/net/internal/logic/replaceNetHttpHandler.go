package logic

import (
	mapset "github.com/deckarep/golang-set/v2"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceNetHttpHandler(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	var (
		rwIndex   = -1
		rIndex    = -1
		paramList []*Field
	)
	paramSet := mapset.NewSet[string]()
	funcDecl, ok := cur.Node().(*FuncDecl)
	if !ok {
		return
	}
	funcType := funcDecl.Type
	paramList = funcType.Params.List
	for index, field := range paramList {
		switch t := field.Type.(type) {
		case *SelectorExpr:
			paramSet.Add(t.Sel.Name)
			if t.Sel.Name == "ResponseWriter" {
				rwIndex = index
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				paramSet.Add(selExpr.Sel.Name)
				if selExpr.Sel.Name == "Request" {
					rIndex = index
				}
			}
		}
	}

	if len(paramList) == 2 && paramSet.Contains("ResponseWriter", "Request") {
		ctx := &Field{
			Names: []*Ident{NewIdent("ctx")},
			Type:  NewIdent("context.Context"),
		}
		c := &Field{
			Names: []*Ident{NewIdent("c")},
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("app"),
					Sel: NewIdent("RequestContext"),
				},
			},
		}
		fields := []*Field{ctx, c}
		funcType.Params.List = fields
		funcSet.Add(funcDecl.Name.Name)
		return
	}

	if len(paramList) > 2 && paramSet.Contains("ResponseWriter", "Request") {
		var fields []*Field
		fields = append(fields, paramList[:rwIndex]...)
		fields = append(fields, &Field{
			Names: []*Ident{NewIdent("c")},
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("app"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rwIndex+2:]...)
		funcType.Params.List = fields
		funcSet.Add(funcDecl.Name.Name)
		return
	}

	if paramSet.Contains("ResponseWriter") {
		var fields []*Field
		fields = append(fields, paramList[:rwIndex]...)
		fields = append(fields, &Field{
			Names: []*Ident{NewIdent("c")},
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("app"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rwIndex+1:]...)
		funcType.Params.List = fields
		funcSet.Add(funcDecl.Name.Name)
		return
	}
	if paramSet.Contains("Request") {
		var fields []*Field
		fields = append(fields, paramList[:rIndex]...)
		fields = append(fields, &Field{
			Names: []*Ident{NewIdent("c")},
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("app"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rIndex+1:]...)
		paramList = fields
		funcSet.Add(funcDecl.Name.Name)
		return
	}
}
