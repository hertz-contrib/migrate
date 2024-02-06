package netHttp

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceRequestURI(cur *astutil.Cursor) {
	replaceAssignStmtRequestURI(cur)
	replaceIfStmtRequestURI(cur)
}

func replaceAssignStmtRequestURI(cur *astutil.Cursor) {
	assignStmt, ok := cur.Node().(*AssignStmt)
	if !ok {
		return
	}
	if len(assignStmt.Rhs) == 1 {
		selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
		if !ok {
			return
		}
		if selExpr.Sel.Name == "RequestURI" {
			if utils.CheckPtrStructName(selExpr, "Request") {
				newExpr := &CallExpr{
					Fun: &Ident{Name: "string"},
					Args: []Expr{
						&CallExpr{
							Fun: &SelectorExpr{
								X: &SelectorExpr{
									X:   &Ident{Name: "c"},
									Sel: &Ident{Name: "Request"},
								},
								Sel: &Ident{Name: "RequestURI"},
							},
						},
					},
				}
				assignStmt.Rhs[0] = newExpr
			}
		}
	}
}

func replaceIfStmtRequestURI(cur *astutil.Cursor) {
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
	if utils.CheckPtrStructName(se, "Request") && se.Sel.Name == "RequestURI" {
		be := &BinaryExpr{
			X: &CallExpr{
				Fun: &Ident{Name: "string"},
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X: &SelectorExpr{
								X:   &Ident{Name: "c"},
								Sel: &Ident{Name: "Request"},
							},
							Sel: &Ident{Name: "RequestURI"},
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
