package logic

import (
	. "go/ast"
	"go/token"
)

func addParamForOptionFunc(pack, funcName, value string, valueType token.Token) *CallExpr {
	return &CallExpr{
		Fun: &SelectorExpr{
			X:   NewIdent(pack),
			Sel: NewIdent(funcName),
		},
		Args: []Expr{
			&BasicLit{
				Kind:  token.INT,
				Value: value,
			},
		},
	}
}
