package icg

import (
	"Grobbit/backend/icg/llvm"
	. "Grobbit/backend/semana"
	. "Grobbit/common"
	"fmt"
	"strings"
)

func ToIRType(tstr string) llvm.IRType {
	// Grobbit only works with int, bool, and void types.
	var ty llvm.IRType
	if tstr == "bool" {
		ty = llvm.I1
	} else if tstr == "int" {
		ty = llvm.I32
	} else if tstr == "float64" {
		ty = llvm.Float
	} else {
		ty = llvm.Void
	}
	return ty
}

func ChooseLabel(cond bool, labelTrue, labelFalse llvm.Label) llvm.Label {
	if cond {
		return labelTrue
	} else {
		return labelFalse
	}
}

func paramRegs(v *IRVisitor, params []*Field) []*llvm.Register {
	var registers []*llvm.Register
	for _, field := range params {
		for _, id := range field.Names {
			obj, _ := resolveId(v, id)
			if obj != nil {
				reg := llvm.NewRegister(ToIRType(obj.Type().String()))
				registers = append(registers, reg)
			}
		}
	}
	return registers
}

// ScopeNavigator is a helper for traversing scopes.
type ScopeNavigator struct {
	level         int
	scopeUniverse *Scope
	scope         *Scope
	childNo       []int
}

func (s *ScopeNavigator) openScope() {
	if s.level == len(s.childNo) {
		s.childNo = append(s.childNo, 0)
	}
	s.scope = s.scope.Child(s.childNo[s.level])
	//fmt.Printf(" openScope %d %s\n", s.level, s.scope.Comment())
	s.level++
}

func (s *ScopeNavigator) closeScope() {
	if s.level < len(s.childNo) {
		s.childNo[s.level] = 0
	}
	s.level--
	//fmt.Printf("closeScope %d %s\n", s.level, s.scope.Comment())
	s.scope = s.scope.Parent()
	s.childNo[s.level]++
}

func (s *ScopeNavigator) uncountOpenClose() {
	s.childNo[s.level]--
}

// ActiveBlock is a Helper for keeping track of the active block.
// If none active block exists, a new is created on demand.

type ActiveBlock struct {
	active *llvm.Block
}

func (ab *ActiveBlock) Get() *llvm.Block {
	if ab.active == nil {
		ab.active = llvm.NewBlock(llvm.NewLabel())
	}
	return ab.active
}

func (ab *ActiveBlock) Set(b *llvm.Block) {
	ab.active = b
}

func (ab *ActiveBlock) Append(instr llvm.Instruction) *llvm.Register {
	if ab.active == nil {
		ab.active = llvm.NewBlock(llvm.NewLabel())
	}
	return ab.active.Append(instr)
}

func (ab *ActiveBlock) Close(instr llvm.Instruction) bool {
	if ab.active == nil {
		return false
	}
	ab.active.Close(instr)
	ab.active = nil
	return true
}

type LabelPair struct {
	labelContinue llvm.Label
	labelBreak    llvm.Label
}

