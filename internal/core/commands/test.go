package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func RunTests(args []string) {
	if len(args) == 0 {
		RunAllTests()
	} else {
		RunSpecificTests(args)
	}
}

func RunAllTests() {
	fmt.Println("Running tests...")

	cmd := exec.Command("go", "test", "-v", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to run tests: %v\n", err)
		return
	}

	fmt.Println("Test result: ok.")
}

func RunSpecificTests(targests []string) {
	fmt.Println("Running specific tests...")

	args := []string{"test", "-v"}
	args = append(args, targests...)

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to run specific tests: %v\n", err)
		return
	}

	fmt.Println("Test results: ok.")
}
