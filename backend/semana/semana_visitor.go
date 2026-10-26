package semana

import (
	. "Grobbit/common"
	"fmt"
	"strings"
)

type SAFlags uint16

const (
	InVarDecl SAFlags = 1 << iota
	InConstDecl
	InFuncDecl
	InParamDecl
	InIdDecl
	InType
	InExpr
)

var flagNames = map[SAFlags]string{
	InVarDecl:   "InVarDecl",
	InConstDecl: "InConstDecl",
	InFuncDecl:  "InFuncDecl",
	InParamDecl: "InParamDecl",
	InIdDecl:    "InIdDecl",
	InType:      "InType",
	InExpr:      "InExpr",
}

func (f SAFlags) String() string {
	var sb strings.Builder
	for k, v := range flagNames {
		if f&k != 0 {
			sb.WriteString(" ")
			sb.WriteString(v)
		}
	}
	return sb.String()
}

type SAVisitor struct {
	VisitorWalker
	errorHandler   ErrorHandler
	universalScope *Scope
	packageScope   *Scope
	fileScope      *Scope
	currentScope   *Scope
	output         strings.Builder
	flags          SAFlags
	currFuncName   string
	enterFuncBlock bool
}

func (v *SAVisitor) printf(format string, a ...any) {
	builder := strings.Builder{}
	for i := 0; i < v.Level; i++ {
		builder.WriteString(". ")
	}
	builder.WriteString(fmt.Sprintf(format, a...))
	v.output.WriteString(builder.String())
}

func (v *SAVisitor) visit(node Node) {
	err := v.VisitDispatcher(v, node)
	if err != nil {
		v.printf("%s\n", err.Error())
	}
}

func (v *SAVisitor) VisitIdentifierNode(node *IdentifierNode) {
	const msgUndefined = "undeclared: %s"
	obj, _ := Resolve(v.currentScope, node, v.flags&InIdDecl != 0)
	if obj == nil {
		if v.flags&InIdDecl == 0 {
			v.logError(node.Tok.Pos, msgUndefined, node.Tok.Lexeme)
		}
		node.SetType(NewInvalidType())
		return
	}
	// NOTE: Check object is a TypeName if InType flag is set.
	node.SetType(obj.Type())
}

func (v *SAVisitor) VisitLiteralNode(node *LiteralNode) {
	if node.Tok.Type == TtInt {
		node.SetType(NewIntType())
	} else if node.Tok.Type == TtBool {
		node.SetType(NewBoolType())
	} else if node.Tok.Type == TtFloat {
		node.SetType(NewFloatType())
	} else if node.Tok.Type == TtChar {
		node.SetType(NewRuneType())
	} else if node.Tok.Type == TtString {
		node.SetType(NewStringType())
	} else {
		panic("unhandled literal type in type checking")
	}
}

func (v *SAVisitor) VisitSelectorExprNode(node *SelectorExprNode) {
	v.visit(node.Expr)
	v.visit(node.Sel)
}

func (v *SAVisitor) VisitCallExprNode(node *CallExprNode) {

	// Small hack ...
	reportUndefinedError := true
	if id, ok := node.Fun.(*IdentifierNode); ok {
		if id.Tok.Lexeme == "Print" || id.Tok.Lexeme == "PrintN" ||
			id.Tok.Lexeme == "Print_F" || id.Tok.Lexeme == "PrintN_F" {
			reportUndefinedError = false
		}
	}

	const msgUndefined = "undeclared function: %s"
	if reportUndefinedError {
		v.visit(node.Fun)
	}
	for _, arg := range node.Args {
		v.visit(arg)
	}
	id := NewIdentifier(TtIdentifier, node.Tok.Lexeme)
	obj, _ := Resolve(v.currentScope, id, false)
	if obj == nil {
		if reportUndefinedError {
			v.logError(node.Tok.Pos, msgUndefined, node.Tok.Lexeme)
		}
		node.SetType(NewInvalidType())
		return
	}
	if typ, ok := obj.Type().(*SignatureType); ok {
		if len(typ.Results) > 0 {
			node.SetType(typ.Results[0])
			return
		}
	}
	node.SetType(NewInvalidType())
}

