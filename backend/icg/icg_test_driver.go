package icg

import (
	. "Grobbit/backend/semana"
	. "Grobbit/common"
	"Grobbit/frontend/toolmade"
	"fmt"
	"os"
)

func TestLLVM(input []byte, filenameOut string) {
	var lines []string

	var errorHandler = func(pos Position, msg string) {
		var txt string
		txt = fmt.Sprintf("Error: (%d %d): %s\n", pos.Line, pos.Col, msg)
		lines = append(lines, txt)
	}

	var parser toolmade.Parser
	astree := parser.ParseAST(input)
	//var pv PrintVisitor
	//fmt.Println(pv.Root(astree))
	var sav SAVisitor
	scopes := sav.Root(astree, errorHandler)
	scopes = scopes
	var irv IRVisitor
	irv.Root(filenameOut, astree, errorHandler, scopes)
	for _, line := range lines {
		fmt.Print(line)
	}
}

func DoTestLLVM(filenameIn, filenameOut string) {
	src, err := os.ReadFile(filenameIn)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	TestLLVM(src, filenameOut)
}
