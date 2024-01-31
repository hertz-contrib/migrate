package utils

import (
	. "go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

var AliasMap map[string]string

func GetAllAliasForPackage(fset *token.FileSet, file *File) (m map[string]string) {
	m = make(map[string]string)
	imports := astutil.Imports(fset, file)
	for _, group := range imports {
		for _, spec := range group {
			packageAlias := spec.Name.String()
			if packageAlias == "<nil>" {
				continue
			}
			packageName, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				continue
			}
			m[packageName] = packageAlias
		}
	}
	return
}

func IsHttpResponseWriter(t *Field) bool {
	packageName := "http"
	expr, ok := t.Type.(*SelectorExpr)
	if !ok {
		return false
	}
	v, ok := AliasMap["net/http"]
	if ok {
		packageName = v
	}
	if expr.X.(*Ident).Name == packageName && expr.Sel.Name == "ResponseWriter" {
		return true
	}
	return false
}

func IsHttpRequest(t *Field) bool {
	packageName := "http"
	expr, ok := t.Type.(*StarExpr)
	if !ok {
		return false
	}
	selectorExpr, ok := expr.X.(*SelectorExpr)
	if !ok {
		return false
	}
	v, ok := AliasMap["net/http"]
	if ok {
		packageName = v
	}
	if selectorExpr.X.(*Ident).Name == packageName && selectorExpr.Sel.Name == "Request" {
		return true
	}
	return false
}

func PackHandleFunc(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	fieldList, ok := cur.Node().(*FieldList)
	if ok {
		if len(fieldList.List) == 2 {
			if IsHttpResponseWriter(fieldList.List[0]) && IsHttpRequest(fieldList.List[1]) {
				astutil.AddImport(fset, file, "context")
				astutil.AddImport(fset, file, "github.com/cloudwego/hertz/pkg/app")
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
				// replace the old parameters with the new ones
				cur.Replace(&FieldList{
					List: []*Field{ctx, c},
				})
			}
		}
	}
}

func PackFprintf(cur *astutil.Cursor, fset *token.FileSet, file *File) {
	var isFmt bool
	var isFprintf bool
	var isResponseWriter bool
	var callExpr *CallExpr
	var ok bool
	exprStmt, ok := cur.Node().(*ExprStmt)
	if ok {
		callExpr, ok = exprStmt.X.(*CallExpr)
		if ok {
			if fun, ok := callExpr.Fun.(*SelectorExpr); ok {
				if ok {
					ident, ok := fun.X.(*Ident)
					if ok {
						isFmt = ident.Name == "fmt"
					}
					isFprintf = fun.Sel.Name == "Fprintf"
				}
			}
		}
	}
	if isFmt && isFprintf {
		// w
		ident, ok := callExpr.Args[0].(*Ident)
		if ok {
			field, ok := ident.Obj.Decl.(*Field)
			if ok {
				expr, ok := field.Type.(*SelectorExpr)
				if ok {
					isResponseWriter = expr.Sel.Name == "ResponseWriter"
				}
			}
		}
	}
	if isResponseWriter {
		callExpr.Fun.(*SelectorExpr).X.(*Ident).Name = "c"          // 修改接收者为c
		callExpr.Fun.(*SelectorExpr).Sel.Name = "String"            // 修改方法名为String
		callExpr.Args[0] = &BasicLit{Kind: token.INT, Value: "200"} // 修改第一个参数为200
	}
}
