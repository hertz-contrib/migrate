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
