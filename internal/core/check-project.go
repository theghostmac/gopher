package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CheckProject(cwd string) error {
	// Check for cmd/main.go
	mainFile := filepath.Join(cwd, "cmd", "main.go")
	if _, err := os.Stat(mainFile); err == nil {
		cwd = filepath.Dir(mainFile) // Use cmd directory if main.go exists
	} else {
		// Allow user to specify the directory if main.go not found
		fmt.Println("cmd/main.go not found in the current directory.")
		fmt.Print("Enter the directory containing your main.go file: ")
		var input string
		fmt.Scanln(&input)
		cwd = input
	}

	// Confirm repository existence
	fmt.Println("Checking if the repository exists...")
	if _, err := os.Stat(cwd); err != nil {
		return fmt.Errorf("repository not found at: %s", cwd)
	}
	fmt.Println("Repository found!")

	fmt.Println("Checking for compilation errors in", cwd, "...")

	cmd := exec.Command("go", "build", "-o", "/dev/null")
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout // Print compilation output for debugging
	cmd.Stderr = os.Stderr // Print error output for debugging

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("compilation errors found")
	}

	fmt.Println("Compilation successful")
	return nil
}
