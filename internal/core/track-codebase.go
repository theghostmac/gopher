package core

import (
	"bufio"
	"fmt"
	"os"

	"io/fs"
	"path/filepath"
	"strings"
)

type TODO struct {
	FilePath   string
	LineNumber int
	Content    string
}

func TrackTODOs(projectDir string) ([]TODO, error) {
	var todos []TODO

	err := filepath.Walk(projectDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".go" {
			fileContents, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", path, err)
			}

			scanner := bufio.NewScanner(strings.NewReader(string(fileContents)))
			lineNum := 1
			for scanner.Scan() {
				line := scanner.Text()

				// More flexible TODO detection
				if strings.Contains(line, "// TODO") {
					todoIndex := strings.Index(line, "// TODO")
					todoContent := strings.TrimSpace(line[todoIndex+7:]) // Extract content after "// TODO"
					todo := TODO{
						FilePath:   path,
						LineNumber: lineNum,
						Content:    todoContent,
					}
					todos = append(todos, todo)
				}

				lineNum++
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return todos, nil
}
