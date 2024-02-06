package netHttp

import (
	. "go/ast"
	"sync"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"

	"golang.org/x/tools/go/ast/astutil"
)

var stringMethodExpr *CallExpr

func ReplaceReqMethod(cur *astutil.Cursor) {
	var once sync.Once
	once.Do(func() {
		stringMethodExpr = &CallExpr{
			Fun: &Ident{Name: "string"},
			Args: []Expr{
				&CallExpr{
					Fun: &SelectorExpr{
						X:   &Ident{Name: "c"},
						Sel: &Ident{Name: "Method"},
					},
				},
			},
		}
	})
	replaceAssignStmtReqMethod(cur)
	replaceIfStmtReqMethod(cur)
	replaceSwitchStmtReqMethod(cur)
}

func replaceSwitchStmtReqMethod(cur *astutil.Cursor) {
	switchStmt, ok := cur.Node().(*SwitchStmt)
	if !ok {
		return
	}
	selExpr, ok := switchStmt.Tag.(*SelectorExpr)
	if !ok {
		return
	}
	if utils.CheckPtrStructName(selExpr, "Request") && selExpr.Sel.Name == "Method" {
		switchStmt.Tag = stringMethodExpr
	}
}

func replaceAssignStmtReqMethod(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok {
		return
	}
	if len(assignStmt.Rhs) == 1 {
		selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "Method" {
			if utils.CheckPtrStructName(selExpr, "Request") {
				assignStmt.Rhs[0] = stringMethodExpr
			}
		}
	}
}
func replaceIfStmtReqMethod(cur *astutil.Cursor) {
	ifStmt, ok := cur.Node().(*IfStmt)
	if !ok {
		return
	}
	binaryExpr, ok := ifStmt.Cond.(*BinaryExpr)
	if !ok {
		return
	}
	se, ok := binaryExpr.X.(*SelectorExpr)
	if !ok {
		return
	}
	if utils.CheckPtrStructName(se, "Request") {
		if se.Sel.Name == "Method" {
			be := &BinaryExpr{
				X:  stringMethodExpr,
				Y:  binaryExpr.Y,
				Op: binaryExpr.Op,
			}
			ifStmt.Cond = be
		}
	}
}
