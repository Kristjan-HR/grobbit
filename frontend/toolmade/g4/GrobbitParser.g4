parser grammar GrobbitParser;
options { tokenVocab=GrobbitLexer; }

///////////////////////////////////////////////// DECLARATIONS /////////////////////////////////////////////////

program:        TtKwPackage pkg_name TtSemicolon (import_decl TtSemicolon)* (top_decl TtSemicolon)* EOF
                ;

top_decl:       const_decl | var_decl | func_decl
                ;

pkg_name:       identifier
                ;

import_decl:    TtKwImport ( import_spec | TtLParen (import_spec TtSemicolon)* TtRParen)
                ;

import_spec:    (pkg_name)? TtString
                ;

const_decl:     TtKwConst (const_spec | TtLParen (const_spec TtSemicolon)* TtRParen)
                ;

const_spec:     identifiers type TtOpAssign basic_literals
                ;

var_decl:       TtKwVar (var_spec | TtLParen (var_spec TtSemicolon)* TtRParen)
                ;

var_spec:       identifiers type (TtOpAssign basic_literals)?
                ;

basic_literals: basic_literal (TtComma basic_literal)*
                ;

basic_literal:  TtInt | TtFloat | TtBool
                ;

func_decl:      TtKwFunc identifier TtLParen (param_decl)? TtRParen (type)? block
                ;

param_decl:     param_spec (TtComma param_spec)*
                ;

param_spec:     identifiers type
                ;

identifiers:    identifier (TtComma identifier)*
                ;

identifier:     TtIdentifier
                ;

type:           identifier
                ;

block:           TtLBrace (statement TtSemicolon)* TtRBrace
                ;

///////////////////////////////////////////////// STATEMENTS /////////////////////////////////////////////////

statement:        decl_stmt
                | simple_stmt
                | return_stmt
                | break_stmt
                | continue_stmt
                | block
                | if_stmt
                | for_stmt
                ;

decl_stmt:        const_decl | var_decl
                ;

simple_stmt:      expr_stmt
                | assign_stmt
                | shortvar_decl
                ;

return_stmt:    TtKwReturn (expression)?
                ;

continue_stmt:  TtKwContinue
                ;

break_stmt:     TtKwBreak
                ;

if_stmt:        TtKwIf (simple_stmt TtSemicolon)? condition block (TtKwElse (if_stmt | block))?
                ;

for_stmt:       TtKwFor (condition | for_clause)? block
                ;

for_clause:     (init=simple_stmt)? TtSemicolon condition? TtSemicolon (post=simple_stmt)?
                ;

condition:      expression
                ;

shortvar_decl:  identifiers TtOpDefine expressions
                ;

assign_stmt:    identifiers TtOpAssign  expressions
                ;

expr_stmt:      call_expr
                ;

///////////////////////////////////////////////// EXPRESSIONS /////////////////////////////////////////////////

expressions:    expression (TtComma expression)*
                ;

expression:       unary_expr
                | expression op=(TtOpMul | TtOpDiv | TtOpMod) expression
                | expression op=(TtOpAdd | TtOpSub) expression
                | expression op=(TtOpEq | TtOpNe | TtOpLt | TtOpLe | TtOpGt | TtOpGe) expression
                | expression op=TtOpAnd expression
                | expression op=TtOpOr expression
                ;

unary_expr:       primary_expr
                | op=(TtOpAdd | TtOpSub | TtOpNot) unary_expr
                ;

primary_expr:     basic_literal
                | (pkg_name TtPeriod)? identifier
                | call_expr
                | TtLParen expression TtRParen
                ;

call_expr:      (pkg_name TtPeriod)? identifier TtLParen (expressions)? TtRParen
                ;
