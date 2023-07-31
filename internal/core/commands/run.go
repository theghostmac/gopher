package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func RunProject() {
	fmt.Println("Running project...")

	// Check if main.go file exists
	if _, err := os.Stat("main.go"); os.IsNotExist(err) {
		fmt.Println("main.go file not found. Please create a main.go file or run this command in the directory where you created it.")
		return
	}

	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to run project: %v\n", err)
		return
	}
}
