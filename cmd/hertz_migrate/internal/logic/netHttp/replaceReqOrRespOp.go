package netHttp

import (
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/types"
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqOrRespOp(call *CallExpr, cur *astutil.Cursor) {
	if sel, ok := call.Fun.(*SelectorExpr); ok {
		if ident, ok := sel.X.(*Ident); ok {
			if utils.CheckObjStarExpr(ident.Obj, "http", "Request") {
				ident.Name = "c"
				switch sel.Sel.Name {
				case "FormFile":
					if as, ok := cur.Parent().(*AssignStmt); ok {
						as.Lhs = as.Lhs[1:]
					}
				case "Cookie":
					if as, ok := cur.Parent().(*AssignStmt); ok {
						as.Lhs = as.Lhs[:1]
					}
					cur.Replace(types.ExportCtxCookie("c", call.Args))
				}
			}
		}
	}
}
