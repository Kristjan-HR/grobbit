package semana

import (
	"fmt"
	"go/ast"
	"go/importer"
	gp "go/parser"
	"go/scanner"
	"go/token"
	"go/types"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
const sourceCode = `
package main

import "fmt"

func foo(a, b int) int {
   	c := a + b
	return c
}
func main() {
	x := 10
	y := "hello"
	// This line would cause a type error if uncommented:
	// z := x + y
	fmt.Println(x, y)
	if j := 1; true {
        const k = 10
		fmt.Println(j, k)
	}

}
`

func Test() {
	fset := token.NewFileSet() // Create a FileSet to manage source positions

	// Parse the source code into an AST file
	f, err := parser.ParseFile(fset, "example.go", sourceCode, 0)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	// Configure the type checker
	conf := types.Config{
		Importer: importer.Default(), // Use the default importer for standard library packages
		Error: func(err error) {
			// Handle type errors, e.g., print them
			fmt.Printf("Type error: %v\n", err)
		},
	}

	// Create a types.Info struct to collect type information
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}

	// Type-check the package
	pkg, err := conf.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		// Note: conf.Error handler will already print type errors,
		// but conf.Check itself returns an error if there are any.
		log.Printf("Package check failed: %v", err)
	} else {
		fmt.Printf("Package %q type-checked successfully.\n", pkg.Path())
	}

	// You can now access type information through the 'info' struct
	// For example, to find the type of 'x':
	for id, obj := range info.Defs {
		if obj != nil {
			fmt.Printf("Variable '%s' has type: %s\n", id.Name, obj.Type().String())
		}
	}

	for _, s := range info.Scopes {
		str := s.String()
		if strings.HasPrefix(str, "function ") || strings.HasPrefix(str, "block ") ||
			strings.HasPrefix(str, "if ") || strings.HasPrefix(str, "for ") {
			continue
		}
		var buf strings.Builder
		s.WriteTo(&buf, 0, true)
		re := regexp.MustCompile(` 0x[0-9a-fA-F]+`)
		outputString := re.ReplaceAllString(buf.String(), "")
		fmt.Println(outputString)
	}

}
*/

func ParseGoFile(fs *token.FileSet, filename string) *ast.File {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fs.AddFile(filename, fs.Base(), len(b))
	astNode, err := gp.ParseFile(fs, filename, nil, 0)
	if err != nil {
		if list, ok := err.(scanner.ErrorList); ok {
			for _, e := range list {
				fmt.Printf("Parsing error: %s\n", e.Error())
			}
		} else {
			log.Fatalf("Unexpected error type: %v", err)
		}
	}
	return astNode
}

func Test(filename string) {
	fmt.Println("TestGoSemana")
	var fs = token.NewFileSet()
	var asTree []*ast.File
	asTree = append(asTree, ParseGoFile(fs, filename))
	//asTree = append(asTree, ParseGoFile(fs, filename+"2"))

	// Configure the type checker
	conf := types.Config{
		Importer: importer.Default(),
		Error: func(err error) {
			// Handle type errors, e.g., print them
			fmt.Printf("Type error: %v\n", err)
		},
	}

	// Create a types.Info struct to collect type information
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}

	// Type-check the package
	pkg, err := conf.Check("main", fs, asTree, info)
	if err != nil {
		// Note: conf.Error handler will already print type errors,
		// but conf.Check itself returns an error if there are any.
		log.Printf("Package check failed: %v", err)
	} else {
		fmt.Printf("Package %q type-checked successfully.\n", pkg.Path())
	}
	/*
		// You can now access type information through the 'info' struct
		// For example, to find the type of 'x':
		for id, obj := range info.Defs {
			if obj != nil {
				fmt.Printf("Variable '%s' has type: %s\n", id.Name, obj.Type().String())
			}
		}
	*/
	var buf strings.Builder
	pkg.Scope().WriteTo(&buf, 0, true)
	re := regexp.MustCompile(` 0x[0-9a-fA-F]+`)
	outputString := re.ReplaceAllString(buf.String(), "")
	fmt.Println(outputString)
}
