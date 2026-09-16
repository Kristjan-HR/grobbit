package handmade

import (
	. "Grobbit/common"
	"Grobbit/frontend/stub"
	"fmt"
	"slices"
)

type Parser struct {
	lexer LexerInterface
	token Token
}

///////////////////////////////////// Helper methods ///////////////////////////////////////

func (parser *Parser) oneOf(ttcheck TokenType, tts ...TokenType) bool {
	return slices.Contains(tts, ttcheck)
}

func (parser *Parser) matchError(msg string) {
	fmt.Println("Parsing error:", msg)
	panic("Parser.match: invalid token")
}

func (parser *Parser) match(tt TokenType) {
	if parser.token.Type != tt {
		msg := fmt.Sprintf("found '%s' but expected '%s' at position %s.", parser.token.Type, tt, parser.token.Pos)
		parser.matchError(msg)
	}
	parser.token = parser.lexer.NextToken()
}

func (parser *Parser) matchIf(tt TokenType) bool {
	if parser.token.Type == tt {
		parser.match(tt)
		return true
	}
	return false
}

///////////////////////////////////// Exported methods ///////////////////////////////////////

func (parser *Parser) ParseSrc(src []byte, handler ErrorHandler) string {
	var lex stub.Lexer // NOTE: You can change to your own lexer, if you like.
	lex.Init(src, handler)
	parser.Init(&lex)
	astree := parser.Parse()
	var pv PrintVisitor
	//ast.Print(nil, astree)  // For a more detailed printing.
	return pv.Root(astree)
}

func (parser *Parser) Init(lex LexerInterface) {
	parser.lexer = lex
	parser.token = parser.lexer.NextToken()
}

func (parser *Parser) Parse() Node {
	var (
		imports []*ImportSpecNode
		decls   []DeclNode
	)
	token := parser.token
	parser.match(TtKwPackage)
	ident := parser.Identifier()
	parser.match(TtSemicolon)
	for parser.token.Type == TtKwImport {
		specs := parser.ImportDecl()
		imports = append(imports, specs...)
		parser.match(TtSemicolon)
	}
	for parser.oneOf(parser.token.Type, TtKwConst, TtKwVar, TtKwFunc) {
		var decl DeclNode
		switch parser.token.Type {
		case TtKwConst:
			decl = parser.ConstDecl()
		case TtKwVar:
			decl = parser.VarDecl()
		case TtKwFunc:
			decl = parser.FuncDecl()
		default:
			// do nothing (should not happen)
		}
		decls = append(decls, decl)
		parser.match(TtSemicolon)
	}
	parser.match(TtEOI)
	return &FileNode{Tok: token, Name: ident, Imports: imports, Decls: decls}
}

func (parser *Parser) ImportDecl() []*ImportSpecNode {
	var (
		spec  *ImportSpecNode
		specs []*ImportSpecNode
	)
	parser.match(TtKwImport)
	if parser.matchIf(TtLParen) {
		for parser.oneOf(parser.token.Type, TtIdentifier, TtString) {
			spec = parser.ImportSpec()
			specs = append(specs, spec)
			parser.match(TtSemicolon)
		}
		parser.match(TtRParen)
	} else {
		spec = parser.ImportSpec()
		specs = append(specs, spec)
	}
	return specs
}

func (parser *Parser) ImportSpec() *ImportSpecNode {
	var identifier *IdentifierNode = nil
	token := parser.token
	if parser.token.Type == TtIdentifier {
		identifier = parser.Identifier()
	}
	path := parser.token.Lexeme
	parser.match(TtString)
	// Token is TtString, or TtIdentifier depending on whether path is prefixed or not.
	return &ImportSpecNode{Tok: token, Name: identifier, Path: path}
}

func (parser *Parser) VarDecl() DeclNode {
	var (
		spec  SpecNode
		specs []SpecNode
	)
	token := parser.token
	parser.match(TtKwVar)
	if parser.matchIf(TtLParen) {
		for parser.token.Type == TtIdentifier {
			spec = parser.VarSpec()
			specs = append(specs, spec)
			parser.match(TtSemicolon)
		}
		parser.match(TtRParen)
	} else {
		spec = parser.VarSpec()
		specs = append(specs, spec)
	}
	return &GenDeclNode{Tok: token, Specs: specs}
}

func (parser *Parser) VarSpec() *ValueSpecNode {
	var (
		exprs []ExprNode
	)
	token := parser.token
	idents := parser.Identifiers()
	typeExpr := parser.TypeName()
	if parser.matchIf(TtOpAssign) {
		expr := parser.BasicLiteral()
		exprs = append(exprs, expr)
		for parser.matchIf(TtComma) {
			expr = parser.BasicLiteral()
			exprs = append(exprs, expr)
		}
	}
	return &ValueSpecNode{Token: token, Names: idents, Type: typeExpr, Values: exprs}
}

