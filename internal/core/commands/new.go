package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	webFlag = "web"
	appFlag = "app"
)

func NewProject(projectName string, flags []string) {
	fmt.Printf("Creating new project named: %s\n", projectName)
	createProjectDirectory(projectName)
	createMainFiles(projectName, flags)

	fmt.Println("Project directory created successfully.")
	fmt.Println("Happy coding!😁")
}

func createProjectDirectory(projectName string) {
	err := os.MkdirAll(projectName, 0755)
	if err != nil {
		fmt.Printf("Failed to create the project directory: %v\n", err)
		os.Exit(1)
	}

	err = os.Chdir(projectName)
	if err != nil {
		fmt.Printf("Failed to change to project directory: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("go", "mod", "init")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to iitialize Go modules: %v\n", err)
		os.Exit(1)
	}

	err = os.Chdir("..")
	if err != nil {
		fmt.Printf("Failed to change back to the previous directory: %v\n", err)
		os.Exit(1)
	}
}

func createMainFiles(projectName string, flags []string) {
	mainFileContent := getMainFileContent(flags)
	testFileContents := getTestFileContent(flags)
	cmdDir := filepath.Join(projectName, "cmd", projectName)
	err := os.MkdirAll(cmdDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create cmd directory: %v\n", err)
		os.Exit(1)
	}

	projectDir := filepath.Join(cmdDir)
	err = os.MkdirAll(projectDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create project directory: %v\n", err)
		os.Exit(1)
	}

	mainFilePath := filepath.Join(projectDir, "main.go")
	err = os.WriteFile(mainFilePath, []byte(mainFileContent), 0644)
	if err != nil {
		fmt.Printf("Failed to create main.go file: %v\n", err)
		os.Exit(1)
	}

	testFilePath := filepath.Join(projectDir, "main_test.go")
	err = os.WriteFile(testFilePath, []byte(testFileContents), 0644)
	if err != nil {
		fmt.Printf("Failed to create main_test.go file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Main files created successfully.")
}

func getMainFileContent(flags []string) string {
	if hasFlag(flags, appFlag) {
		return `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
	}

	if hasFlag(flags, webFlag) {
		return `package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	}

	return ""
}

func getTestFileContent(flags []string) string {
	if hasFlag(flags, appFlag) {
		return `package main_test

import "testing"

func TestMain(t *testing.T) {
	t.Log("Running main test...")
	// Add your main test logic here
}
`
	}

	return ""
}

func hasFlag(flags []string, flag string) bool {
	for _, f := range flags {
		if f == flag {
			return true
		}
	}
	return false
}
