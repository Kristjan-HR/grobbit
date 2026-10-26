package toolmade

import (
	"Grobbit/common"
	"Grobbit/frontend/toolmade/generated"

	"github.com/antlr4-go/antlr/v4"
)

type Parser struct {
}

func (parser *Parser) ParseSrc(src []byte, handler common.ErrorHandler) string {
	// Ignoring handler, needed for interface.
	lexer := generated.NewGrobbitLexer(antlr.NewInputStream(string(src)))
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	par := generated.NewGrobbitParser(tokenStream)
	parse_tree := par.Program()
	visitor := NewGrobbitParserVisitor()
	astree := visitor.Visit(parse_tree)
	var pv common.PrintVisitor
	//ast.Print(nil, astree) // For a more detailed printing.
	return pv.Root(astree)
}

func (parser *Parser) ParseAST(src []byte) common.Node {
	lexer := generated.NewGrobbitLexer(antlr.NewInputStream(string(src)))
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	par := generated.NewGrobbitParser(tokenStream)
	parse_tree := par.Program()
	visitor := NewGrobbitParserVisitor()
	astree := visitor.Visit(parse_tree)
	return astree
}
