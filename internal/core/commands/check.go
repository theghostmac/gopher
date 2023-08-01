package commands

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func CheckCodebase(target string) {
	fmt.Println("Checking the codebase...")

	err := filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isGoFile(path) {
			fmt.Println("Formatting file:", path)
			err := formatFile(path)
			if err != nil {
				fmt.Printf("Failed to format file: %v\n", err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error while checking the codebase: %v\n", err)
	}
}

func formatFile(filename string) error {
	code, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	formatted, err := format.Source(code)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, formatted, 0)
}

func isGoFile(path string) bool {
	return strings.HasSuffix(path, ".go")
}

//package commands
//
//import (
//	"fmt"
//	"go/ast"
//	"go/parser"
//	"go/token"
//	"path/filepath"
//	"strings"
//
//	"golang.org/x/tools/go/analysis"
//	"golang.org/x/tools/go/analysis/singlechecker"
//)
//
//var checkAnalyzer = &analysis.Analyzer{
//	Name: "gopherlint",
//	Doc:  "checks for any TODO comment in the code",
//	Run:  runCheck,
//}
//
//func CheckCodebase(target string) {
//	fmt.Println("Checking the codebase...")
//
//	fset := token.NewFileSet()
//	packages, err := parser.ParseDir(fset, target, nil, parser.ParseComments)
//	if err != nil {
//		fmt.Printf("Failed to parse codebase: %v\n", err)
//		return
//	}
//
//	for _, pkg := range packages {
//		for _, f := range pkg.Files {
//			checkFiles([]*ast.File{f})
//		}
//	}
//}
//
//func CheckDirectory(target string) {
//	fmt.Println("Checking directory...")
//
//	fset := token.NewFileSet()
//	packages, err := parser.ParseDir(fset, target, nil, parser.ParseComments)
//	if err != nil {
//		fmt.Printf("Failed to parse directory: %v\n", err)
//		return
//	}
//
//	files := []*ast.File{}
//	for _, pkg := range packages {
//		for _, f := range pkg.Files {
//			files = append(files, f)
//		}
//	}
//
//	checkFiles(files)
//}
//
//func CheckFile(filename string) {
//	if !isGoFile(filename) {
//		fmt.Println("Invalid file. Please provide a valid .go file.")
//		return
//	}
//
//	fmt.Println("Checking file:", filename)
//
//	fset := token.NewFileSet()
//	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
//	if err != nil {
//		fmt.Printf("Failed to parse file: %v\n", err)
//		return
//	}
//
//	checkFiles([]*ast.File{file})
//}
//
//func checkFiles(files []*ast.File) {
//	pass := &analysis.Pass{
//		Fset:    files[0].Set,
//		Files:   make(map[string]*ast.File),
//		Reportf: func(pos token.Pos, format string, args ...interface{}) { fmt.Printf(format+"\n", args...) },
//	}
//
//	// Collect ASTs for each file.
//	for _, f := range files {
//		pass.Files[f.Name.Name] = f
//	}
//
//	// Run the checkAnalyzer on the collected ASTs.
//	checkAnalyzer.Run(pass)
//}
//
//func runCheck(pass *analysis.Pass) (interface{}, error) {
//	for _, f := range pass.Files {
//		ast.Inspect(f, func(n ast.Node) bool {
//			if c, ok := n.(*ast.Comment); ok {
//				if strings.Contains(c.Text, "TODO") {
//					pass.Reportf(c.Pos(), "TODO comment found: %s", c.Text)
//				}
//			}
//			return true
//		})
//	}
//	return nil, nil
//}
//
//func RegisterCheckAnalyzer() {
//	singlechecker.Main(checkAnalyzer)
//}
//
//func isGoFile(path string) bool {
//	return filepath.Ext(path) == ".go"
//}
