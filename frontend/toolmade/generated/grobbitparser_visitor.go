// Code generated from GrobbitParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // GrobbitParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrobbitParser.
type GrobbitParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrobbitParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GrobbitParser#top_decl.
	VisitTop_decl(ctx *Top_declContext) interface{}

	// Visit a parse tree produced by GrobbitParser#pkg_name.
	VisitPkg_name(ctx *Pkg_nameContext) interface{}

	// Visit a parse tree produced by GrobbitParser#import_decl.
	VisitImport_decl(ctx *Import_declContext) interface{}

	// Visit a parse tree produced by GrobbitParser#import_spec.
	VisitImport_spec(ctx *Import_specContext) interface{}

	// Visit a parse tree produced by GrobbitParser#const_decl.
	VisitConst_decl(ctx *Const_declContext) interface{}

	// Visit a parse tree produced by GrobbitParser#const_spec.
	VisitConst_spec(ctx *Const_specContext) interface{}

	// Visit a parse tree produced by GrobbitParser#var_decl.
	VisitVar_decl(ctx *Var_declContext) interface{}

	// Visit a parse tree produced by GrobbitParser#var_spec.
	VisitVar_spec(ctx *Var_specContext) interface{}

	// Visit a parse tree produced by GrobbitParser#basic_literals.
	VisitBasic_literals(ctx *Basic_literalsContext) interface{}

	// Visit a parse tree produced by GrobbitParser#basic_literal.
	VisitBasic_literal(ctx *Basic_literalContext) interface{}

	// Visit a parse tree produced by GrobbitParser#func_decl.
	VisitFunc_decl(ctx *Func_declContext) interface{}

	// Visit a parse tree produced by GrobbitParser#param_decl.
	VisitParam_decl(ctx *Param_declContext) interface{}

	// Visit a parse tree produced by GrobbitParser#param_spec.
	VisitParam_spec(ctx *Param_specContext) interface{}

	// Visit a parse tree produced by GrobbitParser#identifiers.
	VisitIdentifiers(ctx *IdentifiersContext) interface{}

	// Visit a parse tree produced by GrobbitParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by GrobbitParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by GrobbitParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GrobbitParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by GrobbitParser#decl_stmt.
	VisitDecl_stmt(ctx *Decl_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#simple_stmt.
	VisitSimple_stmt(ctx *Simple_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#continue_stmt.
	VisitContinue_stmt(ctx *Continue_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#break_stmt.
	VisitBreak_stmt(ctx *Break_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#for_stmt.
	VisitFor_stmt(ctx *For_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#for_clause.
	VisitFor_clause(ctx *For_clauseContext) interface{}

	// Visit a parse tree produced by GrobbitParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by GrobbitParser#shortvar_decl.
	VisitShortvar_decl(ctx *Shortvar_declContext) interface{}

	// Visit a parse tree produced by GrobbitParser#assign_stmt.
	VisitAssign_stmt(ctx *Assign_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#expr_stmt.
	VisitExpr_stmt(ctx *Expr_stmtContext) interface{}

	// Visit a parse tree produced by GrobbitParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by GrobbitParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GrobbitParser#unary_expr.
	VisitUnary_expr(ctx *Unary_exprContext) interface{}

	// Visit a parse tree produced by GrobbitParser#primary_expr.
	VisitPrimary_expr(ctx *Primary_exprContext) interface{}

	// Visit a parse tree produced by GrobbitParser#call_expr.
	VisitCall_expr(ctx *Call_exprContext) interface{}
}
