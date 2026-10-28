package icg

import (
	"Grobbit/backend/icg/llvm"
	. "Grobbit/backend/semana"
	. "Grobbit/common"
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
	labelCond, labelPost, labelBody, labelEnd := llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel(), llvm.NewLabel()

	if node.Init != nil {
		v.visit(node.Init)
	}
	blockInit := v.ir.activeBlock.Get()
	blockInit.Close(llvm.NewBranchInstr(labelCond))

	blockCond := v.ir.function.Append(llvm.NewBlock(labelCond, "for_cond"))
	v.ir.activeBlock.Set(blockCond)

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

	blockBody := v.ir.function.Append(llvm.NewBlock(labelBody, "for_body"))
	v.ir.activeBlock.Set(blockBody)
	v.ir.labelsBranch = append(v.ir.labelsBranch, LabelPair{
		labelContinue: labelPost,
		labelBreak:    labelEnd,
	})

	v.visit(node.Body)

	v.ir.labelsBranch = v.ir.labelsBranch[:len(v.ir.labelsBranch)-1]

	v.ir.activeBlock.Close(llvm.NewBranchInstr(labelPost))

	blockPost := v.ir.function.Append(llvm.NewBlock(labelPost, "for_post"))
	v.ir.activeBlock.Set(blockPost)
	if node.Post != nil {
		v.visit(node.Post)
	}
	blockPost = v.ir.activeBlock.Get()

	blockPost.Close(llvm.NewBranchInstr(labelCond))

	blockEnd := v.ir.function.Append(
		llvm.NewBlock(labelEnd, "for_end"))
	v.ir.activeBlock.Set(blockEnd)
}
