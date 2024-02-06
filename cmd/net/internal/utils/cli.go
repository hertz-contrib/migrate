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

func RunGoGet(path string, repo string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	cmd := exec.Command("go", "get", "-u", repo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
}
