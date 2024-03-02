package netHttp

import (
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/types"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/utils"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceRequestOp(sel *SelectorExpr, cur *astutil.Cursor) {
	if ident, ok := sel.X.(*Ident); ok {
		if utils.CheckObjStarExpr(ident.Obj, "http", "Request") {
			switch sel.Sel.Name {
			case "RequestURI":
				cur.Replace(types.ExportRequestURI("c"))
			case "Method":
				cur.Replace(types.ExportReqMethod("c"))
			case "Host":
				cur.Replace(types.ExportReqHost("c"))
			case "ContentLength":
				cur.Replace(types.CallContentLength)
			case "RemoteAddr":
				cur.Replace(types.CallRemoteAddr)
			}
		}
	}

	if _sel, ok := sel.X.(*SelectorExpr); ok {
		if ident, ok := _sel.X.(*Ident); ok {
			if utils.CheckObjStarExpr(ident.Obj, "http", "Request") {
				if _sel.Sel.Name == "URL" {
					_sel.Sel.Name = "URI"
					switch sel.Sel.Name {
					case "Path":
						cur.Replace(types.ExportURIPath("c"))
					case "String":
						cur.Replace(types.ExportURIString("c"))
					}
				}
			}
		}
	}
}
