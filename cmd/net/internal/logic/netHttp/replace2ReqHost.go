package netHttp

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

// ReplaceReqHost replaces r.Host with string(c.Host)
func ReplaceReqHost(cur *astutil.Cursor) {
	replaceAssignStmtReqHost(cur)
	replaceIfStmtReqHost(cur)
	replaceParamListReqHost(cur)
}

func replaceParamListReqHost(cur *astutil.Cursor) {

}

// replaceAssignStmtReqHost replaces r.Host with string(c.Host) in AssignStmt
func replaceAssignStmtReqHost(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok || len(assignStmt.Rhs) != 1 {
		return
	}

	selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
	if !ok || selExpr.Sel.Name != "Host" {
		return
	}

	if utils.CheckPtrStructName(selExpr, "Request") {
		// Create a new expression
		newExpr := &CallExpr{
			Fun: &Ident{Name: "string"},
			Args: []Expr{
				&CallExpr{
					Fun: &SelectorExpr{
						X:   &SelectorExpr{X: &Ident{Name: "c"}, Sel: &Ident{Name: "Request"}},
						Sel: &Ident{Name: "Host"},
					},
				},
			},
		}

		// Replace the right-hand side of the assignment statement
		assignStmt.Rhs = []Expr{newExpr}
	}
}

// replaceIfStmtReqHost replaces r.Host with string(c.Host) in IfStmt
func replaceIfStmtReqHost(cur *astutil.Cursor) {
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

	if utils.CheckPtrStructName(se, "Request") && se.Sel.Name == "Host" {
		be := &BinaryExpr{
			X: &CallExpr{
				Fun: &Ident{Name: "string"},
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X:   &SelectorExpr{X: &Ident{Name: "c"}, Sel: &Ident{Name: "Request"}},
							Sel: &Ident{Name: "Host"},
						},
					},
				},
			},
			Op: binaryExpr.Op,
			Y:  binaryExpr.Y,
		}

		ifStmt.Cond = be
	}
}
