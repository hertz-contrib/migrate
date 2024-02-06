package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqURLPath(cur *astutil.Cursor) {
	replaceBlockStmtReqURLPath(cur)
	replaceExprStmtReqURLPath(cur)
}

func replaceExprStmtReqURLPath(cur *astutil.Cursor) {
	callExpr, ok := cur.Node().(*CallExpr)
	if !ok {
		return
	}
	for i, arg := range callExpr.Args {
		arg, ok := arg.(*SelectorExpr)
		if !ok || arg.Sel.Name != "Path" {
			continue
		}
		selExpr, ok := arg.X.(*SelectorExpr)
		if !ok {
			continue
		}
		if utils.CheckPtrStructName(selExpr, "Request") {
			callExpr.Args[i] = &CallExpr{
				Fun: &Ident{Name: "string"},
				Args: []Expr{
					&CallExpr{
						Fun: &SelectorExpr{
							X: &CallExpr{
								Fun: &SelectorExpr{
									X:   &Ident{Name: "c"},
									Sel: &Ident{Name: "URI"},
								},
							},
							Sel: &Ident{Name: "Path"},
						},
					},
				},
			}
		}
	}
}

func replaceBlockStmtReqURLPath(cur *astutil.Cursor) {
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	for i, stmt := range blockStmt.List {
		assignStmt, ok := stmt.(*AssignStmt)
		if !ok {
			continue
		}
		if len(assignStmt.Rhs) == 1 {
			selExpr, ok := assignStmt.Rhs[0].(*SelectorExpr)
			if !ok {
				continue
			}
			if selExpr.Sel.Name == "Path" {
				assignStmt.Rhs[0] = &CallExpr{
					Fun: &Ident{Name: "string"},
					Args: []Expr{
						&CallExpr{
							Fun: &SelectorExpr{
								X: &CallExpr{
									Fun: &SelectorExpr{
										X:   &Ident{Name: "c"},
										Sel: &Ident{Name: "URI"},
									},
								},
								Sel: &Ident{Name: "Path"},
							},
						},
					},
				}
				blockStmt.List[i] = assignStmt
			}
		}
	}
}
