package icg

import (
	"Grobbit/backend/icg/llvm"
	. "Grobbit/common"
	. "Grobbit/backend/semana"
)

var mapIArthOp = []llvm.IArithmeticOperator{
	TtOpAdd: llvm.Add,
	TtOpSub: llvm.Sub,
	TtOpMul: llvm.Mul,
	TtOpDiv: llvm.Div,
	TtOpMod: llvm.Mod,
}

var mapICmpOp = []llvm.ICmpOperator{
	TtOpEq: llvm.Eq,
	TtOpNe: llvm.Ne,
	TtOpGt: llvm.Sgt,
	TtOpGe: llvm.Sge,
	TtOpLt: llvm.Slt,
	TtOpLe: llvm.Sle,
}

var mapFArthOp = []llvm.FArithmeticOperator{
	TtOpAdd: llvm.FOpAdd,
	TtOpSub: llvm.FOpSub,
	TtOpMul: llvm.FOpMul,
	TtOpDiv: llvm.FOpDiv,
}

var mapFCmpOp = []llvm.FCmpOperator{
	TtOpEq: llvm.FCmpEq,
	TtOpNe: llvm.FCmpNe,
	TtOpGt: llvm.FCmpGt,
	TtOpGe: llvm.FCmpGe,
	TtOpLt: llvm.FCmpLt,
	TtOpLe: llvm.FCmpLe,
}

/*
ArithmeticOp implements a code for the arithmetic operators ( + - * / % ).
Returns a llvm value (llvm.Operand) with the result.
*/
func ArithmeticOp(v *IRVisitor, node *BinaryExprNode) llvm.Operand {
	v.visit(node.Lhs)
	lhsOp := v.ir.exprOperand // Sub-expression returns register where result stored via v.ir.exprOperand
	v.visit(node.Rhs)
	rhsOp := v.ir.exprOperand
	var operand *llvm.Register
	if node.GetType().String() == "float64" {
		operand = v.ir.activeBlock.Append(llvm.NewFArithmeticInstr(mapFArthOp[node.Tok.Type], lhsOp, rhsOp))
	} else {
		operand = v.ir.activeBlock.Append(llvm.NewIArithmeticInstr(mapIArthOp[node.Tok.Type], lhsOp, rhsOp))
	}
	return operand
}

/*
ComparisonOp implements a code for the comparison operators (==, !=, <= <, > >=).
Returns a llvm value (llvm.Operand) with the result.
*/
func ComparisonOp(v *IRVisitor, node *BinaryExprNode) llvm.Operand {
	// TO DO: add your do here (replacing dummy placeholder below).

	// dummy := v.ir.activeBlock.Append(llvm.NewICmpInstr(llvm.Eq, llvm.NewLiteral(llvm.I32, "0"), llvm.NewLiteral(llvm.I32, "0")))
	// return dummy

	
	v.visit(node.Lhs)
	lhsOp := v.ir.exprOperand
	v.visit(node.Rhs)
	rhsOp := v.ir.exprOperand
	block := v.ir.activeBlock.Get()
	var operand *llvm.Register
	if node.Lhs.GetType().String() == "float64" {
		operand = block.Append(llvm.NewFCmpInstr(mapFCmpOp[node.Tok.Type], lhsOp, rhsOp))
	} else {
		operand = block.Append(llvm.NewICmpInstr(mapICmpOp[node.Tok.Type], lhsOp, rhsOp))
	}
	return operand
}

/*
LogicalNotOp implements a code for the logical not operator.
Returns a llvm value (llvm.Operand) with the result.
*/
func LogicalNotOp(v *IRVisitor, node *UnaryExprNode) llvm.Operand {
	// TO DO: add your do here (replacing dummy placeholder below).
	// TODO: use select instruction

	// dummy := v.ir.activeBlock.Append(llvm.NewICmpInstr(llvm.Eq, llvm.NewLiteral(llvm.I32, "0"), llvm.NewLiteral(llvm.I32, "0")))
	// return dummy

	
	v.visit(node.Expr)
	block := v.ir.activeBlock.Get()
	ret := block.Append(llvm.NewSelectInstr(v.ir.exprOperand, llvm.NewLiteral(llvm.I1, "0"), llvm.NewLiteral(llvm.I1, "1")))
	return ret
}

