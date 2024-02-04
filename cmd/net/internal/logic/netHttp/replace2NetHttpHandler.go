package netHttp

import (
	. "go/ast"

	mapset "github.com/deckarep/golang-set/v2"

	"golang.org/x/tools/go/ast/astutil"
)

func Replace2NetHttpHandler(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	noWrapperLine(cur, funcSet)
	oneWrapperLine(cur)
	twoWrapperLine(cur)
	inlineOneWrapperLine(cur)
}

func inlineOneWrapperLine(cur *astutil.Cursor) {
	var (
		rwIndex   = -1
		rIndex    = -1
		paramList []*Field
	)
	paramSet := mapset.NewSet[string]()
	fieldList, ok := cur.Node().(*FieldList)
	if !ok {
		return
	}
	paramList = fieldList.List
	for index, field := range paramList {
		switch t := field.Type.(type) {
		case *SelectorExpr:
			if t.Sel.Name == "ResponseWriter" {
				rwIndex = index
				paramSet.Add(t.Sel.Name)
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					rIndex = index
					paramSet.Add(selExpr.Sel.Name)
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
		paramList = fields
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
		paramList = fields
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
		paramList = fields
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
		fieldList.List = fields
		return
	}

	if !ok || len(fieldList.List) != 2 {
		return
	}
	selExpr, ok := fieldList.List[0].Type.(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "ResponseWriter" {
		return
	}
	starExpr, ok := fieldList.List[1].Type.(*StarExpr)
	if !ok || starExpr.X.(*SelectorExpr).Sel.Name != "Request" {
		return
	}
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
	fieldList.List = fields
}

func twoWrapperLine(cur *astutil.Cursor) {
	var (
		paramList []*Field
	)
	paramSet := mapset.NewSet[string]()

	funcDecl, ok := cur.Node().(*FuncDecl)
	if !ok || funcDecl.Type == nil || funcDecl.Type.Results == nil || len(funcDecl.Type.Results.List) != 1 {
		return
	}
	selExpr, ok := funcDecl.Type.Results.List[0].Type.(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "Handler" || selExpr.X.(*Ident).Name != "http" {
		return
	}
	selExpr.X.(*Ident).Name = "app"
	selExpr.Sel.Name = "HandlerFunc"

	for _, stmt := range funcDecl.Body.List {
		returnStmt, ok := stmt.(*ReturnStmt)
		if !ok {
			continue
		}

		ce, ok := returnStmt.Results[0].(*CallExpr)
		if !ok || len(ce.Args) != 1 {
			continue
		}

		funcLit, ok := ce.Args[0].(*FuncLit)
		if !ok || funcLit.Type == nil || funcLit.Type.Params == nil || len(funcLit.Type.Params.List) != 2 {
			continue
		}
		paramList = funcLit.Type.Params.List

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

		if paramSet.Contains("ResponseWriter", "Request") {
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
			funcLit.Type.Params.List = fields
			returnStmt.Results = []Expr{funcLit}
			return
		}

	}
}

func oneWrapperLine(cur *astutil.Cursor) {
	var (
		paramList []*Field
	)
	paramSet := mapset.NewSet[string]()

	funcDecl, ok := cur.Node().(*FuncDecl)
	if !ok || funcDecl.Type == nil || funcDecl.Type.Results == nil || len(funcDecl.Type.Results.List) != 1 {
		return
	}
	selExpr, ok := funcDecl.Type.Results.List[0].Type.(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "HandlerFunc" || selExpr.X.(*Ident).Name != "http" {
		return
	}
	selExpr.X.(*Ident).Name = "app"

	for _, stmt := range funcDecl.Body.List {
		returnStmt, ok := stmt.(*ReturnStmt)
		if !ok || len(returnStmt.Results) != 1 {
			return
		}

		funcLit, ok := returnStmt.Results[0].(*FuncLit)
		paramList = funcLit.Type.Params.List
		if !ok || len(paramList) != 2 {
			return
		}
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

		if paramSet.Contains("ResponseWriter", "Request") {
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
			funcLit.Type.Params.List = fields
			return
		}
	}
}

func noWrapperLine(cur *astutil.Cursor, funcSet mapset.Set[string]) {
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
			if t.Sel.Name == "ResponseWriter" {
				rwIndex = index
				paramSet.Add(t.Sel.Name)

			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					rIndex = index
					paramSet.Add(selExpr.Sel.Name)
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