func (v *SAVisitor) logError(pos Position, format string, a ...any) {
	str := fmt.Sprintf(format, a...)
	v.errorHandler(pos, str)
}

func (v *SAVisitor) VisitUnaryExprNode(node *UnaryExprNode) {
	const msgInvalidOp1 = "invalid operation: operator '%s' not defined on '%s'"
	const msgInvalidOp2 = "invalid operation: cannot indirect '%s' on type '%s'"
	const msgInvalidOp3 = "invalid operation: cannot take address"
	v.visit(node.Expr)
	exprType := node.Expr.GetType()
	switch node.Tok.Lexeme {
	case "+", "-":
		if !TypeIsOneOf(exprType, TypeInt|TypeFloat) {
			v.logError(node.Tok.Pos, msgInvalidOp1, node.Tok.Lexeme, exprType.String())
		}
		node.SetType(exprType)
	case "!":
		if !TypeIsOneOf(exprType, TypeBool) {
			v.logError(node.Tok.Pos, msgInvalidOp1, node.Tok.Lexeme, exprType.String())
		}
		node.SetType(NewBoolType())
	case "*": // Dereferencing is handled in StarExprNode, but also include in case you used UnaryExprNode.
		// Ensure is (non-nil) pointer type
		if !TypeIsOneOf(exprType, TypePointer) || TypeIsNilPointer(exprType) {
			v.logError(node.Tok.Pos, msgInvalidOp2, node.Tok.Lexeme, exprType.String())
			node.SetType(NewInvalidType())
		} else {
			node.SetType(exprType.(*PointerType).GetBaseType())
		}
	case "&":
		if !IsAddressable(v.currentScope, node.Expr) {
			v.logError(node.Tok.Pos, msgInvalidOp3)
		}
		node.SetType(NewPointerType(exprType))
	default:
		panic("unhandled unary expression operator: " + node.Tok.Lexeme)
	}

}

func (v *SAVisitor) VisitBinaryExprNode(node *BinaryExprNode) {
	const msgInvalidOp1 = "invalid operation: operator '%s' not defined on types '%s' and '%s'"
	v.visit(node.Lhs)
	v.visit(node.Rhs)
	exprTypeLhs, exprTypeRhs := node.Lhs.GetType(), node.Rhs.GetType()
	switch node.Tok.Lexeme {
	case "&&", "||":
		if !TypeIsOneOf(exprTypeLhs, TypeBool) && !TypeIsOneOf(exprTypeRhs, TypeBool) {
			v.logError(node.Tok.Pos, msgInvalidOp1, node.Tok.Lexeme, exprTypeLhs.String(), exprTypeRhs.String())
		}
		node.SetType(NewBoolType())
	case "+", "-", "*", "/":
		if !TypeIsOneOf(exprTypeLhs, TypeInt|TypeFloat) && !TypesAreIdentical(exprTypeLhs, exprTypeRhs) {
			v.logError(node.Tok.Pos, msgInvalidOp1, node.Tok.Lexeme, exprTypeLhs.String(), exprTypeRhs.String())
		}
		node.SetType(exprTypeLhs)
	case "%":
		if !TypeIsOneOf(exprTypeLhs, TypeInt) && !TypeIsOneOf(exprTypeRhs, TypeInt) {
			v.logError(node.Tok.Pos, msgInvalidOp1, node.Tok.Lexeme, exprTypeLhs.String(), exprTypeRhs.String())
		}
		node.SetType(NewIntType())
	case ">=", ">", "<", "<=":
		if !TypeIsOneOf(exprTypeLhs, TypeInt|TypeFloat|TypeString|TypeRune) && !TypesAreComparable(exprTypeLhs, exprTypeRhs) {
			v.logError(node.Tok.Pos, msgInvalidOp1, node.Tok.Lexeme, exprTypeLhs.String(), exprTypeRhs.String())
		}
		node.SetType(NewBoolType())
	case "==", "!=":
		if !TypeIsOneOf(exprTypeLhs, TypeInt|TypeFloat|TypeString|TypeRune|TypePointer) && !TypesAreComparable(exprTypeLhs, exprTypeRhs) {
			v.logError(node.Tok.Pos, msgInvalidOp1, node.Tok.Lexeme, exprTypeLhs.String(), exprTypeRhs.String())
		}
		node.SetType(NewBoolType())
	default:
		panic("unhandled binary expression operator: " + node.Tok.Lexeme)
	}

}

