package initialprint

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ANSI escape codes for text coloring
const (
	ansiColorCyan   = "\033[96m"
	ansiColorGreen  = "\033[92m"
	ansiColorMagenta = "\033[95m"
	ansiColorYellow = "\033[33m"
	ansiColorReset  = "\033[0m"
)



func PrintPrompt(currentDir string) {
	username := getUsername()
	hostname := getHostname()
	dir := getCurrentRelativeDirectory(currentDir)
	osName := "POS"
	fmt.Printf("%s@%s %s %s:%s$ ", ansiColorCyan+username, ansiColorGreen+hostname, ansiColorMagenta+osName, ansiColorYellow+dir, ansiColorReset)
}

func getUsername() string {
	user, err := user.Current()
	if err != nil {
		return "user"
	}
	userName := extractUser(user.Username)
	return userName
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown-host"
	}
	return hostname
}

func getCurrentRelativeDirectory(currentDir string) string {
	dir, _ := filepath.Rel(currentDir, getCurrentDirectory())
	if dir == "." {
		return "~"
	}
	return dir
}

func getCurrentDirectory() string {
	dir, _ := os.Getwd()
	return dir
}

func extractUser(input string ) string {

	// Find the position of the backslash ("\") character
	index := strings.LastIndex(input, "\\")
	if index != -1 {
		// Extract the substring after the backslash
		result := input[index+1:]
		return result
	} else {
		return "user"
	}

}