type IRInfo struct {
	module          *llvm.Module   // active module
	function        *llvm.Function // active function (nil, none is active)
	activeBlock     ActiveBlock
	exprOperand     llvm.Operand // used by expression nodes to return computed 'value'.
	params          []*llvm.Register
	blockAlloc      *llvm.Block
	blockInitialize *llvm.Block
	blockStmts      *llvm.Block
	blockReturn     *llvm.Block
	returnVar       *llvm.Register
	labelsBranch    []LabelPair
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type IRVisitor struct {
	VisitorWalker
	errorHandler ErrorHandler
	output       strings.Builder
	flags        SAFlags
	sp           ScopeNavigator
	ir           IRInfo
}

func (v *IRVisitor) printf(format string, a ...any) {
	builder := strings.Builder{}
	for i := 0; i < v.Level; i++ {
		builder.WriteString(". ")
	}
	builder.WriteString(fmt.Sprintf(format, a...))
	v.output.WriteString(builder.String())
}

func (v *IRVisitor) logError(pos Position, format string, a ...any) {
	str := fmt.Sprintf(format, a...)
	v.errorHandler(pos, str)
}

func (v *IRVisitor) visit(node Node) {
	err := v.VisitDispatcher(v, node)
	if err != nil {
		v.printf("%s\n", err.Error())
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *IRVisitor) VisitIdentifierNode(node *IdentifierNode) {
	if v.flags&InExpr != 0 {
		if IsVariable(v.sp.scope, node) || IsConst(v.sp.scope, node) {
			dest := resolveDest(v, node)
			instr := llvm.NewLoadInstr(dest)
			v.ir.exprOperand = v.ir.activeBlock.Append(instr)
		}
	}
}

func (v *IRVisitor) VisitLiteralNode(node *LiteralNode) {
	v.ir.exprOperand = llvm.NewLiteral(ToIRType(node.GetType().String()), node.Tok.Lexeme)
}

func (v *IRVisitor) VisitSelectorExprNode(node *SelectorExprNode) {
	// No need to implement, not part of MiniGrobbit.
	v.visit(node.Expr)
	v.visit(node.Sel)
}

func (v *IRVisitor) VisitCallExprNode(node *CallExprNode) {
	v.visit(node.Fun)
	var argsOperands []llvm.Operand
	for _, arg := range node.Args {
		v.visit(arg)
		argsOperands = append(argsOperands, v.ir.exprOperand)
	}
	name := "unknown"
	var expr ExprNode = node.Fun
	if sel, ok := expr.(*SelectorExprNode); ok {
		expr = sel.Expr
	}
	if id, ok := expr.(*IdentifierNode); ok {
		name = id.Tok.Lexeme
	}
	instr := llvm.NewCallInstr(ToIRType(node.GetType().String()), name, argsOperands...)
	v.ir.exprOperand = v.ir.activeBlock.Append(instr)
}

func (v *IRVisitor) VisitUnaryExprNode(node *UnaryExprNode) {
	var opRet llvm.Operand
	if node.Tok.Type == TtOpNot {
		opRet = LogicalNotOp(v, node)
	} else {
		v.visit(node.Expr)
		block := v.ir.activeBlock.Get()
		opExpr := v.ir.exprOperand // set during visiting node.Expr
		if node.Tok.Type == TtOpSub {
			opRet = block.Append(llvm.NewIArithmeticInstr(llvm.Sub, llvm.NewLiteral(opExpr.GetType(), "0"), opExpr))
		} else if node.Tok.Type == TtOpAdd {
			opRet = opExpr // Do nothing.
		} else {
			v.logError(node.Tok.Pos, "Ops, Grobbit does not support unary operator '%s' (ignoring).", node.Tok.Lexeme)
			opRet = opExpr // Do nothing.
		}
	}
	v.ir.exprOperand = opRet // continue to propagate the result
}

func (v *IRVisitor) VisitBinaryExprNode(node *BinaryExprNode) {
	var opRet llvm.Operand
	switch node.Tok.Type {
	case TtOpAdd, TtOpSub, TtOpMul, TtOpDiv, TtOpMod:
		opRet = ArithmeticOp(v, node)
	case TtOpEq, TtOpNe, TtOpLt, TtOpLe, TtOpGe, TtOpGt:
		opRet = ComparisonOp(v, node)
	case TtOpOr:
		opRet = LogicalOrOp(v, node)
	case TtOpAnd:
		opRet = LogicalAndOp(v, node)
	default:
		panic(fmt.Sprintf("unknown operator type: %v", node.Tok.Type))
	}
	v.ir.exprOperand = opRet // continue to propagate the result
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *IRVisitor) VisitDeclStmtNode(node *DeclStmtNode) {
	// No ICG to do here.
	v.visit(node.Decl)
}

func (v *IRVisitor) VisitExprStmtNode(node *ExprStmtNode) {
	// No ICG to do here.
	v.flags |= InExpr
	v.visit(node.Expr)
	v.flags &^= InExpr
}

func (v *IRVisitor) VisitAssignStmtNode(node *AssignStmtNode) {

	// Grobbit limitations.
	for _, expr := range node.Lhs {
		if _, ok := expr.(*IdentifierNode); !ok {
			v.logError(node.Tok.Pos, "Ops, Grobbit does not support right-hand side expressions in assignments.")
			return
		}
	}
	if len(node.Lhs) != len(node.Rhs) {
		v.logError(node.Tok.Pos, "Ops, mismatch in the number of right-hand and left-hand expressions.")
		return
	}

	var rhsExpr []llvm.Operand

	v.flags |= InExpr
	for _, expr := range node.Rhs {
		v.visit(expr)
		rhsExpr = append(rhsExpr, v.ir.exprOperand)
	}
	v.flags &^= InExpr

	if node.Tok.Type == TtOpDefine {
		v.flags |= InIdDecl
		for i, expr := range node.Lhs {
			if id, ok := expr.(*IdentifierNode); ok {
				declareConstOrVar(v, id, rhsExpr[i], v.ir.activeBlock.Get())
			}
		}
		v.flags &^= InIdDecl
	} else {
		for i, expr := range node.Lhs {
			if id, ok := expr.(*IdentifierNode); ok {
				dest := resolveDest(v, id)
				instr := llvm.NewStoreInstr(rhsExpr[i], dest)
				v.ir.activeBlock.Append(instr)
			}
		}
	}

}

func (v *IRVisitor) VisitReturnStmtNode(node *ReturnStmtNode) {
	if node.Result == nil && v.ir.returnVar != nil {
		v.logError(node.Token.Pos, "Ops, mismatch in the number of returned values in return statements.")
	}
	v.flags |= InExpr
	var exprOp llvm.Operand = nil
	if node.Result != nil {
		v.visit(node.Result)
		exprOp = v.ir.exprOperand
	}
	v.flags &^= InExpr

	// The return statement does not generate 'ret' statement, but rather branches to the functions return block,
	// which has the return statement. The value to return, if any, is stored in a (alloca generated) variable.
	if !IsNil(exprOp) && v.ir.returnVar != nil {
		v.ir.activeBlock.Append(llvm.NewStoreInstr(exprOp, v.ir.returnVar))
	}
	// Close active block, with a branch to the functions return block.
	v.ir.activeBlock.Close(llvm.NewBranchInstr(v.ir.blockReturn.GetLabel()))
}

func (v *IRVisitor) VisitBranchStmtNode(node *BranchStmtNode) {
	l := len(v.ir.labelsBranch)
	if l > 0 {
		if node.Tok.Type == TtKwBreak {
			v.ir.activeBlock.Close(llvm.NewBranchInstr(v.ir.labelsBranch[l-1].labelBreak))
		} else if node.Tok.Type == TtKwContinue {
			v.ir.activeBlock.Close(llvm.NewBranchInstr(v.ir.labelsBranch[l-1].labelContinue))
		}
	}
}

func (v *IRVisitor) VisitBlockStmtNode(node *BlockStmtNode) {
	// No ICG to do in here.
	v.sp.openScope()
	for _, stmt := range node.List {
		v.visit(stmt)
	}
	v.sp.closeScope()
}

/*
VisitIfStmtNode generates following blocks:

	   Upon entering:
			current:				<-- ir.activeBlock
				...
				nil (block is open, does not have terminal instruction yet)

	   Upon exiting:
			current:
				...
				<init_stmts>
				br cond
	     	cond:
				op = <cond_expr>
				'br op then else' or 'br then end' if no else part in the if-statement.
	     	then:
				<then_stmt>
				br end
		  [ else:
	            <else_stmt>
				br  end ]
	     	end:					<-- ir.activeBlock
				nil (block is open, does not have terminal instruction yet)
*/
func (v *IRVisitor) VisitIfStmtNode(node *IfStmtNode) {

	v.sp.openScope()

	labelCond, labelThen, labelElse, labelEnd := llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel()

	blockInit := v.ir.activeBlock.Get() // We do not create a new block for init, can append to active block.
	if !IsNil(node.Init) {
		v.visit(node.Init)
	}
	blockInit = v.ir.activeBlock.Get() // the active block could have changed!
	blockInit.Close(llvm.NewBranchInstr(labelCond))

	blockCond := v.ir.function.Append(llvm.NewBlock(labelCond, "cond"))
	v.ir.activeBlock.Set(blockCond) // ensure we set the active block (before visiting Cond)
	v.flags |= InExpr
	v.visit(node.Cond)
	v.flags &^= InExpr
	condOp := v.ir.exprOperand
	blockCond = v.ir.activeBlock.Get() // the active block could have changed!

	blockThen := v.ir.function.Append(llvm.NewBlock(labelThen, "then"))
	v.ir.activeBlock.Set(blockThen) // ensure we set the active block (before visiting Body)
	v.visit(node.Body)
	blockThen = v.ir.activeBlock.Get()
	blockThen.Close(llvm.NewBranchInstr(labelEnd))

	if !IsNil(node.Else) {
		blockElse := v.ir.function.Append(llvm.NewBlock(labelElse, "else"))
		v.ir.activeBlock.Set(blockElse) // ensure we set the active block (before visiting Body)
		v.visit(node.Else)
		blockElse = v.ir.activeBlock.Get()
		blockElse.Close(llvm.NewBranchInstr(labelEnd))
		blockCond.Close(llvm.NewBranchCondInstr(condOp, labelThen, labelElse))
	} else {
		blockCond.Close(llvm.NewBranchCondInstr(condOp, labelThen, labelEnd))
	}

	blockEnd := v.ir.function.Append(llvm.NewBlock(labelEnd, "end")) // We leave this block open and as the active block.
	v.ir.activeBlock.Set(blockEnd)

	v.sp.closeScope()
}

func (v *IRVisitor) VisitForStmtNode(node *ForStmtNode) {
	v.sp.openScope()
	ForStmt(v, node)
	v.sp.closeScope()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *IRVisitor) VisitImportSpecNode(node *ImportSpecNode) {
	v.logError(node.Tok.Pos,
		"Ops, Grobbit does not implement imports (can use Print(int), Print_F(double), PrintN(n, ...), PrintN_F(n, ...).")
	// v.visit(v, node.Name)  Always nil, Tok lexeme used instead to identify prefix.
}

func (v *IRVisitor) VisitValueSpecNode(node *ValueSpecNode) {
	v.flags |= InIdDecl
	for _, id := range node.Names {
		v.visit(id)
	}
	v.flags &^= InIdDecl
	if !IsNil(node.Type) {
		v.flags |= InType
		v.visit(node.Type)
		v.flags &^= InType
	}
	v.flags |= InExpr
	for _, expr := range node.Values {
		v.visit(expr)
		str := expr.GetType().String()
		if str != "int" && str != "bool" && str != "float64" {
			v.logError(node.Token.Pos, "Ops, Grobbit only supports int, bool, and float64 types.")
		}
	}
	v.flags &^= InExpr
	for i := 0; i < len(node.Names); i++ {
		if i < len(node.Values) {
			if literal, ok := node.Values[i].(*LiteralNode); ok {
				declareConstOrVar(v, node.Names[i], llvm.NewLiteral(ToIRType(literal.GetType().String()), literal.Tok.Lexeme), nil)
				continue
			} else {
				v.logError(node.Token.Pos, "Ops, Grobbit only supports initialization with literal values.")
			}
		}
		declareConstOrVar(v, node.Names[i], llvm.NewLiteral(ToIRType(node.Type.GetType().String()), "0"), nil)
	}
}

func (v *IRVisitor) VisitGenDeclNode(node *GenDeclNode) {
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

func (v *IRVisitor) VisitFuncDeclNode(node *FuncDeclNode) {

	llvm.ResetRegLabCounter()

	labelAlloc, labelInitialize, labelStmts, labelReturn := llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel()

	var retType llvm.IRType
	if node.Type.Type == nil {
		retType = llvm.Void
	} else {
		retType = ToIRType(node.Type.Type.GetType().String())
	}

	name := node.Name.Tok.Lexeme
	if name == "main" {
		name = "main_Grobbit_16__"
	}
	v.ir.function = llvm.NewFunction(name, nil, retType)
	v.ir.module.Append(v.ir.function)
	v.ir.activeBlock.Set(nil)
	v.ir.params = nil
	v.ir.exprOperand = nil
	v.ir.returnVar = nil
	v.ir.blockAlloc = llvm.NewBlock(labelAlloc)
	v.ir.blockInitialize = llvm.NewBlock(labelInitialize)
	v.ir.blockStmts = llvm.NewBlock(labelStmts)
	v.ir.blockReturn = llvm.NewBlock(labelReturn)
	v.ir.function.Append(v.ir.blockAlloc)
	v.ir.function.Append(v.ir.blockInitialize)
	v.ir.function.Append(v.ir.blockStmts)

	v.sp.openScope()

	if node.Type.Type != nil {
		// Allocate a space on the stack for the return value and initialize it with the default value (for type).
		t := ToIRType(node.Type.Type.GetType().String())
		v.ir.returnVar = v.ir.blockAlloc.Append(llvm.NewAllocaInstr(t))
		v.ir.blockInitialize.Append(llvm.NewStoreInstr(llvm.NewLiteral(t, "0"), v.ir.returnVar))
	}

	params := paramRegs(v, node.Type.Params)
	v.flags |= InFuncDecl
	v.flags |= InParamDecl
	i := 0
	for _, field := range node.Type.Params {
		v.flags |= InIdDecl
		for _, id := range field.Names {
			v.visit(id)
			declareConstOrVar(v, id, params[i], nil)
			i += 1
		}
		v.flags &^= InIdDecl
		v.flags |= InType
		v.visit(field.Type)
		v.flags &^= InType
	}
	v.flags &^= InParamDecl
	v.ir.function.SetParams(params)

	if node.Type.Type != nil {
		v.flags |= InType
		v.visit(node.Type.Type)
		v.flags &^= InType
	}
	v.sp.closeScope()
	v.sp.uncountOpenClose()

	v.ir.activeBlock.Set(v.ir.blockStmts)
	v.visit(node.Body)
	v.ir.blockStmts = v.ir.activeBlock.Get()
	v.flags &^= InFuncDecl

	v.ir.blockAlloc.Close(llvm.NewBranchInstr(labelInitialize))
	v.ir.blockInitialize.Close(llvm.NewBranchInstr(labelStmts))
	v.ir.blockStmts.Close(llvm.NewBranchInstr(labelReturn))
	var ret llvm.Operand
	if v.ir.returnVar != nil {
		ret = v.ir.blockReturn.Append(llvm.NewLoadInstr(v.ir.returnVar))
	} else {
		ret = llvm.NewLiteral(llvm.Void, "")
	}
	v.ir.blockReturn.Close(llvm.NewReturnInstr(ret))
	v.ir.function.Append(v.ir.blockReturn)
	v.ir.function = nil
}

func (v *IRVisitor) VisitFileNode(node *FileNode) {
	v.ir.module = llvm.NewModule()
	universeDeclarations(v)

	v.sp.openScope()

	v.visit(node.Name)
	for _, spec := range node.Imports {
		v.visit(spec)
	}
	for _, decl := range node.Decls {
		v.visit(decl)
	}
	v.sp.closeScope()
}

func (v *IRVisitor) Root(filenameOut string, node Node, handler ErrorHandler, scope *Scope) {
	v.errorHandler = handler
	v.sp = ScopeNavigator{level: 0, scopeUniverse: scope.Parent(), scope: scope}
	v.output.Reset()
	v.visit(node)
	v.ir.module.WriteToFile(filenameOut + "_norab")
	v.ir.module.Relabel()
	fmt.Println("Writing out file: ", filenameOut)
	v.ir.module.WriteToFile(filenameOut)
}
