package netHttp

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/migrate/cmd/net/internal/utils"
	. "go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func CollectHandlerFuncName(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	collectTmpFuncName(cur, funcSet)
	collectCommonFuncName(cur, funcSet)
	collectExprStmtName(cur, funcSet)
}

func collectExprStmtName(cur *astutil.Cursor, mapSet mapset.Set[string]) {
	var funcName string
	exprStmt, ok := cur.Node().(*ExprStmt)
	if !ok {
		return
	}
	callExpr, ok := exprStmt.X.(*CallExpr)
	if !ok {
		return
	}
	ident, ok := callExpr.Fun.(*Ident)
	if !ok {
		return
	}
	funcName = ident.Name
	for _, i := range callExpr.Args {
		switch t := i.(type) {
		case *Ident:
			if utils.CheckStarProp(t, "Request") {
				mapSet.Add(funcName)
			}
			if utils.CheckProp(t, "ResponseWriter") {
				mapSet.Add(funcName)
			}
		}
	}
}

func collectTmpFuncName(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	var (
		funcName  string
		paramList []*Field
	)
	blockStmt, ok := cur.Node().(*BlockStmt)
	if !ok {
		return
	}

	for _, stmt := range blockStmt.List {
		as, ok := stmt.(*AssignStmt)
		if !ok {
			return
		}
		if len(as.Lhs) == 1 {
			ident, ok := as.Lhs[0].(*Ident)
			if !ok {
				return
			}
			funcName = ident.Name
		}
		if len(as.Rhs) == 1 {
			funcLit, ok := as.Rhs[0].(*FuncLit)
			if !ok {
				return
			}
			paramList = funcLit.Type.Params.List
			for _, field := range paramList {
				switch t := field.Type.(type) {
				case *SelectorExpr:
					if t.Sel.Name == "ResponseWriter" {
						funcSet.Add(funcName)
					}
				case *StarExpr:
					selExpr, ok := t.X.(*SelectorExpr)
					if ok {
						if selExpr.Sel.Name == "Request" {
							funcSet.Add(funcName)
						}
					}
				}
			}
		}
	}
}

func collectCommonFuncName(cur *astutil.Cursor, funcSet mapset.Set[string]) {
	var (
		paramList []*Field
	)
	funcDecl, ok := cur.Node().(*FuncDecl)
	if !ok {
		return
	}
	funcType := funcDecl.Type
	paramList = funcType.Params.List

	for _, field := range paramList {
		switch t := field.Type.(type) {
		case *SelectorExpr:
			if t.Sel.Name == "ResponseWriter" {
				funcSet.Add(funcDecl.Name.Name)
			}
		case *StarExpr:
			selExpr, ok := t.X.(*SelectorExpr)
			if ok {
				if selExpr.Sel.Name == "Request" {
					funcSet.Add(funcDecl.Name.Name)
				}
			}
		}
	}
}
