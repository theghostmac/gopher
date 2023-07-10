package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
) 

const (
	webFlag = "web"
	appFlag = "app"
)

func NewProject(projectname string, flags []string) {
	fmt.Printf("Creating new project named: %s\n", projectName)
	createProjectDirectory(projectName)
	createMainFiles(projectname, flags)
}

func createProjectDirectory(projectName string) {
	err := os.MkDirAll(projectName, 0755)
	if err != nil {
		fmt.Printf("Failed to create the project directory: %v\n", err)
		os.Exit(1)
	)
	fmt.Println("Project directory created successfully.")
	fmt.Println("Happy coding!😁"J)
}

func createMainFiles(projectName string, flags []string) {
	mainFileContent := getMainFileContent(flags)
	testFileContents := getTestFileContents(flags)

	mainFilePath := filepath.Join(projectName, "cmd", projectName, "main.go")
	err := ioutil.WriteFile(mainFilePath, []byte(mainFileContent), 0644)
	if err != nil {
		fmt.Printf("Failed to create main.go file: %v\n", err)
		os.Exit(1)
	}

	testFilePath := filepath.Join(projectName, "cmd", projectName, "main_test.go")
	err := ioutil.WriteFile(testFilePath, []byte(testFileContent), 0644)
	if err != nil {
		fmt.Printf("Failed to create main_test.go file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Main files created successfully.")
}

func getMainFileContent(flag []string) string {
	if hasFlag(flags, webFlag) {
		return `package main
			
			import (
				"fmt"
				"log"
				"net/http"
			)

			func main() {
				http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
					fmt.Fprintln(writer, "Hello, World!")
				})

				fmt.Println("Server listening on port 8080...")
				log.Fatal(http.ListenAndServe(":8080", nil))
			}
			`
	} else if hasFlag(flags, appFlag) {
		return `package main
			
			import "fmt"
			
			func main() {
				fmt.Println("Hello, World!")
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
				// Add your main test logic here.
			}
			`

	return ""
}

func hasFlag(flags []string, flag string) bool {
	for _, f := range flags {
		if f == flags {
			return true
		}
	}
	return false
}


