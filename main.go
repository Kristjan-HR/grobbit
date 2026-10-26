package main

import (
	"Grobbit/backend/icg"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// Also, create your own test.
	filename := "tests/icg/test_student.gr" // Set your default filename here

	// Check if at least one argument was passed after the program name
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	//handmade.DoTestLexer()
	//handmade.DoTestParser()
	//toolmade.DoTestLexer()
	//toolmade.DoTestParser()

	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	newFilename := filepath.Join("./out", strings.TrimSuffix(base, ext)+".ll")
	icg.DoTestLLVM(filename, newFilename)

	// To create an executable run: clang out/main.c out/test_01.ll
	// Creates an executable named: a.out
}