func (parser *Parser) FuncDecl() DeclNode {
	var (
		params []*Field
		result ExprNode
	)
	token := parser.token
	parser.match(TtKwFunc)
	ident := parser.Identifier()
	parser.match(TtLParen)
	if parser.token.Type != TtRParen {
		params = parser.ParamDecl()
	}
	parser.match(TtRParen)
	if parser.token.Type != TtLBrace {
		result = parser.TypeName()
	}
	body := parser.BlockStatement()
	return &FuncDeclNode{Tok: token, Name: ident, Type: &FuncType{Params: params, Type: result}, Body: body}
}

func (parser *Parser) ParamDecl() []*Field {
	var fields []*Field
	field := parser.ParamSpec()
	fields = append(fields, field)
	for parser.matchIf(TtComma) {
		field = parser.ParamSpec()
		fields = append(fields, field)
	}
	return fields
}

func (parser *Parser) ParamSpec() *Field {
	idents := parser.Identifiers()
	expr := parser.TypeName()
	return &Field{Names: idents, Type: expr}
}

func (parser *Parser) TypeName() ExprNode {
	return parser.Identifier()
}

func (parser *Parser) Identifiers() []*IdentifierNode {
	var idents []*IdentifierNode
	ident := parser.Identifier()
	idents = append(idents, ident)
	for parser.matchIf(TtComma) {
		ident := parser.Identifier()
		idents = append(idents, ident)
	}
	return idents
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func (parser *Parser) Statement() StmtNode {
	var stmt StmtNode = nil
	switch parser.token.Type {
	case TtKwConst, TtKwVar:
		stmt = parser.DeclStatement()
	case TtKwReturn:
		stmt = parser.ReturnStatement()
	case TtKwBreak:
		stmt = parser.BreakStatement()
	case TtKwContinue:
		stmt = parser.ContinueStatement()
	case TtKwIf:
		stmt = parser.IfStatement()
	case TtKwFor:
		stmt = parser.ForStatement()
	case TtLBrace:
		stmt = parser.BlockStatement()
	default:
		stmt = parser.SimpleStatement()
	}
	return stmt
}

func (parser *Parser) DeclStatement() *DeclStmtNode {
	var decl DeclNode
	token := parser.token
	switch parser.token.Type {
	case TtKwConst:
		decl = parser.ConstDecl()
	case TtKwVar:
		decl = parser.VarDecl()
	}
	return &DeclStmtNode{Tok: token, Decl: decl}
}

func (parser *Parser) BlockStatement() *BlockStmtNode {
	var stmts []StmtNode
	token := parser.token
	parser.match(TtLBrace)
	for parser.token.Type != TtRBrace {
		stmt := parser.Statement()
		stmts = append(stmts, stmt)
		parser.matchIf(TtSemicolon)
	}
	parser.match(TtRBrace)
	return &BlockStmtNode{Tok: token, List: stmts}
}

func (parser *Parser) ReturnStatement() *ReturnStmtNode {
	var expr ExprNode
	token := parser.token
	parser.match(TtKwReturn)
	if !parser.oneOf(parser.token.Type, TtSemicolon, TtRBrace) {
		expr = parser.Expression()
	}
	return &ReturnStmtNode{Token: token, Result: expr}
}

func (parser *Parser) ContinueStatement() StmtNode {
	token := parser.token
	parser.match(TtKwContinue)
	return &BranchStmtNode{Tok: token}
}

func (parser *Parser) ForStatement() *ForStmtNode {
	var init, cond, post StmtNode

	token := parser.token
	parser.match(TtKwFor)

	if parser.token.Type != TtLBrace {

		// for ; condition ; post { ... }
		// or for ;; { ... }
		if parser.matchIf(TtSemicolon) {

			if parser.token.Type != TtSemicolon {
				cond = parser.SimpleStatement()
			}

			parser.match(TtSemicolon)

			if parser.token.Type != TtLBrace {
				post = parser.SimpleStatement()
			}

		} else {

			// Could be either:
			// for condition { ... }
			// or:
			// for init ; condition ; post { ... }

			first := parser.SimpleStatement()

			if parser.matchIf(TtSemicolon) {
				init = first

				if parser.token.Type != TtSemicolon {
					cond = parser.SimpleStatement()
				}

				parser.match(TtSemicolon)

				if parser.token.Type != TtLBrace {
					post = parser.SimpleStatement()
				}

			} else {
				cond = first
			}
		}
	}

	var condExpr ExprNode

	if cond != nil {
		exprStmt, ok := cond.(*ExprStmtNode)

		if !ok {
			msg := fmt.Sprintf(
				"Condition needs to be an expression at line %d.",
				parser.token.Pos.Line,
			)
			parser.matchError(msg)
		}

		condExpr = exprStmt.Expr
	}

	block := parser.BlockStatement()

	return &ForStmtNode{
		Tok:  token,
		Init: init,
		Cond: condExpr,
		Post: post,
		Body: block,
	}
}
func (parser *Parser) SimpleStatement() StmtNode {
	var stmt StmtNode = nil
	var exprsLhs, exprsRhs []ExprNode
	exprsLhs = parser.Expressions()
	token := parser.token
	if parser.oneOf(parser.token.Type, TtOpDefine, TtOpAssign) {
		parser.match(parser.token.Type)
		exprsRhs = parser.Expressions()
		// We leave to it to semantic analysis to check that only identifier expressions on the left-hand-side
		stmt = &AssignStmtNode{Tok: token, Lhs: exprsLhs, Rhs: exprsRhs}
	} else {
		if len(exprsLhs) != 1 {
			msg := fmt.Sprintf("Expected only a single expression at position %s.", parser.token.Pos)
			parser.matchError(msg)
			// terminates program
		} else { // Expression statement
			// We leave to it to semantic analysis to check the correct expression type (call expression)
			stmt = &ExprStmtNode{Tok: token, Expr: exprsLhs[0]}
		}
	}
	return stmt
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func (parser *Parser) Expressions() []ExprNode {
	var exprs []ExprNode
	expr := parser.Expression()
	exprs = append(exprs, expr)
	for parser.matchIf(TtComma) {
		expr = parser.Expression()
		exprs = append(exprs, expr)
	}
	return exprs
}

func (parser *Parser) Expression() ExprNode {
	var expr ExprNode = nil
	expr = parser.ExprAnd()
	token := parser.token
	for parser.matchIf(TtOpOr) { // TtOpOr has the lowest precedence (see Go specs)
		rhs := parser.ExprAnd()
		expr = &BinaryExprNode{Tok: token, Lhs: expr, Rhs: rhs}
		token = parser.token
	}
	return expr
}

func (parser *Parser) Identifier() *IdentifierNode {
	node := &IdentifierNode{Tok: parser.token}
	parser.match(TtIdentifier)
	return node
}

func (parser *Parser) BasicLiteral() ExprNode {
	if !parser.isBasicLiteral() {
		msg := fmt.Sprintf(
			"Expected a literal but found '%s' at position %s.",
			parser.token.Type,
			parser.token.Pos,
		)
		parser.matchError(msg)
	}

	token := parser.token
	parser.match(parser.token.Type)

	return &LiteralNode{Tok: token}
}

////////////////// You implement the methods below (add methods as needed) ////////////////////

func (parser *Parser) ConstDecl() DeclNode {
	var (
		spec  SpecNode
		specs []SpecNode
	)

	token := parser.token
	parser.match(TtKwConst)

	if parser.matchIf(TtLParen) {
		for parser.token.Type == TtIdentifier {
			spec = parser.ConstSpec()
			specs = append(specs, spec)
			parser.match(TtSemicolon)
		}
		parser.match(TtRParen)
	} else {
		spec = parser.ConstSpec()
		specs = append(specs, spec)
	}

	return &GenDeclNode{Tok: token, Specs: specs}
}

func (parser *Parser) ConstSpec() *ValueSpecNode {
	var exprs []ExprNode

	token := parser.token
	idents := parser.Identifiers()
	typeExpr := parser.TypeName()

	parser.match(TtOpAssign)

	expr := parser.BasicLiteral()
	exprs = append(exprs, expr)

	for parser.matchIf(TtComma) {
		expr = parser.BasicLiteral()
		exprs = append(exprs, expr)
	}

	return &ValueSpecNode{
		Token:  token,
		Names:  idents,
		Type:   typeExpr,
		Values: exprs,
	}
}

func (parser *Parser) BreakStatement() StmtNode {
	token := parser.token
	parser.match(TtKwBreak)
	return &BranchStmtNode{Tok: token}
}

func (parser *Parser) IfStatement() *IfStmtNode {
	var (
		statement      StmtNode
		condition      ExprNode
		else_statement StmtNode
	)

	token := parser.token
	parser.match(TtKwIf)
	statement_or_condition := parser.SimpleStatement() // as the set of possible conditions is a subset of the set of possible simple statements

	if parser.matchIf(TtSemicolon) {
		statement = statement_or_condition
		condition = parser.Expression()
	} else {
		expression, ok := statement_or_condition.(*ExprStmtNode)
		if !ok {
			parser.matchError("expected an expression as the if condition")
		}
		condition = expression.Expr
	}

	body := parser.BlockStatement()

	if parser.matchIf(TtKwElse) {
		switch parser.token.Type {
		case TtKwIf:
			else_statement = parser.IfStatement()
		case TtLBrace:
			else_statement = parser.BlockStatement()
		default:
			parser.matchError("expected \"if\" or \"{\" after \"else\"")
		}
	}

	return &IfStmtNode{Tok: token, Init: statement, Cond: condition, Body: body, Else: else_statement}
}

func (parser *Parser) ExprAnd() ExprNode {
	var expr ExprNode = nil
	expr = parser.ExprRelational()
	token := parser.token

	for parser.matchIf(TtOpAnd) {
		rhs := parser.ExprRelational()
		expr = &BinaryExprNode{Tok: token, Lhs: expr, Rhs: rhs}
		token = parser.token
	}

	return expr
}

func (parser *Parser) ExprRelational() ExprNode {
	expr := parser.ExprAdditive()

	for parser.oneOf(
		parser.token.Type,
		TtOpEq,
		TtOpNe,
		TtOpLt,
		TtOpLe,
		TtOpGt,
		TtOpGe,
	) {
		token := parser.token
		parser.match(parser.token.Type)

		rhs := parser.ExprAdditive()

		expr = &BinaryExprNode{
			Tok: token,
			Lhs: expr,
			Rhs: rhs,
		}
	}

	return expr
}

func (parser *Parser) ExprAdditive() ExprNode {
	expr := parser.ExprMultiplicative()

	for parser.oneOf(
		parser.token.Type,
		TtOpAdd,
		TtOpSub,
	) {
		token := parser.token
		parser.match(parser.token.Type)
		rhs := parser.ExprMultiplicative()
		expr = &BinaryExprNode{
			Tok: token,
			Lhs: expr,
			Rhs: rhs,
		}
	}

	return expr
}

func (parser *Parser) ExprMultiplicative() ExprNode {

	expr := parser.ExprUnary()
	for parser.oneOf(
		parser.token.Type,
		TtOpMul,
		TtOpDiv,
		TtOpMod,
	) {
		token := parser.token
		parser.match(parser.token.Type)

		rhs := parser.ExprUnary()

		expr = &BinaryExprNode{
			Tok: token,
			Lhs: expr,
			Rhs: rhs,
		}
	}

	return expr
}

func (parser *Parser) ExprUnary() ExprNode {
	if parser.oneOf(
		parser.token.Type,
		TtOpAdd,
		TtOpSub,
		TtOpNot,
	) {
		token := parser.token
		parser.match(parser.token.Type)

		expr := parser.ExprUnary()

		return &UnaryExprNode{
			Tok:  token,
			Expr: expr,
		}
	}

	return parser.ExprPrimary()
}

func (parser *Parser) ExprPrimary() ExprNode {
	if parser.isBasicLiteral() {
		return parser.BasicLiteral()
	}

	if parser.matchIf(TtLParen) {
		expr := parser.Expression()
		parser.match(TtRParen)
		return expr
	}

	if parser.token.Type != TtIdentifier {
		msg := fmt.Sprintf(
			"Expected a primary expression but found '%s' at position %s.",
			parser.token.Type,
			parser.token.Pos,
		)
		parser.matchError(msg)
	}

	ident, token := parser.QualifiedIdentifier()

	if parser.token.Type == TtLParen {
		parser.match(TtLParen)

		var args []ExprNode
		if parser.token.Type != TtRParen {
			args = parser.Expressions()
		}

		parser.match(TtRParen)

		return &CallExprNode{
			Tok:  token,
			Fun:  ident,
			Args: args,
		}
	}

	return ident
}

func (parser *Parser) QualifiedIdentifier() (ExprNode, Token) {
	first := parser.Identifier()
	var expr ExprNode = first
	token := first.Tok

	if parser.matchIf(TtPeriod) {
		second := parser.Identifier()

		token.Lexeme = token.Lexeme + "." + second.Tok.Lexeme
		token.PosEnd = second.Tok.PosEnd

		expr = &SelectorExprNode{
			Tok:  token,
			Expr: expr,
			Sel:  second,
		}
	}

	return expr, token
}

func (parser *Parser) isBasicLiteral() bool {
	if parser.oneOf(
		parser.token.Type,
		TtInt,
		TtFloat,
		TtString,
	) {
		return true
	}

	return parser.token.Type == TtIdentifier &&
		(parser.token.Lexeme == "true" || parser.token.Lexeme == "false")
}

// Add functions as needed to parse expressions with the precedence (and associativity) of Grobbit operators correct
// (same as in Go). Note that you need to rewrite the grammar for reflecting the correct operator precedence.
