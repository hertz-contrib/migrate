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

package netHttp

import (
	. "go/ast"
	"go/token"

	mapset "github.com/deckarep/golang-set/v2"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceNetHttpHandler(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	noWrapperLine(cur, fset, file)
	oneWrapperLine(cur, fset, file)
	twoWrapperLine(cur, fset, file)
	fieldListReplaceNetHttpHandler(cur, fset, file)
}

func fieldListReplaceNetHttpHandler(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	var (
		rwIndex   = -1
		rIndex    = -1
		paramList []*Field
		hasName   = true
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
				if field.Names == nil {
					hasName = false
				}
				paramSet.Add(t.Sel.Name)
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					rIndex = index
					if field.Names == nil {
						hasName = false
					}
					paramSet.Add(selExpr.Sel.Name)
				}
			}
		}
	}
	if paramSet.IsEmpty() {
		return
	}

	ctxName := []*Ident{NewIdent("ctx")}
	cName := []*Ident{NewIdent("c")}
	if !hasName {
		ctxName = nil
		cName = nil
	}

	if len(paramList) == 2 && paramSet.Contains("ResponseWriter", "Request") {
		ctx := &Field{
			Names: ctxName,
			Type:  NewIdent("context.Context"),
		}
		c := &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
					Sel: NewIdent("RequestContext"),
				},
			},
		}
		fields := []*Field{ctx, c}
		fieldList.List = fields

		return
	}

	if len(paramList) > 2 && paramSet.Contains("ResponseWriter", "Request") {
		var fields []*Field
		fields = append(fields, paramList[:rwIndex]...)
		fields = append(fields, &Field{
			Names: ctxName,
			Type:  NewIdent("context.Context"),
		})
		fields = append(fields, &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rwIndex+2:]...)
		fieldList.List = fields

		return
	}

	if paramSet.Contains("ResponseWriter") {
		var fields []*Field
		fields = append(fields, paramList[:rwIndex]...)
		fields = append(fields, &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rwIndex+1:]...)
		fieldList.List = fields
		return
	}
	if paramSet.Contains("Request") {
		var fields []*Field
		fields = append(fields, paramList[:rIndex]...)
		fields = append(fields, &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
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
		Names: ctxName,
		Type:  NewIdent("context.Context"),
	}
	c := &Field{
		Names: cName,
		Type: &StarExpr{
			X: &SelectorExpr{
				X:   NewIdent("hzapp"),
				Sel: NewIdent("RequestContext"),
			},
		},
	}
	fields := []*Field{ctx, c}
	fieldList.List = fields

}

func twoWrapperLine(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	var (
		paramList []*Field
		hasName   = true
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
	selExpr.X.(*Ident).Name = "hzapp"
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
					if field.Names == nil {
						hasName = false
					}
					paramSet.Add(t.Sel.Name)
				}
			case *StarExpr:
				selExpr, ok := t.X.(*SelectorExpr)
				if ok {
					if selExpr.Sel.Name == "Request" {
						if field.Names == nil {
							hasName = false
						}
						paramSet.Add(selExpr.Sel.Name)
					}
				}
			}
		}

		ctxName := []*Ident{NewIdent("ctx")}
		cName := []*Ident{NewIdent("c")}
		if !hasName {
			ctxName = nil
			cName = nil
		}

		if paramSet.Contains("ResponseWriter", "Request") {
			ctx := &Field{
				Names: ctxName,
				Type:  NewIdent("context.Context"),
			}
			c := &Field{
				Names: cName,
				Type: &StarExpr{
					X: &SelectorExpr{
						X:   NewIdent("hzapp"),
						Sel: NewIdent("RequestContext"),
					},
				},
			}
			fields := []*Field{ctx, c}
			funcLit.Type.Params.List = fields
			returnStmt.Results = []Expr{funcLit}
		}
	}
}

func oneWrapperLine(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	var (
		paramList []*Field
		hasName   = true
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
	selExpr.X.(*Ident).Name = "hzapp"
	for _, stmt := range funcDecl.Body.List {
		returnStmt, ok := stmt.(*ReturnStmt)
		if !ok || len(returnStmt.Results) != 1 {
			return
		}

		funcLit, ok := returnStmt.Results[0].(*FuncLit)
		if !ok || funcLit.Type == nil || funcLit.Type.Params == nil {
			return
		}
		paramList = funcLit.Type.Params.List
		if !ok || len(paramList) != 2 {
			return
		}
		for _, field := range paramList {
			switch t := field.Type.(type) {
			case *SelectorExpr:
				if t.Sel.Name == "ResponseWriter" {
					if field.Names == nil {
						hasName = false
					}
					paramSet.Add(t.Sel.Name)
				}
			case *StarExpr:
				selExpr, ok := t.X.(*SelectorExpr)
				if ok {
					if selExpr.Sel.Name == "Request" {
						if field.Names == nil {
							hasName = false
						}
						paramSet.Add(selExpr.Sel.Name)
					}
				}
			}
		}
		ctxName := []*Ident{NewIdent("ctx")}
		cName := []*Ident{NewIdent("c")}
		if !hasName {
			ctxName = nil
			cName = nil
		}
		if paramSet.Contains("ResponseWriter", "Request") {
			ctx := &Field{
				Names: ctxName,
				Type:  NewIdent("context.Context"),
			}
			c := &Field{
				Names: cName,
				Type: &StarExpr{
					X: &SelectorExpr{
						X:   NewIdent("hzapp"),
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

func noWrapperLine(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	var (
		rwIndex   = -1
		rIndex    = -1
		hasName   = true
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
				if field.Names == nil {
					hasName = false
				}
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					rIndex = index
					paramSet.Add(selExpr.Sel.Name)
					if field.Names == nil {
						hasName = false
					}
				}
			}
		}
	}
	ctxName := []*Ident{NewIdent("ctx")}
	cName := []*Ident{NewIdent("c")}
	if !hasName {
		ctxName = nil
		cName = nil
	}
	if len(paramList) == 2 && paramSet.Contains("ResponseWriter", "Request") {
		ctx := &Field{
			Names: ctxName,
			Type:  NewIdent("context.Context"),
		}
		c := &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
					Sel: NewIdent("RequestContext"),
				},
			},
		}
		fields := []*Field{ctx, c}
		funcType.Params.List = fields
		return
	}

	if len(paramList) > 2 && paramSet.Contains("ResponseWriter", "Request") {
		var fields []*Field
		fields = append(fields, paramList[:rwIndex]...)
		fields = append(fields, &Field{
			Names: ctxName,
			Type:  NewIdent("context.Context"),
		})
		fields = append(fields, &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rwIndex+2:]...)
		funcType.Params.List = fields
		return
	}

	if paramSet.Contains("ResponseWriter") {
		var fields []*Field
		fields = append(fields, paramList[:rwIndex]...)
		fields = append(fields, &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rwIndex+1:]...)
		funcType.Params.List = fields
		return
	}
	if paramSet.Contains("Request") {
		var fields []*Field
		fields = append(fields, paramList[:rIndex]...)
		fields = append(fields, &Field{
			Names: cName,
			Type: &StarExpr{
				X: &SelectorExpr{
					X:   NewIdent("hzapp"),
					Sel: NewIdent("RequestContext"),
				},
			},
		})
		fields = append(fields, paramList[rIndex+1:]...)
		funcType.Params.List = fields
		return
	}
}
