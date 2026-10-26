package toolmade

// antlr4 -Dlanguage=Go -visitor -no-listener -package generated -o ../generated GrobbitParser.g4

import (
	"Grobbit/common"
	"Grobbit/frontend/toolmade/generated"
	"fmt"
)
import "github.com/antlr4-go/antlr/v4"

type GrobbitParserVisitor struct {
	*generated.BaseGrobbitParserVisitor
}

func NewGrobbitParserVisitor() *GrobbitParserVisitor {
	return &GrobbitParserVisitor{&generated.BaseGrobbitParserVisitor{}}
}

func (v *GrobbitParserVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func VisitExpectingType[T any](v *GrobbitParserVisitor, tree antlr.ParseTree) T {
	value := tree.Accept(v)
	switch ret := value.(type) {
	case T:
		return ret
	default:
		panic(fmt.Sprintf("VisitType[%T] unexpected.", value))
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

// VisitProgram implements:
//
// program:	TtKwPackage pkg_name TtSemicolon (import_decl TtSemicolon)* (top_decl TtSemicolon)* EOF
//
//	;
func (v *GrobbitParserVisitor) VisitProgram(ctx *generated.ProgramContext) interface{} {
	var (
		imports []*common.ImportSpecNode
		decls   []common.DeclNode
	)
	token := mapToken(ctx.GetStart())
	ident := VisitExpectingType[*common.IdentifierNode](v, ctx.Pkg_name())
	for _, node := range ctx.AllImport_decl() {
		imps := VisitExpectingType[[]*common.ImportSpecNode](v, node)
		imports = append(imports, imps...)
	}
	for _, node := range ctx.AllTop_decl() {
		decl := VisitExpectingType[common.DeclNode](v, node)
		decls = append(decls, decl)
	}
	return &common.FileNode{Tok: token, Name: ident, Imports: imports, Decls: decls}
}

// VisitTop_decl implements:
//
// top_decl:	const_decl | var_decl | func_decl
//
//	;
func (v *GrobbitParserVisitor) VisitTop_decl(ctx *generated.Top_declContext) interface{} {
	var decl common.DeclNode
	if ctx.Const_decl() != nil {
		decl = VisitExpectingType[*common.GenDeclNode](v, ctx.Const_decl())
	} else if ctx.Var_decl() != nil {
		decl = VisitExpectingType[*common.GenDeclNode](v, ctx.Var_decl())
	} else {
		decl = VisitExpectingType[*common.FuncDeclNode](v, ctx.Func_decl())
	}
	return decl
}

// VisitPkg_name implements:
//
// pkg_name:	identifier
//
//	;
func (v *GrobbitParserVisitor) VisitPkg_name(ctx *generated.Pkg_nameContext) interface{} {
	return VisitExpectingType[*common.IdentifierNode](v, ctx.Identifier())
}

// VisitImport_decl implements:
//
// import_decl:	TtKwImport (import_spec | TtLParen (import_spec TtSemicolon)* TtRParen)
//
//	;
func (v *GrobbitParserVisitor) VisitImport_decl(ctx *generated.Import_declContext) interface{} {
	var imports []*common.ImportSpecNode
	for _, node := range ctx.AllImport_spec() {
		imp := VisitExpectingType[*common.ImportSpecNode](v, node)
		imports = append(imports, imp)
	}
	return imports
}

// VisitImport_spec implements:
//
// import_spec: (pkg_name)? TtString
//
//	;
func (v *GrobbitParserVisitor) VisitImport_spec(ctx *generated.Import_specContext) interface{} {
	var ident *common.IdentifierNode = nil
	token := mapToken(ctx.GetStart())
	if ctx.Pkg_name() != nil {
		ident = VisitExpectingType[*common.IdentifierNode](v, ctx.Pkg_name())
	}
	path := ctx.TtString().GetText()
	return &common.ImportSpecNode{Tok: token, Name: ident, Path: path}
}

// VisitConst_decl implements:
//
//	const_decl:	TtKwVar (const_spec | TtLParen (const_spec TtSemicolon)* TtRParen)
//				;
func (v *GrobbitParserVisitor) VisitConst_decl(ctx *generated.Const_declContext) interface{} {
	var (
		specs []common.SpecNode
	)
	token := mapToken(ctx.GetStart())
	for _, node := range ctx.AllConst_spec() {
		spec := VisitExpectingType[common.SpecNode](v, node)
		specs = append(specs, spec)
	}
	return &common.GenDeclNode{Tok: token, Specs: specs}
}

// VisitConst_spec implements:
//
// const_spec:	identifiers type TtOpAssign basic_literals
//
//	;
func (v *GrobbitParserVisitor) VisitConst_spec(ctx *generated.Const_specContext) interface{} {
	var (
		typ   common.ExprNode
		exprs []common.ExprNode
	)
	token := mapToken(ctx.GetStart())
	idents := VisitExpectingType[[]*common.IdentifierNode](v, ctx.Identifiers())
	typ = VisitExpectingType[common.ExprNode](v, ctx.Type_())
	exprs = VisitExpectingType[[]common.ExprNode](v, ctx.Basic_literals())
	return &common.ValueSpecNode{Token: token, Names: idents, Type: typ, Values: exprs}
}

// VisitVar_decl implements:
//
//	var_decl:   TtKwVar (var_spec | TtLParen (var_spec TtSemicolon)* TtRParen)
//				;
func (v *GrobbitParserVisitor) VisitVar_decl(ctx *generated.Var_declContext) interface{} {
	var (
		specs []common.SpecNode
	)
	token := mapToken(ctx.GetStart())
	for _, node := range ctx.AllVar_spec() {
		spec := VisitExpectingType[common.SpecNode](v, node)
		specs = append(specs, spec)
	}
	return &common.GenDeclNode{Tok: token, Specs: specs}
}

// VisitVar_spec implements:
//
// var_spec:       	identifiers type (TtOpAssign basic_literals)?
//
//	;
func (v *GrobbitParserVisitor) VisitVar_spec(ctx *generated.Var_specContext) interface{} {
	var (
		typ   common.ExprNode
		exprs []common.ExprNode
	)
	token := mapToken(ctx.GetStart())
	idents := VisitExpectingType[[]*common.IdentifierNode](v, ctx.Identifiers())
	typ = VisitExpectingType[common.ExprNode](v, ctx.Type_())
	if ctx.Basic_literals() != nil {
		exprs = VisitExpectingType[[]common.ExprNode](v, ctx.Basic_literals())
	}
	return &common.ValueSpecNode{Token: token, Names: idents, Type: typ, Values: exprs}
}

// VisitBasic_literals implements:
//
//	identifiers:	basic_literal (TtComma basic_literal)*
//					;
func (v *GrobbitParserVisitor) VisitBasic_literals(ctx *generated.Basic_literalsContext) interface{} {
	var literals []common.ExprNode
	for _, node := range ctx.AllBasic_literal() {
		literal := VisitExpectingType[common.ExprNode](v, node)
		literals = append(literals, literal)
	}
	return literals
}

// VisitBasic_literal implements:
//
//		basic_literal:  	TtInt | TtFloat | TtBool
//	 	              	;
func (v *GrobbitParserVisitor) VisitBasic_literal(ctx *generated.Basic_literalContext) interface{} {
	token := mapToken(ctx.GetStart())
	return &common.LiteralNode{Tok: token}
}

// VisitFunc_decl implements:
//
// func_decl:	TtKwFunc identifier TtLParen (param_decl)? TtRParen (type)? block
//
//	;
func (v *GrobbitParserVisitor) VisitFunc_decl(ctx *generated.Func_declContext) interface{} {
	var funcType common.FuncType
	token := mapToken(ctx.GetStart())
	ident := VisitExpectingType[*common.IdentifierNode](v, ctx.Identifier())
	if ctx.Param_decl() != nil {
		funcType.Params = VisitExpectingType[[]*common.Field](v, ctx.Param_decl())
	}
	if ctx.Type_() != nil {
		funcType.Type = VisitExpectingType[common.ExprNode](v, ctx.Type_())
	}
	body := VisitExpectingType[*common.BlockStmtNode](v, ctx.Block())
	return &common.FuncDeclNode{Tok: token, Name: ident,
		Type: &funcType, Body: body}
}

// VisitParam_decl implements:
//
//	param_decl:	param_spec (TtComma param_spec)*
//				;
func (v *GrobbitParserVisitor) VisitParam_decl(ctx *generated.Param_declContext) interface{} {
	var specs []*common.Field
	for _, node := range ctx.AllParam_spec() {
		spec := VisitExpectingType[*common.Field](v, node)
		specs = append(specs, spec)
	}
	return specs
}

// VisitParam_spec implements:
//
//	param_spec: 	identifiers type
func (v *GrobbitParserVisitor) VisitParam_spec(ctx *generated.Param_specContext) interface{} {
	names := VisitExpectingType[[]*common.IdentifierNode](v, ctx.Identifiers())
	expr := VisitExpectingType[common.ExprNode](v, ctx.Type_())
	return &common.Field{Names: names, Type: expr}
}

// VisitIdentifiers implements:
//
//	identifiers:	identifier (TtComma identifier)*
//					;
func (v *GrobbitParserVisitor) VisitIdentifiers(ctx *generated.IdentifiersContext) interface{} {
	var idents []*common.IdentifierNode
	for _, node := range ctx.AllIdentifier() {
		ident := VisitExpectingType[*common.IdentifierNode](v, node)
		idents = append(idents, ident)
	}
	return idents
}

// VisitIdentifier implements:
//
//	identifier:	TtIdentifier
//	      		;
func (v *GrobbitParserVisitor) VisitIdentifier(ctx *generated.IdentifierContext) interface{} {
	token := mapToken(ctx.GetStart())
	return &common.IdentifierNode{Tok: token}
}

// VisitType implements:
//
// type:	identifier
//
//	;
func (v *GrobbitParserVisitor) VisitType(ctx *generated.TypeContext) interface{} {
	return VisitExpectingType[*common.IdentifierNode](v, ctx.Identifier())
}

// VisitBlock implements:
//
//	block:      	TtLBrace (statement TtSemicolon)* TtRBrace
//					;
func (v *GrobbitParserVisitor) VisitBlock(ctx *generated.BlockContext) interface{} {
	var stmts []common.StmtNode
	token := mapToken(ctx.GetStart())
	for _, node := range ctx.AllStatement() {
		stmt := VisitExpectingType[common.StmtNode](v, node)
		stmts = append(stmts, stmt)
	}
	return &common.BlockStmtNode{Tok: token, List: stmts}
}

//////////////////////////////////////// Statements //////////////////////////////////////////////////////

// VisitStatement implements:
//
//	  statement:      decl_stmt
//	  	        	| simple_stmt
//				    | break_stmt
//	         	    | continue_stmt
//	             	| return_stmt
//	  	            | block
//				    | if_stmt
//	                | for_stmt
//	                ;
func (v *GrobbitParserVisitor) VisitStatement(ctx *generated.StatementContext) interface{} {
	var stmt common.StmtNode
	if ctx.Decl_stmt() != nil {
		stmt = VisitExpectingType[*common.DeclStmtNode](v, ctx.Decl_stmt())
	} else if ctx.Simple_stmt() != nil {
		stmt = VisitExpectingType[common.StmtNode](v, ctx.Simple_stmt())
	} else if ctx.Break_stmt() != nil {
		stmt = VisitExpectingType[*common.BranchStmtNode](v, ctx.Break_stmt())
	} else if ctx.Continue_stmt() != nil {
		stmt = VisitExpectingType[*common.BranchStmtNode](v, ctx.Continue_stmt())
	} else if ctx.Return_stmt() != nil {
		stmt = VisitExpectingType[*common.ReturnStmtNode](v, ctx.Return_stmt())
	} else if ctx.Block() != nil {
		stmt = VisitExpectingType[*common.BlockStmtNode](v, ctx.Block())
	} else if ctx.If_stmt() != nil {
		stmt = VisitExpectingType[*common.IfStmtNode](v, ctx.If_stmt())
	} else if ctx.For_stmt() != nil {
		stmt = VisitExpectingType[*common.ForStmtNode](v, ctx.For_stmt())
	} else {
		panic(v)
	}
	return stmt
}

// VisitDecl_stmt implements:
//
//	decl_stmt:        const_decl | var_decl
//					;
func (v *GrobbitParserVisitor) VisitDecl_stmt(ctx *generated.Decl_stmtContext) interface{} {
	var decl *common.GenDeclNode
	token := mapToken(ctx.GetStart())
	if ctx.Const_decl() != nil {
		decl = VisitExpectingType[*common.GenDeclNode](v, ctx.Const_decl())
	} else {
		decl = VisitExpectingType[*common.GenDeclNode](v, ctx.Var_decl())
	}
	return &common.DeclStmtNode{Tok: token, Decl: decl}
}

// VisitSimple_stmt implements:
//
//	simple_stmt:      expr_stmt
//					| shortvar_decl
//					| assign_stmt
func (v *GrobbitParserVisitor) VisitSimple_stmt(ctx *generated.Simple_stmtContext) interface{} {
	var stmt common.StmtNode = nil
	if ctx.Expr_stmt() != nil {
		stmt = VisitExpectingType[*common.ExprStmtNode](v, ctx.Expr_stmt())
	} else if ctx.Shortvar_decl() != nil {
		stmt = VisitExpectingType[*common.AssignStmtNode](v, ctx.Shortvar_decl())
	} else if ctx.Assign_stmt() != nil {
		stmt = VisitExpectingType[*common.AssignStmtNode](v, ctx.Assign_stmt())
	}
	return stmt
}

// VisitReturn_stmt implements:
//
//	return_stmt:    TtKwReturn (expression)?
//					;
func (v *GrobbitParserVisitor) VisitReturn_stmt(ctx *generated.Return_stmtContext) interface{} {
	var result common.ExprNode
	token := mapToken(ctx.GetStart())
	if ctx.Expression() != nil {
		result = VisitExpectingType[common.ExprNode](v, ctx.Expression())
	}
	return &common.ReturnStmtNode{Token: token, Result: result}
}

// VisitContinue_stmt implements:
//
//	continue_stmt:  TtKwContinue
//					;
func (v *GrobbitParserVisitor) VisitContinue_stmt(ctx *generated.Continue_stmtContext) interface{} {
	tk := ctx.TtKwContinue().GetSymbol()
	return &common.BranchStmtNode{Tok: mapToken(tk)}
}

// VisitBreak_stmt implements:
//
//	break_stmt:     TtKwBreak
//	;
func (v *GrobbitParserVisitor) VisitBreak_stmt(ctx *generated.Break_stmtContext) interface{} {
	tk := ctx.TtKwBreak().GetSymbol()
	return &common.BranchStmtNode{Tok: mapToken(tk)}
}

// VisitIf_stmt implements:
//
// if_stmt:	TtKwIf (simple_stmt TtSemicolon)? condition block (TtKwElse (if_stmt | block))?
//
//	;
func (v *GrobbitParserVisitor) VisitIf_stmt(ctx *generated.If_stmtContext) interface{} {
	var (
		init_stmt common.StmtNode = nil
		else_stmt common.StmtNode = nil
	)

	token := mapToken(ctx.GetStart())
	if ctx.Simple_stmt() != nil {
		init_stmt = VisitExpectingType[common.StmtNode](v, ctx.Simple_stmt())
	}
	cond := VisitExpectingType[common.ExprNode](v, ctx.Condition())
	block := VisitExpectingType[*common.BlockStmtNode](v, ctx.Block(0))
	if ctx.TtKwElse() != nil {
		if ctx.If_stmt() != nil {
			else_stmt = VisitExpectingType[common.StmtNode](v, ctx.If_stmt())
		} else {
			else_stmt = VisitExpectingType[common.StmtNode](v, ctx.Block(1))
		}
	}

	return &common.IfStmtNode{Tok: token, Init: init_stmt, Cond: cond, Body: block, Else: else_stmt}
}

type ForClause struct {
	cond                 common.ExprNode
	init_stmt, post_stmt common.StmtNode
}

// VisitFor_stmt implements:
//
//	for_stmt:    TtKwFor (condition | for_clause)? block
//				;
func (v *GrobbitParserVisitor) VisitFor_stmt(ctx *generated.For_stmtContext) interface{} {
	var clause = ForClause{nil, nil, nil}
	token := mapToken(ctx.GetStart())
	if ctx.Condition() != nil {
		clause.cond = VisitExpectingType[common.ExprNode](v, ctx.Condition())
	} else if ctx.For_clause() != nil {
		clause = VisitExpectingType[ForClause](v, ctx.For_clause())
	}
	block := VisitExpectingType[*common.BlockStmtNode](v, ctx.Block())
	return &common.ForStmtNode{Tok: token, Init: clause.init_stmt, Cond: clause.cond,
		Post: clause.post_stmt, Body: block}
}

// VisitFor_clause implements:
//
//	for_clause:     simple_stmt? TtSemicolon condition? TtSemicolon simple_stmt?
func (v *GrobbitParserVisitor) VisitFor_clause(ctx *generated.For_clauseContext) interface{} {
	var clause = ForClause{nil, nil, nil}

	if initCtx := ctx.GetInit(); initCtx != nil {
		clause.init_stmt = VisitExpectingType[common.StmtNode](v, initCtx)
	}
	if ctx.Condition() != nil {
		clause.cond = VisitExpectingType[common.ExprNode](v, ctx.Condition())
	}
	if postCtx := ctx.GetPost(); postCtx != nil {
		clause.post_stmt = VisitExpectingType[common.StmtNode](v, postCtx)
	}
	return clause
}

// VisitCondition implements:
//
// condition:     expression
//
//	;
func (v *GrobbitParserVisitor) VisitCondition(ctx *generated.ConditionContext) interface{} {
	return VisitExpectingType[common.ExprNode](v, ctx.Expression())
}

// VisitShortvar_decl implements:
//
//		shortvar_decl:  identifiers TtOpDefine expressions
//	                 ;
func (v *GrobbitParserVisitor) VisitShortvar_decl(ctx *generated.Shortvar_declContext) interface{} {
	var lhs []common.ExprNode
	token := mapToken(ctx.TtOpDefine().GetSymbol())
	idents := VisitExpectingType[[]*common.IdentifierNode](v, ctx.Identifiers())
	for _, ident := range idents {
		lhs = append(lhs, ident)
	}
	rhs := VisitExpectingType[[]common.ExprNode](v, ctx.Expressions())
	return &common.AssignStmtNode{Tok: token, Lhs: lhs, Rhs: rhs}
	// NOTE Maybe have lhs of AssignStmtNode be []IdentifierNode?
}

// VisitAssign_stmt implements:
//
//	assign_stmt:    identifiers TtOpAssign expressions
//					;
func (v *GrobbitParserVisitor) VisitAssign_stmt(ctx *generated.Assign_stmtContext) interface{} {
	var lhs []common.ExprNode
	token := mapToken(ctx.TtOpAssign().GetSymbol())
	idents := VisitExpectingType[[]*common.IdentifierNode](v, ctx.Identifiers())
	for _, ident := range idents {
		lhs = append(lhs, ident)
	}
	rhs := VisitExpectingType[[]common.ExprNode](v, ctx.Expressions())
	return &common.AssignStmtNode{Tok: token, Lhs: lhs, Rhs: rhs}
	// NOTE Maybe have lhs of AssignStmtNode be []IdentifierNode?
}

// VisitExpr_stmt implements:
//
//	expr_stmt:      call_expr
//					;
func (v *GrobbitParserVisitor) VisitExpr_stmt(ctx *generated.Expr_stmtContext) interface{} {
	token := mapToken(ctx.GetStart())
	expr := VisitExpectingType[common.ExprNode](v, ctx.Call_expr())
	return &common.ExprStmtNode{Tok: token, Expr: expr}
}

//////////////////////////////////////// Expressions //////////////////////////////////////////////////////

// VisitExpressions implements:
//
//	expressions:    expression (TtComma expression)*
//	;
func (v *GrobbitParserVisitor) VisitExpressions(ctx *generated.ExpressionsContext) interface{} {
	var exprs []common.ExprNode
	for _, node := range ctx.AllExpression() {
		expr := VisitExpectingType[common.ExprNode](v, node)
		exprs = append(exprs, expr)
	}
	return exprs
}

// VisitExpression implements:
//
// expression: 	  unary_expr
//
//	| expression op=(TtOpMul | TtOpDiv | TtOpMod) expression
//	| expression op=(TtOpAdd | TtOpSub) expression
//	| expression op=(TtOpEq | TtOpNe | TtOpLt | TtOpLe | TtOpGt | TtOpGe) expression
//	| expression op=TtOpAnd expression
//	| expression op=TtOpOr expression
//	;
func (v *GrobbitParserVisitor) VisitExpression(ctx *generated.ExpressionContext) interface{} {
	if ctx.Unary_expr() != nil {
		return VisitExpectingType[common.ExprNode](v, ctx.Unary_expr())
	}
	lhs := VisitExpectingType[common.ExprNode](v, ctx.Expression(0))
	token := mapToken(ctx.GetOp())
	rhs := VisitExpectingType[common.ExprNode](v, ctx.Expression(1))

	return &common.BinaryExprNode{Tok: token, Lhs: lhs, Rhs: rhs}
}

// VisitUnary_expr implements:
//
//	unary_expr:       primary_expr
//	                | op=(TtOpAdd | TtOpSub | TtOpNot) unary_expr
func (v *GrobbitParserVisitor) VisitUnary_expr(ctx *generated.Unary_exprContext) interface{} {
	if ctx.Primary_expr() != nil {
		return VisitExpectingType[common.ExprNode](v, ctx.Primary_expr())
	}
	token := mapToken(ctx.GetOp())
	rhs := VisitExpectingType[common.ExprNode](v, ctx.Unary_expr())
	return &common.UnaryExprNode{Tok: token, Expr: rhs}
}

// VisitPrimary_expr implements:
//
// primary_expr:      basic_literal
//
//	| (pkg_name TtPeriod)? identifier
//	| call_expr
//	| TtLParen expression TtRParen
//	;
func (v *GrobbitParserVisitor) VisitPrimary_expr(ctx *generated.Primary_exprContext) interface{} {
	var expr common.ExprNode
	if ctx.Basic_literal() != nil {
		expr = VisitExpectingType[common.ExprNode](v, ctx.Basic_literal())
	} else if ctx.Call_expr() != nil {
		expr = VisitExpectingType[common.ExprNode](v, ctx.Call_expr())
	} else if ctx.TtLParen() != nil {
		expr = VisitExpectingType[common.ExprNode](v, ctx.Expression())
	} else {
		expr = VisitExpectingType[*common.IdentifierNode](v, ctx.Identifier())
		if ctx.Pkg_name() != nil {
			token := mapToken(ctx.TtPeriod().GetSymbol())
			ident := VisitExpectingType[*common.IdentifierNode](v, ctx.Pkg_name())
			expr = &common.SelectorExprNode{Tok: token, Expr: expr, Sel: ident}
		}
	}
	return expr
}

// VisitCall_expr implements:
//
// call_expr:       (pkg_name TtPeriod)? identifier TtLParen (expressions)? TtRParen
//
//	;
func (v *GrobbitParserVisitor) VisitCall_expr(ctx *generated.Call_exprContext) interface{} {
	var expr common.ExprNode
	//token := mapToken(ctx.TtLParen().GetSymbol())
	token := mapToken(ctx.Identifier().GetStart())
	expr = VisitExpectingType[*common.IdentifierNode](v, ctx.Identifier())
	if ctx.Pkg_name() != nil {
		token := mapToken(ctx.TtPeriod().GetSymbol())
		ident := VisitExpectingType[*common.IdentifierNode](v, ctx.Pkg_name())
		expr = &common.SelectorExprNode{Tok: token, Expr: expr, Sel: ident}
	}
	var args []common.ExprNode
	if ctx.Expressions() != nil {
		args = VisitExpectingType[[]common.ExprNode](v, ctx.Expressions())
	}
	return &common.CallExprNode{Tok: token, Fun: expr, Args: args}
}
