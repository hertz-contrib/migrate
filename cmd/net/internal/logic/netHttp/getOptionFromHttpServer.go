package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/config"

	"golang.org/x/tools/go/ast/astutil"
)

func GetOptionsFromHttpServer(cur *astutil.Cursor) {
	block, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}

	// 遍历函数体的语句列表，找到 svr := http.Server {...} 这个赋值语句
	index := findHttpServerAssignment(block)
	if index == -1 {
		return
	}

	processHttpServerOptions(block, index)
}

// 找到 http.Server 赋值语句的索引
func findHttpServerAssignment(block *BlockStmt) int {
	for i, stmt := range block.List {
		assign, ok := stmt.(*AssignStmt)
		if !ok || len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
			continue
		}

		_, ok = assign.Lhs[0].(*Ident)
		if !ok {
			continue
		}

		compLit, ok := assign.Rhs[0].(*CompositeLit)
		if !ok {
			continue
		}

		selExpr, ok := compLit.Type.(*SelectorExpr)
		if !ok || selExpr.X.(*Ident).Name != "http" || selExpr.Sel.Name != "Server" {
			continue
		}

		return i
	}

	return -1
}

// 处理 http.Server 赋值语句，更新配置并移除该语句
func processHttpServerOptions(block *BlockStmt, index int) {
	compLit := block.List[index].(*AssignStmt).Rhs[0].(*CompositeLit)

	for _, elt := range compLit.Elts {
		if kvExpr, ok := elt.(*KeyValueExpr); ok {
			key := kvExpr.Key.(*Ident).Name
			config.Map[key] = kvExpr.Value
		}
	}
	// Remove the http.Server assignment statement
	//block.List = append(block.List[:index], block.List[index+1:]...)
}
