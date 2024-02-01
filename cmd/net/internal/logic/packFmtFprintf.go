package logic

import (
	. "go/ast"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
)

func PackFprintf(cur *astutil.Cursor) {
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}

	for i, stmt := range blockStmt.List {
		if !isTargetStmt(stmt) {
			continue
		}

		ident, ok := getSelectorIdent(stmt)
		if !ok {
			continue
		}

		field, ok := getField(ident)
		if !ok {
			continue
		}

		expr, ok := getTypeSelectorExpr(field)
		if !ok {
			continue
		}

		if !isResponseWriterType(expr) {
			continue
		}

		updateStmts(blockStmt, i, stmt)
		return
	}
}

func isTargetStmt(stmt Stmt) bool {
	exprStmt, ok := stmt.(*ExprStmt)
	if !ok {
		return false
	}

	callExpr, ok := exprStmt.X.(*CallExpr)
	if !ok {
		return false
	}

	fprintf, ok := callExpr.Fun.(*SelectorExpr)
	if !ok {
		return false
	}

	ident, ok := fprintf.X.(*Ident)
	if ok {
		return ident.Name == "fmt" && fprintf.Sel.Name == "Fprintf"
	}
	return false
}

func getSelectorIdent(stmt Stmt) (*Ident, bool) {
	exprStmt := stmt.(*ExprStmt)
	callExpr := exprStmt.X.(*CallExpr)
	ident, ok := callExpr.Args[0].(*Ident)
	return ident, ok
}

func getField(ident *Ident) (*Field, bool) {
	field, ok := ident.Obj.Decl.(*Field)
	return field, ok
}

func getTypeSelectorExpr(field *Field) (*SelectorExpr, bool) {
	expr, ok := field.Type.(*SelectorExpr)
	return expr, ok
}

func isResponseWriterType(expr *SelectorExpr) bool {
	return expr.Sel.Name == "ResponseWriter"
}

func updateStmts(blockStmt *BlockStmt, index int, stmt Stmt) {
	fprintf := stmt.(*ExprStmt).X.(*CallExpr).Fun.(*SelectorExpr)
	fprintf.X.(*Ident).Name = "c"      // 修改接收者为c
	fprintf.Sel.Name = "SetBodyString" // 修改方法名为String

	callExpr := stmt.(*ExprStmt).X.(*CallExpr)
	callExpr.Args = callExpr.Args[1:] // 删除第一个参数

	newStmts := make([]Stmt, 0, len(blockStmt.List)*2)
	newStmts = append(newStmts, blockStmt.List[:index]...)
	newStmts = append(newStmts, &ExprStmt{
		X: &CallExpr{
			Fun: &SelectorExpr{
				X:   NewIdent("c"),
				Sel: NewIdent("SetStatusCode"),
			},
			Args: []Expr{&BasicLit{Kind: token.INT, Value: "200"}},
		},
	})
	newStmts = append(newStmts, &ExprStmt{X: callExpr})
	blockStmt.List = newStmts
}
