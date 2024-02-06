// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	value, ok := m[httpProp]
	if !ok {
		return nil
	}
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
		switch vv := value.(type) {
		case *BinaryExpr:
			lit, ok := vv.X.(*BasicLit)
			if ok {
				return addBasicParamForOptionFunc(pack, funcName, lit.Value, lit.Kind)
			}
		case *SelectorExpr:
			return addBasicParamForOptionFunc(pack, funcName, "1", token.INT)
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
