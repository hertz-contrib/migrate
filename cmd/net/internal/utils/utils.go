package utils

import . "go/ast"

// CheckPtrStructName is a function used to check struc name
// like r.FormFile, can check r *http.Request struct name is 'Request'
func CheckPtrStructName(selExpr *SelectorExpr, name string) bool {
	if ident, ok := selExpr.X.(*Ident); ok {
		if field, ok := ident.Obj.Decl.(*Field); ok {
			starExpr, ok := field.Type.(*StarExpr)
			if !ok {
				return false
			}
			selExpr, ok := starExpr.X.(*SelectorExpr)
			if !ok {
				return false
			}
			if selExpr.Sel.Name == name {
				return true
			}
		}
	}
	return false
}

func CheckStructName(selExpr *SelectorExpr, name string) bool {
	if ident, ok := selExpr.X.(*Ident); ok {
		if field, ok := ident.Obj.Decl.(*Field); ok {
			selExpr, ok := field.Type.(*SelectorExpr)
			if !ok {
				return false
			}
			if selExpr.Sel.Name == name {
				return true
			}
		}
	}
	return false
}
