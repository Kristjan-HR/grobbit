package handmade

import (
	. "Grobbit/common"
	"strings"
	"unicode"
	//"fmt"
)

type Lexer struct {
	src          []rune
	n            int
	i            int
	ch           rune
	eoi          bool
	token        Token
	pos          Position
	errorHandler ErrorHandler
}

var charEscapeSet = map[rune]bool{
	't':  true,
	'n':  true,
	'\\': true,
	'\'': true,
}

var stringEscapeSet = map[rune]bool{
	't':  true,
	'n':  true,
	'\\': true,
	'"':  true,
}

///////////////////// Helper functions ////////////////////////////

func (lexer *Lexer) peekRune() (rune, bool) {
	if lexer.eoi || lexer.i+1 >= lexer.n {
		return rune(0), true
	}
	return lexer.src[lexer.i+1], false
}

func (lexer *Lexer) peekRuneIs(check rune) bool {
	r, err := lexer.peekRune()
	return !err && r == check
}

func (lexer *Lexer) nextRune() {
	if lexer.eoi {
		return
	}
	lexer.i++
	if lexer.i >= lexer.n {
		lexer.ch, lexer.eoi = rune(0), true
		return
	}
	ch := lexer.src[lexer.i]
	if ch == '\n' { // Note: '\r'
		lexer.pos.Line += 1
		lexer.pos.Col = -1
	}
	lexer.pos.Col += 1
	lexer.ch, lexer.eoi = ch, false
}

type pair struct {
	r rune
	t TokenType
}

func (lexer *Lexer) setToken(Tt TokenType, pairs ...pair) {
	if len(pairs) == 0 {
		lexer.token = Token{Type: Tt, Lexeme: string(lexer.ch), Pos: lexer.pos}
		lexer.nextRune()
		lexer.token.PosEnd = lexer.pos
	} else {
		startCh, startPos := lexer.ch, lexer.pos
		lexer.nextRune()
		if !lexer.eoi {
			for _, p := range pairs {
				if lexer.ch == p.r {
					lexer.token = Token{Type: p.t, Lexeme: string(startCh) + string(lexer.ch), Pos: startPos}
					lexer.nextRune()
					lexer.token.PosEnd = lexer.pos
					return
				}
			}
		}
		lexer.token = Token{Type: Tt, Lexeme: string(startCh), Pos: startPos, PosEnd: lexer.pos}
	}
}

func (lexer *Lexer) resetToken(Tt TokenType) {
	lexer.token = Token{Type: Tt, Lexeme: lexer.token.Lexeme + string(lexer.ch), Pos: lexer.token.Pos}
	lexer.nextRune()
	lexer.token.PosEnd = lexer.pos
}

func (lexer *Lexer) processWhiteSpaces() bool {
	n := 0
	for !lexer.eoi && unicode.IsSpace(lexer.ch) {
		n += 1
		lexer.nextRune()
	}
	return n > 0
}

///////////////////// Public functions ////////////////////////////

// Init function initialises the lexical analysis.
func (lexer *Lexer) Init(src []byte, handler ErrorHandler) {
	lexer.src, lexer.n, lexer.i = []rune(string(src)), len(src), -1
	lexer.eoi = false
	lexer.pos = Position{Line: 1, Col: 0}
	lexer.errorHandler = handler
	lexer.nextRune()
}

