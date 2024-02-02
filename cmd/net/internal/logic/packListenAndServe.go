package logic

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	"golang.org/x/tools/go/ast/astutil"
)

func PackListenAndServe(cur *astutil.Cursor, cfg *config.Config) {
	selExpr, ok := cur.Node().(*SelectorExpr)
	if ok {
		if selExpr.Sel == nil {
			return
		}
		if selExpr.Sel.Name == "ListenAndServe" {
			selExpr.X.(*Ident).Name = cfg.ServerVar
			selExpr.Sel.Name = "Spin"
		}
	}
}
