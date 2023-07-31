package commands

import (
	"fmt"
	"os/exec"
)

func AddDependencies(dependencies []string) {
	if len(dependencies) == 0 {
		fmt.Println("Usage: gopher add <dependency1> <dependency2> ...")
		return
	}

	for _, deps := range dependencies {
		fmt.Printf("Fetching and installing dependencies: %s\n", deps)
		err := exec.Command("go", "get", deps).Run()
		if err != nil {
			fmt.Printf("Failed to fetch and install dependency %s: %v\n", deps, err)
		} else {
			fmt.Printf("Dependency %s fetched and installed successfully.\n", deps)
		}
	}
}
