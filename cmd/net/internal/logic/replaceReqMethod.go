package logic

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceReqMethod(cur *astutil.Cursor) {
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
			if ident, ok := selExpr.X.(*Ident); ok {
				if field, ok := ident.Obj.Decl.(*Field); ok {
					starExpr, ok := field.Type.(*StarExpr)
					if !ok {
						return
					}
					selExpr, ok := starExpr.X.(*SelectorExpr)
					if !ok {
						return
					}
					if selExpr.Sel.Name == "Request" {
						newExpr := &CallExpr{
							Fun: &Ident{Name: "string"},
							Args: []Expr{
								&CallExpr{
									Fun: &SelectorExpr{
										X: &SelectorExpr{
											X:   &Ident{Name: "c"},
											Sel: &Ident{Name: "Request"},
										},
										Sel: &Ident{Name: "Method"},
									},
								},
							},
						}
						assignStmt.Rhs[0] = newExpr
					}
				}
			}
		}
	}
	//if fn, ok := cur.Node().(*FuncDecl); ok {
	//	for _, stmt := range fn.Body.List {
	//		if assignStmt, ok := stmt.(*AssignStmt); ok {
	//			for _, expr := range assignStmt.Rhs {
	//				if ident, ok := expr.(*Ident); ok && ident.Name == "Method" {
	//					// 构造新的表达式 string(c.Request.Method())
	//					newExpr := &CallExpr{
	//						Fun: &Ident{Name: "string"},
	//						Args: []Expr{
	//							&CallExpr{
	//								Fun: &SelectorExpr{
	//									X: &SelectorExpr{
	//										X:   &Ident{Name: "c"},
	//										Sel: &Ident{Name: "Request"},
	//									},
	//									Sel: &Ident{Name: "Method"},
	//								},
	//								Args: nil,
	//							},
	//						},
	//					}
	//					cur.Replace(newExpr)
	//				}
	//			}
	//		}
	//	}
	//}
}
