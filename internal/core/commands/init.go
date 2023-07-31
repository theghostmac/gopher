package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// InitializeProject receives a projectName, builds a directory, and initializes go.mod in it.
func InitializeProject() {
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("A go.mod file already exists. The project is already initialized.")
		return
	}

	fmt.Println("Initializing new Go project with Go modules...")

	cmd := exec.Command("go", "mod", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to initialize Go modules: %v\n", err)
		return
	}

	fmt.Println("Project initialized successfully with Go modules.")
}
