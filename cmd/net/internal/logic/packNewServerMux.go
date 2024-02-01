package logic

import (
	. "go/ast"
	"go/token"

	"github.com/hertz-contrib/migrate/cmd/net/internal/config"

	"golang.org/x/tools/go/ast/astutil"
)

func PackNewServeMux(cur *astutil.Cursor, fset *token.FileSet, file *File, cfg *config.Config) {
	assign, ok := cur.Node().(*AssignStmt)
	if ok {
		if len(assign.Lhs) == 1 && len(assign.Rhs) == 1 {
			if callExpr, ok := assign.Rhs[0].(*CallExpr); ok {
				if fun, ok := callExpr.Fun.(*SelectorExpr); ok {
					ident, ok := fun.X.(*Ident)
					if !ok {
						return
					}
					if ident.Name == "http" && fun.Sel.Name == "NewServeMux" {
						astutil.AddImport(fset, file, "github.com/cloudwego/hertz/pkg/app/server")
						callExpr.Fun.(*SelectorExpr).X.(*Ident).Name = "server"
						callExpr.Fun.(*SelectorExpr).Sel.Name = "Default"
						cfg.SrvVar = assign.Lhs[0].(*Ident).Name
						newOptions(callExpr, cfg)
					}
				}
			}
		}
	}
}

func newOptions(callExpr *CallExpr, opts *config.Config) {
	var args []Expr
	if opts.Addr != "" {
		optionFunc := addParamForOptionFunc("server", "WithHostPorts", opts.Addr, token.STRING)
		args = append(args, optionFunc)
	}
	if opts.IdleTimeout != "" {
		optionFunc := addParamForOptionFunc("server", "WithIdleTimeout", opts.IdleTimeout, token.INT)
		args = append(args, optionFunc)
	}
	if opts.WriteTimeout != "" {
		optionFunc := addParamForOptionFunc("server", "WithWriteTimeout", opts.WriteTimeout, token.INT)
		args = append(args, optionFunc)
	}
	if opts.ReadTimeout != "" {
		optionFunc := addParamForOptionFunc("server", "WithReadTimeout", opts.ReadTimeout, token.INT)
		args = append(args, optionFunc)
	}
	callExpr.Args = args
}
