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
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func RunGoImports(path string) {
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}
	log.Println("Running goimports on", abs)
	cmd := exec.Command("go", "run", "-mod=mod", "golang.org/x/tools/cmd/goimports", "-w", abs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
}

func RunGoGet(path, repo string, version string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	var cmd *exec.Cmd
	if version != "" {
		cmd = exec.Command("go", "get", repo, version)
	} else {
		cmd = exec.Command("go", "get", repo)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
}
