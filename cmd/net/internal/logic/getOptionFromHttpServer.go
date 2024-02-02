package logic

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/config"

	"golang.org/x/tools/go/ast/astutil"
)

func GetOptionsFromHttpServer(cur *astutil.Cursor, cfg *config.Config) {
	block, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}

	// 遍历函数体的语句列表，找到 svr := http.Server {...} 这个赋值语句
	index := findHttpServerAssignment(block)
	if index == -1 {
		return
	}

	processHttpServerOptions(block, index, cfg)
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
func processHttpServerOptions(block *BlockStmt, index int, cfg *config.Config) {
	compLit := block.List[index].(*AssignStmt).Rhs[0].(*CompositeLit)

	for _, elt := range compLit.Elts {
		if kvExpr, ok := elt.(*KeyValueExpr); ok {
			key := kvExpr.Key.(*Ident).Name
			switch t := kvExpr.Value.(type) {
			case *BasicLit:
				handleBasicLitOption(cfg, key, t)
			case *CallExpr:
				handleCallerExprOption(cfg, key, t)
			case *BinaryExpr:
				handleBinaryExprOption(cfg, key, t)
			}
		}
	}

	block.List = append(block.List[:index], block.List[index+1:]...)
}

func handleCallerExprOption(cfg *config.Config, key string, t *CallExpr) {
	switch key {
	case "Addr":
		cfg.Addr = t.Args[0].(*BasicLit).Value
	}
}

// 处理 BasicLit 类型的配置选项
func handleBasicLitOption(cfg *config.Config, key string, lit *BasicLit) {
	switch key {
	case "Addr":
		cfg.Addr = lit.Value
	}
}

// 处理 BinaryExpr 类型的配置选项
func handleBinaryExprOption(cfg *config.Config, key string, expr *BinaryExpr) {
	switch key {
	case "IdleTimeout":
		cfg.IdleTimeout = expr.X.(*BasicLit).Value
	case "WriteTimeout":
		cfg.WriteTimeout = expr.X.(*BasicLit).Value
	case "ReadTimeout":
		cfg.ReadTimeout = expr.X.(*BasicLit).Value
	}
}
