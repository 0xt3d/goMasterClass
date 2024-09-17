package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	tools := []string{"curl", "git", "htop"}

	for _, tool := range tools {
		if !isToolInstalled(tool) {
			fmt.Printf("Installing %s...\n", tool)
			if err := installTool(tool); err != nil {
				fmt.Printf("Failed to install %s: %v\n", tool, err)
			} else {
				fmt.Printf("%s installed successfully.\n", tool)
			}
		} else {
			fmt.Printf("%s is already installed.\n", tool)
		}
	}
}

func isToolInstalled(tool string) bool {
	cmd := exec.Command("which", tool)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func installTool(tool string) error {
	cmd := exec.Command("sudo", "apt-get", "install", "-y", tool)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}