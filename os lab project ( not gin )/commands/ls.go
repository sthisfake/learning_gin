package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func ListFiles(args []string) {
	// Run "ls" command with the provided arguments
	cmd := exec.Command("ls", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing 'ls':", err)
	}
}