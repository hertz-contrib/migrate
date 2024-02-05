package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"strings"
)

func main() {
	src := `
package main

import (
	"net/http"
)

type Config struct{}

func (svc *Config) wjbool(r *http.Request) (error, bool) {
	getId := func(r *http.Request) (int64, error) {
		return 0, nil
	}
	return nil, false
}
`

	// 解析源代码
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// 使用 astutil 包来遍历 AST
	astutil.Apply(f, func(c *astutil.Cursor) bool {
		// 查找函数声明
		if fnDecl, ok := c.Node().(*ast.FuncDecl); ok {
			// 提取函数名和参数列表中的标识符
			funcName := fnDecl.Name.Name
			params := extractParams(fnDecl.Type.Params)

			// 打印找到的函数信息
			fmt.Printf("Found function: %s\n", funcName)
			fmt.Printf("Parameters: %s\n", params)
		}
		return true
	}, nil)

	fmt.Println("Search complete.")
}

// extractParams 提取参数列表中的标识符
func extractParams(params *ast.FieldList) string {
	var identifiers []string
	for _, field := range params.List {
		for _, ident := range field.Names {
			identifiers = append(identifiers, ident.Name)
		}
	}
	return strings.Join(identifiers, ", ")
}
