package netHttp

import (
	. "go/ast"
	"go/token"
)

func addBasicParamForOptionFunc(pack, funcName, value string, valueType token.Token) *CallExpr {
	return &CallExpr{
		Fun: &SelectorExpr{
			X:   NewIdent(pack),
			Sel: NewIdent(funcName),
		},
		Args: []Expr{
			&BasicLit{
				Kind:  valueType,
				Value: value,
			},
		},
	}
}

func addParamInOption(pack, funcName, httpProp string, m map[string]any) *CallExpr {
	value := m[httpProp]
	switch httpProp {
	case "Addr":
		switch v := value.(type) {
		case *BasicLit:
			return addBasicParamForOptionFunc(pack, funcName, v.Value, v.Kind)
		case *CallExpr:
			if s, ok := packCallExpr2Str(v); ok {
				return addBasicParamForOptionFunc(pack, funcName, s, token.STRING)
			}
		}
	case "IdleTimeout", "WriteTimeout", "ReadTimeout":
		if value == nil {
			return nil
		}
		v, ok := value.(*BinaryExpr).X.(*BasicLit)
		if ok {
			return addBasicParamForOptionFunc(pack, funcName, v.Value, v.Kind)
		}
	}

	return nil
}

func packCallExpr2Str(ce *CallExpr) (string, bool) {
	argstring := ""
	selExpr, ok := ce.Fun.(*SelectorExpr)
	if !ok {
		return "", false
	}
	ident, ok := selExpr.X.(*Ident)
	if !ok {
		return "", false
	}
	packageName := ident.Name
	fName := selExpr.Sel.Name
	args := ce.Args
	for _, arg := range args {
		argstring += arg.(*BasicLit).Value + ","
	}
	return packageName + "." + fName + "(" + argstring + ")", true
}