func (v *SAVisitor) VisitDeclStmtNode(node *DeclStmtNode) {
	v.visit(node.Decl)
}

func (v *SAVisitor) VisitExprStmtNode(node *ExprStmtNode) {
	v.flags |= InExpr
	v.visit(node.Expr)
	v.flags &^= InExpr
}

func (v *SAVisitor) VisitAssignStmtNode(node *AssignStmtNode) {
	const msgAssignMismatch1 = "assignment mismatch: %d variable(s) but %d value(s)"
	const msgAssignMismatch2 = "cannot use %s as %s value in assignment"
	const msgAssignMismatch3 = "cannot assign (not addressable)"
	if node.Tok.Type == TtOpDefine {
		v.flags |= InIdDecl
	}
	for _, expr := range node.Lhs {
		v.visit(expr)
	}
	v.flags &^= InIdDecl
	for _, expr := range node.Rhs {
		v.flags |= InExpr
		v.visit(expr)
		v.flags &^= InExpr
	}

	if node.Tok.Type == TtOpDefine {
		// Short-var declaration
		DeclareShortHand(v, v.currentScope, node)
	} else {
		// Regular assignment
		if len(node.Lhs) != len(node.Rhs) {
			v.logError(node.Tok.Pos, msgAssignMismatch1, len(node.Lhs), len(node.Rhs))
		} else {
			for i := 0; i < len(node.Lhs); i++ {
				if !IsAssignable(v.currentScope, node.Lhs[i]) {
					v.logError(node.Tok.Pos, msgAssignMismatch3)
				}
				if !TypesAreComparable(node.Lhs[i].GetType(), node.Rhs[i].GetType()) {
					v.logError(node.Tok.Pos, msgAssignMismatch2, node.Rhs[i].GetType().String(), node.Lhs[i].GetType().String())
				}
			}
		}
	}
}

func (v *SAVisitor) VisitReturnStmtNode(node *ReturnStmtNode) {
	var types []Type

	if node.Result != nil {
		v.flags |= InExpr
		v.visit(node.Result)
		v.flags &^= InExpr
		types = append(types, node.Result.GetType())

	}
	// NOTE: Get Signature of Function, ensure matches.
	// Loop through, ensure each comparable # match.
}

func (v *SAVisitor) VisitBranchStmtNode(node *BranchStmtNode) {
	const msgNotEnclosed = "%s not in for statement"
	if NearestEnclosingScope(v.currentScope, "for") == nil {
		v.logError(node.Tok.Pos, msgNotEnclosed, node.Tok.Lexeme)
	}
}

func (v *SAVisitor) VisitBlockStmtNode(node *BlockStmtNode) {
	newScopeCreated := false
	if v.currentScope.isFunc && v.enterFuncBlock {
		// For function body, the scope has already been created in FuncDeclNode.
		// So, do not create a new block scope here.
		v.enterFuncBlock = false
	} else {
		v.currentScope = NewScope(v.currentScope, "block")
		newScopeCreated = true
	}
	for _, stmt := range node.List {
		v.visit(stmt)
	}
	if newScopeCreated {
		v.currentScope = v.currentScope.Parent()
	}
}

func (v *SAVisitor) VisitIfStmtNode(node *IfStmtNode) {
	// In Go and Grobbit, If-statements are in their own scope (in addition to the scope of the body-block).
	v.currentScope = NewScope(v.currentScope, "if")
	if !IsNil(node.Init) {
		v.visit(node.Init)
	}
	v.flags |= InExpr
	v.visit(node.Cond)
	v.flags &^= InExpr
	v.visit(node.Body)
	if !IsNil(node.Else) {
		v.visit(node.Else)
	}
	v.currentScope = v.currentScope.Parent()
}

