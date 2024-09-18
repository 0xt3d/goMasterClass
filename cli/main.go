package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	tools := []string{"curl", "git", "htop", "docker"}

	for _, tool := range tools {
		if tool == "docker" {
			if !isToolInstalled(tool) {
				fmt.Printf("Installing %s...\n", tool)
				if err := installDocker(); err != nil {
					fmt.Printf("Failed to install %s: %v\n", tool, err)
				} else {
					fmt.Printf("%s installed successfully.\n", tool)
				}
			} else {
				fmt.Printf("%s is already installed.\n", tool)
			}
		} else {
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

func installDocker() error {
	// Update the apt package index and install packages to allow apt to use a repository over HTTPS
	cmd := exec.Command("sudo", "apt-get", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("sudo", "apt-get", "install", "-y", "apt-transport-https", "ca-certificates", "curl", "software-properties-common")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Add Docker’s official GPG key
	cmd = exec.Command("curl", "-fsSL", "https://download.docker.com/linux/ubuntu/gpg", "|", "sudo", "apt-key", "add", "-")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Add Docker apt repository
	cmd = exec.Command("sudo", "add-apt-repository", "deb", "[arch=amd64]", "https://download.docker.com/linux/ubuntu", "focal", "stable")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Update the apt package index again
	cmd = exec.Command("sudo", "apt-get", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Install Docker CE
	cmd = exec.Command("sudo", "apt-get", "install", "-y", "docker-ce")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
