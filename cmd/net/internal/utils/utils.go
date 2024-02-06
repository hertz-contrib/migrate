package utils

import (
	"fmt"
	. "go/ast"
	"os"
	"path/filepath"
	"regexp"
)

// CheckPtrStructName is a function used to check struc name
// like r.FormFile, can check r *http.Request struct name is 'Request'
func CheckPtrStructName(selExpr *SelectorExpr, name string) bool {
	if ident, ok := selExpr.X.(*Ident); ok {
		if ident.Obj == nil {
			return false
		}
		if field, ok := ident.Obj.Decl.(*Field); ok {
			starExpr, ok := field.Type.(*StarExpr)
			if !ok {
				return false
			}
			_selExpr, ok := starExpr.X.(*SelectorExpr)
			if !ok {
				return false
			}
			if _selExpr.Sel.Name == name {
				return true
			}
		}
	}
	return false
}

func CheckStarProp(ident *Ident, name string) bool {
	if ident.Obj == nil || ident.Obj.Decl == nil {
		return false
	}
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
	return false
}

func CheckProp(ident *Ident, name string) bool {
	if ident.Obj == nil || ident.Obj.Decl == nil {
		return false
	}
	if field, ok := ident.Obj.Decl.(*Field); ok {
		selExpr, ok := field.Type.(*SelectorExpr)
		if !ok {
			return false
		}
		if selExpr.Sel.Name == name {
			return true
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

func ReplaceParamsInStr(s string) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	resultString := re.ReplaceAllString(s, ":$1")
	return resultString
}

func CollectGoFiles(directory string) ([]string, error) {
	var goFiles []string
	abs, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}
	err = filepath.Walk(abs, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}

		return nil
	})

	return goFiles, err
}

func SearchAllDirHasGoMod(path string) (dirs []string) {
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = filepath.Walk(abs, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			modFilePath := filepath.Join(path, "go.mod")
			if _, err := os.Stat(modFilePath); err == nil {
				dirs = append(dirs, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	return dirs
}
