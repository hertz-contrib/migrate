package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/config"
	"golang.org/x/tools/go/ast/astutil"
)

func PackListenAndServe(cur *astutil.Cursor) {
	selExpr, ok := cur.Node().(*SelectorExpr)
	if ok {
		if selExpr.Sel == nil {
			return
		}
		if selExpr.Sel.Name == "ListenAndServe" {
			v, ok := config.Map["server"]
			if ok {
				selExpr.X.(*Ident).Name = v.(string)
				selExpr.Sel.Name = "Spin"
			}
		}
	}
}
