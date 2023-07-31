package commands

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	binaryName = "gopher"
)

func BuildProject() {
	fmt.Println("Building projects...")

	// Check if a go.mod file exists
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println("go.mod file is not found.")
	}

	// Get the main package directory
	mainPackage, err := getMainPackage()
	if err != nil {
		fmt.Printf("Failed to get the main package directory: %v\n", err)
		return
	}

	// Build the project
	cmd := exec.Command("go", "build", "-o", binaryName, mainPackage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to build project: %v\n", err)
		return
	}

	fmt.Println("Project built successfully")
}

func getMainPackage() (string, error) {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
