package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func ChangeDirectory(args []string , currentDir string)  string {
	if len(args) == 0 {
		fmt.Println("cd: missing operand") 
	}

	targetDir := args[0]
	if targetDir == ".." {
		// Move one directory up
		currentDir, _ = filepath.Split(currentDir[:len(currentDir)-1])

	} else {
		// Move to the specified directory
		newDir := filepath.Join(currentDir, targetDir)
		absPath, err := filepath.Abs(newDir)
		if err == nil {
			currentDir = absPath
			err = os.Chdir(absPath)
			if err != nil {
				fmt.Println("cd:", err)
			}
		} else {
			fmt.Println("cd:", err)
		}
	}

	return currentDir
}