func (v *SAVisitor) VisitForStmtNode(node *ForStmtNode) {
	// In Go and Grobbit, For-statements are in their own scope (in addition to the scope of the body-block).
	v.currentScope = NewScope(v.currentScope, "for")
	if !IsNil(node.Init) {
		v.visit(node.Init)
	}
	if !IsNil(node.Cond) {
		v.flags |= InExpr
		v.visit(node.Cond)
		v.flags &^= InExpr
	}
	if !IsNil(node.Post) {
		v.visit(node.Post)
	}
	v.visit(node.Body)
	v.currentScope = v.currentScope.Parent()
}

func (v *SAVisitor) VisitImportSpecNode(node *ImportSpecNode) {
	_ = node
	// v.visit(v, node.Name)  Always nil, Tok lexeme used instead to identify prefix.
	DeclarePkgName(v.currentScope, node)
}

func (v *SAVisitor) VisitValueSpecNode(node *ValueSpecNode) {
	v.flags |= InIdDecl
	for _, id := range node.Names {
		v.visit(id)
	}
	v.flags &^= InIdDecl
	v.flags |= InType
	if !IsNil(node.Type) {
		v.visit(node.Type)
	}
	v.flags &^= InType
	v.flags |= InExpr
	for _, expr := range node.Values {
		v.visit(expr)
	}
	v.flags &^= InExpr

	if v.flags&InFuncDecl != 0 {
		DeclareValueSpec(v, v.currentScope, LocalVar, node, v.flags)
	} else {
		DeclareValueSpec(v, v.packageScope, PackageVar, node, v.flags)
	}

}

func (v *SAVisitor) VisitGenDeclNode(node *GenDeclNode) {
	var flag SAFlags
	if node.Tok.Type == TtKwConst {
		flag = InConstDecl
	} else {
		flag = InVarDecl
	}
	v.flags |= flag
	for _, spec := range node.Specs {
		v.visit(spec)
	}
	v.flags &^= flag
}

func (v *SAVisitor) VisitFuncDeclNode(node *FuncDeclNode) {
	// For functions, create the scope for the body right away, to be able to add the parameters, if any.
	v.flags |= InFuncDecl
	v.currFuncName = node.Name.Tok.Lexeme
	v.currentScope = NewScope(v.currentScope, "function")
	v.currentScope.isFunc = true
	v.flags |= InParamDecl
	for _, field := range node.Type.Params {
		v.flags |= InIdDecl
		for _, id := range field.Names {
			v.visit(id)
		}
		v.flags &^= InIdDecl
		v.flags |= InType
		v.visit(field.Type)
		v.flags &^= InType
		DeclareVars(v, v.currentScope, ParamVar, field.Type.GetType(), field.Names...)
	}
	v.flags &^= InParamDecl
	if node.Type.Type != nil {
		v.flags |= InType
		v.visit(node.Type.Type)
		v.flags &^= InType
	}
	DeclareFunc(v.packageScope, node) // Do this before visiting the function block in case
	// we need to look the function signature up in there.
	v.enterFuncBlock = true // Mark that we are entering the function block, to prevent BlockStmt from generating a new scope.
	v.visit(node.Body)
	v.enterFuncBlock = false
	v.currentScope = v.currentScope.Parent()
	v.currFuncName = ""
	v.flags &^= InFuncDecl
}

func (v *SAVisitor) VisitFileNode(node *FileNode) {
	// Note: v.visit(node.Name)
	for _, spec := range node.Imports {
		v.visit(spec)
	}
	for _, decl := range node.Decls {
		v.visit(decl)
	}
}

func (v *SAVisitor) Root(node Node, handler ErrorHandler) *Scope {
	v.errorHandler = handler
	v.output.Reset()
	v.universalScope = NewScope(nil, "universal")
	PopulateScope(v, v.universalScope)
	v.packageScope = NewScope(v.universalScope, "package")
	v.fileScope = NewScope(v.packageScope, "file")
	v.currentScope = v.fileScope

	v.visit(node)
	return v.packageScope
}
