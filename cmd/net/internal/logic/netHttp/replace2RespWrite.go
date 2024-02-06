package netHttp

import (
	. "go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

func ReplaceRespWrite(cur *astutil.Cursor) {
	stmt, ok := cur.Node().(*BlockStmt)
	if !ok || len(stmt.List) == 0 {
		return
	}
	var setStatusCodeInserted bool

	for i, s := range stmt.List {
		es, ok := s.(*ExprStmt)
		if !ok {
			continue
		}
		ce, ok := es.X.(*CallExpr)
		if !ok {
			continue
		}
		selExpr, ok := ce.Fun.(*SelectorExpr)
		if !ok || selExpr.Sel == nil {
			continue
		}

		// 检查是否已经插入了 c.SetStatusCode
		if selExpr.Sel.Name == "SetStatusCode" {
			setStatusCodeInserted = true
			continue
		}

		if selExpr.Sel.Name == "Write" {
			var _es *ExprStmt
			if !setStatusCodeInserted {
				_es = &ExprStmt{
					X: &CallExpr{
						Fun: &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("SetStatusCode"),
						},
						Args: []Expr{NewIdent("200")},
					},
				}
				stmt.List = append(stmt.List[:i], append([]Stmt{_es}, stmt.List[i:]...)...)
			}

			ce.Fun = &SelectorExpr{
				X: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("Response"),
				},
				Sel: NewIdent("SetBody"),
			}
		}
	}
}