/*
LogicalOrOp implements a short-circuit code for the logical or operator.
Returns a llvm value (llvm.Operand) with the result.
*/
func LogicalOrOp(v *IRVisitor, node *BinaryExprNode) llvm.Operand {
	// TO DO: add your do here (replacing dummy placeholder below).
	// TODO:
	// block_lhs:
	//  %lhs = eval lhs
	// 	br i1 %lhs, %block_end, %block_rhs
	// block_rhs:
	//  %rhs = eval rhs
	//  br %block_end
	// block_end
	//  %res = phi i1 [%lhs, %block_lhs], [%rhs, %block_rhs]

	labelEnd, labelRhs := llvm.NewLabel(), llvm.NewLabel()
	v.visit(node.Lhs)

	blockLhs := v.ir.activeBlock.Get()
	lhsOp := v.ir.exprOperand
	blockLhs.Close(llvm.NewBranchCondInstr(lhsOp, labelEnd, labelRhs))
	labelLhs := blockLhs.GetLabel()

	blockRhs := v.ir.function.Append(llvm.NewBlock(labelRhs, "rhs"))
	v.ir.activeBlock.Set(blockRhs)
	v.visit(node.Rhs)
	blockRhs = v.ir.activeBlock.Get()
	rhsOp := v.ir.exprOperand
	blockRhs.Close(llvm.NewBranchInstr(labelEnd))
	labelRhsFinal := blockRhs.GetLabel()
	blockEnd := v.ir.function.Append(llvm.NewBlock(labelEnd, "end"))
	ret := blockEnd.Append(llvm.NewPhiInstr(
		lhsOp.GetType(),
		llvm.ValueLabel{Op: lhsOp, Label: labelLhs},
		llvm.ValueLabel{Op: rhsOp, Label: labelRhsFinal},
	))
	v.ir.activeBlock.Set(blockEnd)
	return ret
}

/*
LogicalAndOp implements a short-circuit code for the logical and operator.
Returns a llvm value (llvm.Operand) with the result.
*/
func LogicalAndOp(v *IRVisitor, node *BinaryExprNode) llvm.Operand {
	// TO DO: add your do here (replacing dummy placeholder below).
	// TODO:
	// block_lhs:
	//  %lhs = eval lhs
	//  br i1 %lhs, %block_rhs, %block_end
	// block_rhs:
	//  %rhs = eval rhs
	//  br %block_end
	// block_end:
	//  %res = phi i1 [false, %block_lhs], [%rhs, %block_rhs]

	// dummy := v.ir.activeBlock.Append(llvm.NewICmpInstr(llvm.Eq, llvm.NewLiteral(llvm.I32, "0"), llvm.NewLiteral(llvm.I32, "0")))
	// return dummy

	labelEnd, labelRhs := llvm.NewLabel(), llvm.NewLabel()

	v.visit(node.Lhs)
	blockLhs := v.ir.activeBlock.Get()
	lhsOp := v.ir.exprOperand
	blockLhs.Close(llvm.NewBranchCondInstr(lhsOp, labelRhs, labelEnd))
	labelLhs := blockLhs.GetLabel()

	blockRhs := v.ir.function.Append(llvm.NewBlock(labelRhs, "rhs"))
	v.ir.activeBlock.Set(blockRhs)
	v.visit(node.Rhs)
	blockRhs = v.ir.activeBlock.Get()
	rhsOp := v.ir.exprOperand
	blockRhs.Close(llvm.NewBranchInstr(labelEnd))
	labelRhsFinal := blockRhs.GetLabel()
	blockEnd := v.ir.function.Append(llvm.NewBlock(labelEnd, "end"))
	ret := blockEnd.Append(llvm.NewPhiInstr(
		lhsOp.GetType(),
		llvm.ValueLabel{Op: llvm.NewLiteral(llvm.I1, "0"), Label: labelLhs},
		llvm.ValueLabel{Op: rhsOp, Label: labelRhsFinal},
	))
	v.ir.activeBlock.Set(blockEnd)
	return ret
}

