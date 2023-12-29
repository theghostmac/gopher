package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type ProjectType string

const (
	WebApp  ProjectType = "web"
	CLIApp  ProjectType = "cli"
	General ProjectType = "general"
)

func CreateNewProject(name string, projectType ProjectType) error {
	// Create base project directory
	err := os.Mkdir(name, 0755)
	if err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create common project structure.
	commonDirs := []string{"cmd", "tests", "internal"}
	for _, dir := range commonDirs {
		err := createDirectory(filepath.Join(name, dir))
		if err != nil {
			return err
		}
	}

	// Create main.go in cmd directory (corrected path)
	err = createMainGoFile(filepath.Join(name, "cmd", "main.go")) // Used "main.go" directly
	if err != nil {
		return err
	}

	// TODO: Additional setup based on project type.
	switch projectType {
	case WebApp:
		err := createDirectory(filepath.Join(name, "internal", "api"))
		if err != nil {
			return err
		}
	case CLIApp:
		err := createDirectory(filepath.Join(name, "internal", "commands"))
		if err != nil {
			return err
		}
	default:
		// Additional setup for General apps.
	}

	// Initialize the repository.
	if err := initializeRepo(name); err != nil {
		return fmt.Errorf("failed to initialize repository: %w", err)
	}

	return nil
}

func createDirectory(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	return nil
}

func createMainGoFile(path string) error {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory for main.go: %w", err)
	}

	mainContent := `package main

import "fmt"

func main() {
    fmt.Println("Hello, Gopher!")
}`

	err := os.WriteFile(filepath.Join(path, "main.go"), []byte(mainContent), 0644) // Keep "main.go" here
	if err != nil {
		return fmt.Errorf("failed to create main.go file: %w", err)
	}

	return nil
}

func initializeRepo(projectName string) error {
	// Change to the project directory
	if err := os.Chdir(projectName); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	// Initialize Go module
	fmt.Println("Initializing Go module...")
	if err := exec.Command("go", "mod", "init").Run(); err != nil {
		return fmt.Errorf("failed to initialize Go module: %w", err)
	}
	fmt.Println("Go module initialized successfully!")

	// Initialize Git repository
	fmt.Println("Initializing Git repository...")
	if err := exec.Command("git", "init").Run(); err != nil {
		return fmt.Errorf("failed to initialize Git repository: %w", err)
	}
	fmt.Println("Git repository initialized successfully!")

	// Create README.md
	fmt.Println("Creating README.md...")
	if err := exec.Command("touch", "README.md").Run(); err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	fmt.Println("README.md created successfully!")

	return nil
}
