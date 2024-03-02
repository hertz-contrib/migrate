package netHttp

import (
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/types"
	. "go/ast"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceHttpOp(call *CallExpr, cur *astutil.Cursor) {
	if sel, ok := call.Fun.(*SelectorExpr); ok {
		if ident, ok := sel.X.(*Ident); ok {
			if ident.Name == "http" {
				switch sel.Sel.Name {
				case "NotFound":
					cur.Replace(types.CallNotFound)
				case "Redirect":
					// http.Redirect(w http.ResponseWriter, r *http.Request, url string, code int)
					redirectURIString := call.Args[2]
					cur.Replace(types.ExportCallRedirect("c", call.Args[3], redirectURIString))
				case "Error":
					// remove w http.ResponseWriter
					callArgs := call.Args[1:]
					lit, ok := callArgs[0].(*BasicLit)
					if ok && lit.Kind == token.STRING {
						if lit.Value == "\"\"" {
							cur.Replace(&CallExpr{
								Fun:  types.ExportCtxOp("c", "AbortWithStatus"),
								Args: callArgs[1:],
							})
						} else {
							cur.Replace(&CallExpr{
								Fun:  types.ExportCtxOp("c", "AbortWithMsg"),
								Args: callArgs,
							})
						}
					}
				}
			}
		}
	}
}