/*
	 ForStmt:
	 VisitForStmtNode generates following blocks:

		   Upon entering:
				current:				<-- ir.activeBlock
					...
					nil (block is open, does not have terminal instruction yet)

		   Upon exiting:
				current:
					...
					<init_stmts>
					'br cond' (if cond) else 'br body'
			   [ cond:
					op = <cond_expr>
					br op body end ]
		      [	post:
					...
					'br cond' (if cond) else 'br body' ]
		     	body:
					<body_stmt>
					'br post' (if post) else 'br cond' (if cond) or 'br body'
		     	end:					<-- ir.activeBlock
					nil (block is open, does not have terminal instruction yet)
*/
func ForStmt(v *IRVisitor, node *ForStmtNode) {
	// TO DO: add your do here (replacing dummy placeholder below).

	// TODO: most difficult
	// for [cond] {}
	// for [init] ; [cond] ; [post] {}
	// blocks exist, even if empty
	// for loops have scope for their init, and the block has a scope

	// TODO:
	// for_init:
	//  ...
	//  br label %for_cond
	// for_cond:
	//  ...
	//  br i1 cond, %for_body, %end
	// for_post:
	//  ...
	//  br label %for_cond
	// for_body
	//  ...
	//  br label %for_post
	// end:
	//  ...
	labelCond, labelPost, labelBody, labelEnd := llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel()
	// for_init:
	if node.Init != nil {
		v.visit(node.Init)
	} // now we jump to condition block
	blockInit := v.ir.activeBlock.Get()
	blockInit.Close(llvm.NewBranchInstr(labelCond))
	//we create condition block
	blockCond := v.ir.function.Append(llvm.NewBlock(labelCond, "for_cond"))
	v.ir.activeBlock.Set(blockCond)
	//evaluate condition if it exists
	if node.Cond != nil {
		v.flags |= InExpr
		v.visit(node.Cond)
		v.flags &^= InExpr
		condOp := v.ir.exprOperand
		blockCond = v.ir.activeBlock.Get()
		blockCond.Close(llvm.NewBranchCondInstr(condOp, labelBody, labelEnd))
	} else {
		blockCond.Close(llvm.NewBranchInstr(labelBody))
	}
	// create the body 
	blockBody := v.ir.function.Append(llvm.NewBlock(labelBody, "for_body"))
	v.ir.activeBlock.Set(blockBody)
	v.ir.labelsBranch = append(v.ir.labelsBranch, LabelPair{
		//with break --> jump to end, with continue --> jump to post
    labelContinue: labelPost,
    labelBreak:    labelEnd,})
	//exec instructions of body
	v.visit(node.Body)
	//each break affects only its loop
	v.ir.labelsBranch = v.ir.labelsBranch[:len(v.ir.labelsBranch)-1]
	//avoid closing active block again if it's laready closed
	v.ir.activeBlock.Close(llvm.NewBranchInstr(labelPost))
	//create post block
	blockPost := v.ir.function.Append(llvm.NewBlock(labelPost, "for_post"))
	v.ir.activeBlock.Set(blockPost)
	if node.Post != nil {
		v.visit(node.Post)
	}
	blockPost = v.ir.activeBlock.Get()
	//check the condition again
	blockPost.Close(llvm.NewBranchInstr(labelCond))
	//block that makes the program continue when it reaches the end or when it sees continue statement
	blockEnd := v.ir.function.Append(
    llvm.NewBlock(labelEnd, "for_end"),)
	v.ir.activeBlock.Set(blockEnd)
}