// NextToken reads and returns the next token.
func (lexer *Lexer) NextToken() Token {

	// TODO: Modify the code in here.

	for lexer.processWhiteSpaces() {
		// Nothing ...
	}

	if lexer.eoi {
		lexer.token = Token{
			Type:   TtEOI,
			Lexeme: "",
			Pos:    lexer.pos,
			PosEnd: lexer.pos,
		}
		return lexer.token
	}
	switch lexer.ch {
	case ',':
		lexer.setToken(TtComma)
		//lexer.nextRune()

	case ';':
		lexer.setToken(TtSemicolon)
		//lexer.nextRune()
	case '(':
		lexer.setToken(TtLParen)
		//lexer.nextRune()
	case ')':
		lexer.setToken(TtRParen)
		//lexer.nextRune()
	case '+':
		lexer.setToken(
			TtOpAdd,
			pair{'=', TtOpAddAssign},
			pair{'+', TtOpInc},
		)
	case '-':
		lexer.setToken(
			TtOpSub,
			pair{'=', TtOpSubAssign},
			pair{'-', TtOpDec},
		)
	case '%':
		lexer.setToken(
			TtOpMod,
			pair{'=', TtOpModAssign},
		)

	case '=':
		lexer.setToken(
			TtOpAssign,
			pair{'=', TtOpEq},
		)

	case '!':
		lexer.setToken(
			TtOpNot,
			pair{'=', TtOpNe},
		)

	case '<':
		lexer.setToken(
			TtOpLt,
			pair{'=', TtOpLe},
			pair{'-', TtOpArrow},
		)
	case '*':
		lexer.setToken(TtOpMul, pair{'=', TtOpMulAssign})
		//lexer.nextRune()
	case '.':
		next, isEnd := lexer.peekRune()

		if !isEnd && unicode.IsDigit(next) {
			start := lexer.pos
			lexeme, _ := lexer.processNumber()

			lexer.token = Token{
				Type:   TtFloat,
				Lexeme: lexeme,
				Pos:    start,
				PosEnd: lexer.pos,
			}
		} else {
			lexer.setToken(TtPeriod)
		}
	case '>':
		lexer.setToken(
			TtOpGt,
			pair{'=', TtOpGe},
		)

	case ':':
		lexer.setToken(
			TtColon,
			pair{'=', TtOpDefine},
		)

	case '&':
		lexer.setToken(
			TtOpBitAnd,
			pair{'=', TtOpBitAndAssign},
			pair{'&', TtOpAnd},
		)
	case '|':
		lexer.setToken(
			TtOpBitOr,
			pair{'=', TtOpBitOrAssign},
			pair{'|', TtOpOr},
		)
	case '{':
		lexer.setToken(TtLBrace)

	case '}':
		lexer.setToken(TtRBrace)
	case '/':
		if lexer.peekRuneIs('/') {
			lexer.singleLineComment()
			return lexer.NextToken()
		} else if lexer.peekRuneIs('*') {
			lexer.multilineComment()
			return lexer.NextToken()
		}

		lexer.setToken(
			TtOpDiv,
			pair{'=', TtOpDivAssign},
		)
	case '"':
		start := lexer.pos
		lexeme := lexer.processString()

		lexer.token = Token{
			Type:   TtString,
			Lexeme: lexeme,
			Pos:    start,
			PosEnd: lexer.pos,
		}
	default:
		if unicode.IsLetter(lexer.ch) || lexer.ch == '_' {
			start := lexer.pos
			lexeme := lexer.identifier()
			//fmt.Println("DEBUG:", lexeme, Lookup(lexeme))
			lexer.token = Token{
				Type:   Lookup(lexeme),
				Lexeme: lexeme,
				Pos:    start,
				PosEnd: lexer.pos,
			}
		} else if unicode.IsDigit(lexer.ch) {
			start := lexer.pos
			lexeme, isFloat := lexer.processNumber()

			tokenType := TtInt

			if isFloat {
				tokenType = TtFloat
			}

			lexer.token = Token{
				Type:   tokenType,
				Lexeme: lexeme,
				Pos:    start,
				PosEnd: lexer.pos,
			}

		} else {
			lexer.setToken(TtUnknown)
		}
	}
	return lexer.token
}

func (lexer *Lexer) identifier() string {
	var builder strings.Builder
	for !lexer.eoi && (unicode.IsLetter(lexer.ch) || unicode.IsDigit(lexer.ch) || lexer.ch == '_') {
		builder.WriteRune(lexer.ch)
		lexer.nextRune()
	}
	return builder.String()
}

func (lexer *Lexer) singleLineComment() {
	for !lexer.eoi && lexer.ch != '\n' {
		lexer.nextRune()
	}
}

func (lexer *Lexer) multilineComment() {
	start := lexer.pos

	lexer.nextRune()
	lexer.nextRune()

	for !lexer.eoi {
		if lexer.ch == '*' && lexer.peekRuneIs('/') {
			lexer.nextRune()
			lexer.nextRune()
			return
		}
		lexer.nextRune()
	}

	lexer.errorHandler(start, "multiline comment not terminated")
}

func (lexer *Lexer) processNumber() (string, bool) {
	var builder strings.Builder
	isFloat := false

	for !lexer.eoi && unicode.IsDigit(lexer.ch) {
		builder.WriteRune(lexer.ch)
		lexer.nextRune()
	}

	if !lexer.eoi && lexer.ch == '.' {
		isFloat = true
		builder.WriteRune(lexer.ch)
		lexer.nextRune()

		for !lexer.eoi && unicode.IsDigit(lexer.ch) {
			builder.WriteRune(lexer.ch)
			lexer.nextRune()
		}
	}

	if !lexer.eoi && (lexer.ch == 'e' || lexer.ch == 'E') {
		isFloat = true
		builder.WriteRune(lexer.ch)
		lexer.nextRune()

		if !lexer.eoi && (lexer.ch == '+' || lexer.ch == '-') {
			builder.WriteRune(lexer.ch)
			lexer.nextRune()
		}

		for !lexer.eoi && unicode.IsDigit(lexer.ch) {
			builder.WriteRune(lexer.ch)
			lexer.nextRune()
		}
	}

	return builder.String(), isFloat
}

func (lexer *Lexer) processString() string {
	var builder strings.Builder
	start := lexer.pos

	// Add the opening quotation mark
	builder.WriteRune(lexer.ch)
	lexer.nextRune()

	for !lexer.eoi {

		// Closing quotation mark: end of string
		if lexer.ch == '"' {
			builder.WriteRune(lexer.ch)
			lexer.nextRune()
			return builder.String()
		}

		// Escape sequence
		if lexer.ch == '\\' {
			builder.WriteRune(lexer.ch)
			lexer.nextRune()

			if lexer.eoi {
				lexer.errorHandler(start, "string not terminated")
				return builder.String()
			}

			if !stringEscapeSet[lexer.ch] {
				lexer.errorHandler(lexer.pos, "unknown escape sequence")
			}

			builder.WriteRune(lexer.ch)
			lexer.nextRune()
			continue
		}

		// A real newline cannot appear inside this string literal
		if lexer.ch == '\n' {
			lexer.errorHandler(start, "string not terminated")
			return builder.String()
		}

		builder.WriteRune(lexer.ch)
		lexer.nextRune()
	}

	lexer.errorHandler(start, "string not terminated")
	return builder.String()
}
