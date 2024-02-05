package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqHeader(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(assignStmt.Rhs) != 1 {
		return
	}

	selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "Header" {
		return
	}

	if utils.CheckPtrStructName(selExpr, "Request") {
		callExpr := &SelectorExpr{
			X: &SelectorExpr{
				X:   NewIdent("c"),
				Sel: NewIdent("Request"),
			},
			Sel: NewIdent("Header"),
		}
		// Replace the right-hand side of the assignment statement
		assignStmt.Rhs[0] = callExpr
	}
}
