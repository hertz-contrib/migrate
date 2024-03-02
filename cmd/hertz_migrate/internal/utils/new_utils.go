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

package utils

import (
	. "go/ast"
)

func CheckPtrPkgAndStructName(s *StarExpr, packageName, structName string) bool {
	if se, ok := s.X.(*SelectorExpr); ok {
		return CheckSelPkgAndStruct(se, packageName, structName)
	}
	return false
}

func CheckCallPkgAndMethodName(c *CallExpr, packageName, methodName string) bool {
	if fs, ok := c.Fun.(*SelectorExpr); ok {
		return CheckSelPkgAndStruct(fs, packageName, methodName)
	}
	return false
}

func CheckSelPkgAndStruct(s *SelectorExpr, packageName, structName string) bool {
	if i, ok := s.X.(*Ident); ok {
		if packageName == "" {
			return s.Sel.Name == structName
		}
		if i.Name != packageName {
			return false
		}
		if s.Sel.Name == structName {
			return true
		}
	}
	return false
}

func CheckSelObj(s *SelectorExpr, packageName, structName string) bool {
	if ident, ok := s.X.(*Ident); ok {
		if ident.Obj != nil {
			if field, ok := ident.Obj.Decl.(*Field); ok {
				if ss, ok := field.Type.(*SelectorExpr); ok {
					return CheckSelPkgAndStruct(ss, packageName, structName)
				}
			}
		}
	}
	return false
}

func CheckFuncDeclReturnOne(fun *FuncDecl, pkg, name string) bool {
	if fun.Type.Results == nil {
		return false
	}
	if len(fun.Type.Results.List) != 1 {
		return false
	}
	starField := fun.Type.Results.List[0]
	if s, ok := starField.Type.(*StarExpr); ok {
		if sel, ok := s.X.(*SelectorExpr); ok {
			return CheckSelPkgAndStruct(sel, pkg, name)
		}
	}
	return true
}

func CheckObjStarExpr(obj *Object, pkg, name string) bool {
	if obj == nil || obj.Decl == nil {
		return false
	}

	if field, ok := obj.Decl.(*Field); ok {
		if starExpr, ok := field.Type.(*StarExpr); ok {
			if selExpr, ok := starExpr.X.(*SelectorExpr); ok {
				return CheckSelPkgAndStruct(selExpr, pkg, name)
			}
		}
	}

	if assignStmt, ok := obj.Decl.(*AssignStmt); ok {
		if call, ok := assignStmt.Rhs[0].(*CallExpr); ok {
			if i, ok := call.Fun.(*Ident); ok {
				if i.Obj == nil {
					return false
				}
				if i.Obj.Kind == Fun {
					return CheckFuncDeclReturnOne(i.Obj.Decl.(*FuncDecl), pkg, name)
				}
			}

			//if se, ok := call.Fun.(*SelectorExpr); ok {
			//	if i, ok := se.X.(*Ident); ok {
			//		return CheckObjStarExpr(i.Obj, pkg, name)
			//	}
			//}
		}
	}
	return false
}

func CheckObjSelExpr(obj *Object, pkg, name string) bool {
	if obj == nil || obj.Decl == nil {
		return false
	}

	if assignStmt, ok := obj.Decl.(*AssignStmt); ok {

		if call, ok := assignStmt.Rhs[0].(*CallExpr); ok {
			if se, ok := call.Fun.(*SelectorExpr); ok {
				return CheckSelPkgAndStruct(se, pkg, name)
			}
		}

		if ue, ok := assignStmt.Rhs[0].(*UnaryExpr); ok {
			if clit, ok := ue.X.(*CompositeLit); ok {
				if sel, ok := clit.Type.(*SelectorExpr); ok {
					return CheckSelPkgAndStruct(sel, pkg, name)
				}
			}
		}
	}
	return false
}
