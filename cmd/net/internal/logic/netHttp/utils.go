package netHttp

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
