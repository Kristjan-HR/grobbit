package toolmade

import (
	. "Grobbit/common"
	"Grobbit/frontend/toolmade/generated"

	"github.com/antlr4-go/antlr/v4"
)

var tokenTypes = map[int]TokenType{
	antlr.TokenEOF:                        TtEOI,
	generated.GrobbitLexerTtUnknown:       TtUnknown,
	generated.GrobbitLexerTtIdentifier:    TtIdentifier,
	generated.GrobbitLexerTtInt:           TtInt,
	generated.GrobbitLexerTtFloat:         TtFloat,
	generated.GrobbitLexerTtString:        TtString,
	generated.GrobbitLexerTtOpAdd:         TtOpAdd,
	generated.GrobbitLexerTtOpSub:         TtOpSub,
	generated.GrobbitLexerTtOpMul:         TtOpMul,
	generated.GrobbitLexerTtOpDiv:         TtOpDiv,
	generated.GrobbitLexerTtOpMod:         TtOpMod,
	generated.GrobbitLexerTtOpAnd:         TtOpAnd,
	generated.GrobbitLexerTtOpOr:          TtOpOr,
	generated.GrobbitLexerTtOpEq:          TtOpEq,
	generated.GrobbitLexerTtOpLt:          TtOpLt,
	generated.GrobbitLexerTtOpGt:          TtOpGt,
	generated.GrobbitLexerTtOpAssign:      TtOpAssign,
	generated.GrobbitLexerTtOpNot:         TtOpNot,
	generated.GrobbitLexerTtOpNe:          TtOpNe,
	generated.GrobbitLexerTtOpLe:          TtOpLe,
	generated.GrobbitLexerTtOpGe:          TtOpGe,
	generated.GrobbitLexerTtOpDefine:      TtOpDefine,
	generated.GrobbitLexerTtLParen:        TtLParen,
	generated.GrobbitLexerTtLBrace:        TtLBrace,
	generated.GrobbitLexerTtComma:         TtComma,
	generated.GrobbitLexerTtPeriod:        TtPeriod,
	generated.GrobbitLexerTtRParen:        TtRParen,
	generated.GrobbitLexerTtRBrace:        TtRBrace,
	generated.GrobbitLexerTtSemicolon:     TtSemicolon,
	generated.GrobbitLexerTtKwBreak:       TtKwBreak,
	generated.GrobbitLexerTtKwCase:        TtKwCase,
	generated.GrobbitLexerTtKwChan:        TtKwChan,
	generated.GrobbitLexerTtKwConst:       TtKwConst,
	generated.GrobbitLexerTtKwContinue:    TtKwContinue,
	generated.GrobbitLexerTtKwDefault:     TtKwDefault,
	generated.GrobbitLexerTtKwDefer:       TtKwDefer,
	generated.GrobbitLexerTtKwElse:        TtKwElse,
	generated.GrobbitLexerTtKwFallthrough: TtKwFallthrough,
	generated.GrobbitLexerTtKwFor:         TtKwFor,
	generated.GrobbitLexerTtKwFunc:        TtKwFunc,
	generated.GrobbitLexerTtKwGo:          TtKwGo,
	generated.GrobbitLexerTtKwGoto:        TtKwGo,
	generated.GrobbitLexerTtKwIf:          TtKwIf,
	generated.GrobbitLexerTtKwImport:      TtKwImport,
	generated.GrobbitLexerTtKwInterface:   TtKwInterface,
	generated.GrobbitLexerTtKwMap:         TtKwMap,
	generated.GrobbitLexerTtKwPackage:     TtKwPackage,
	generated.GrobbitLexerTtKwRange:       TtKwRange,
	generated.GrobbitLexerTtKwReturn:      TtKwReturn,
	generated.GrobbitLexerTtKwSelect:      TtKwSelect,
	generated.GrobbitLexerTtKwStruct:      TtKwStruct,
	generated.GrobbitLexerTtKwSwitch:      TtKwSwitch,
	generated.GrobbitLexerTtKwType:        TtKwType,
	generated.GrobbitLexerTtKwVar:         TtKwVar,
	generated.GrobbitLexerTtBool:          TtBool,
}

type Lexer struct {
	generated    *generated.GrobbitLexer
	errorHandler ErrorHandler
}

func mapToken(tk antlr.Token) Token {
	ttype := tokenTypes[tk.GetTokenType()]
	pos := Position{Line: tk.GetLine(), Col: tk.GetColumn() + 1}
	endPos := Position{Line: tk.GetLine(), Col: tk.GetColumn() + 1 + tk.GetStop() - tk.GetStart()}
	return Token{Type: ttype, Lexeme: tk.GetText(), Pos: pos, PosEnd: endPos}
}

///////////////////// Public functions ////////////////////////////

func (lexer *Lexer) Init(src []byte, handler ErrorHandler) {
	lexer.generated = generated.NewGrobbitLexer(antlr.NewInputStream(string(src)))
	lexer.errorHandler = handler
}

func (lexer *Lexer) NextToken() Token {
	tk := lexer.generated.NextToken()
	return mapToken(tk)
}
