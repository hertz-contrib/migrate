package netHttp

import (
	. "go/ast"

	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqCookie(cur *astutil.Cursor) {
	var cookieName string
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}
	for _, stmt := range blockStmt.List {
		if cookieName != "" {
			switch ss := stmt.(type) {
			case *IfStmt:
				binaryExpr, ok := ss.Cond.(*BinaryExpr)
				if !ok {
					continue
				}
				selExpr, ok := binaryExpr.X.(*SelectorExpr)
				if ok {
					if selExpr.X.(*Ident).Name == cookieName && selExpr.Sel.Name == "Value" {
						binaryExpr.X = &CallExpr{
							Fun:  NewIdent("string"),
							Args: []Expr{NewIdent(cookieName)},
						}
					}
					continue
				}
			case *AssignStmt:
				if len(ss.Lhs) != 1 || len(ss.Rhs) != 1 {
					continue
				}
				selExpr, ok := ss.Rhs[0].(*SelectorExpr)
				if ok {
					if selExpr.X.(*Ident).Name == cookieName && selExpr.Sel.Name == "Value" {
						ss.Rhs[0] = &CallExpr{
							Fun:  NewIdent("string"),
							Args: []Expr{NewIdent(cookieName)},
						}
					}
					continue
				}
			}

		} else {
			assignStmt, ok := stmt.(*AssignStmt)
			if ok {
				if len(assignStmt.Lhs) != 2 || len(assignStmt.Rhs) != 1 {
					continue
				}
				callExpr, ok := assignStmt.Rhs[0].(*CallExpr)
				if !ok {
					continue
				}
				selExpr, ok := callExpr.Fun.(*SelectorExpr)
				if !ok {
					continue
				}
				if utils.CheckPtrStructName(selExpr, "Request") && selExpr.Sel.Name == "Cookie" {
					assignStmt.Lhs = assignStmt.Lhs[:len(assignStmt.Lhs)-1]
					cookieName = assignStmt.Lhs[0].(*Ident).Name
					selExpr.X = NewIdent("c")
				}
			}
		}
	}

}
