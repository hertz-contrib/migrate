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
	"fmt"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal/global"
	. "go/ast"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// CheckPtrStructName is a function used to check struct name
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
		if info.IsDir() {
			for _, dir := range global.IgnoreDirs {
				if strings.Contains(info.Name(), dir) {
					return nil
				}
			}
		}
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
		fmt.Println("[Error] search go.mod dir fail, error ", err)
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
