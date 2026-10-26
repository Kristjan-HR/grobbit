package icg

import (
	"Grobbit/backend/icg/llvm"
	. "Grobbit/backend/semana"
	. "Grobbit/common"
)

var mapIdToReg = make(map[Object]*llvm.Register)
var mapIdToGlobalId = make(map[Object]*llvm.GlobalId)

func isGlobalScope(sp *ScopeNavigator, s *Scope) bool {
	if s == sp.scopeUniverse {
		return true // Universe scope
	}
	for i := 0; i < sp.scopeUniverse.NumChildren(); i++ {
		if s == sp.scopeUniverse.Child(i) {
			return true // Package scope
		}
	}
	return false
}

func valueOrDefault(operand llvm.Operand) string {
	if !IsNil(operand) {
		return operand.GetValue()
	} else {
		return "0"
	}
}

func doMapIdToGlobalId(v *IRVisitor, obj Object, id *IdentifierNode) *llvm.GlobalId {
	gid := mapIdToGlobalId[obj]
	if gid == nil {
		v.logError(id.Tok.Pos, "Ops, global id '%s' should be in map.", id.Tok.Lexeme)
	}
	return gid
}

func doMapIdToReg(v *IRVisitor, obj Object, id *IdentifierNode) *llvm.Register {
	reg := mapIdToReg[obj]
	if reg == nil {
		v.logError(id.Tok.Pos, "Ops, register '%s' should be in map.", id.Tok.Lexeme)
	}
	return reg
}

func resolveId(v *IRVisitor, id *IdentifierNode) (Object, *Scope) {
	obj, scope := Resolve(v.sp.scope, id, false)
	if obj == nil {
		v.logError(id.Tok.Pos, "Ops, identifier '%s' should be in symbol table.", id.Tok.Lexeme)
	}
	return obj, scope
}

func resolveDest(v *IRVisitor, id *IdentifierNode) llvm.Operand {
	obj, scope := resolveId(v, id)
	if obj == nil {
		return nil
	}
	var dest llvm.Operand
	if isGlobalScope(&v.sp, scope) {
		dest = doMapIdToGlobalId(v, obj, id)
	} else {
		dest = doMapIdToReg(v, obj, id)
	}
	return dest
}

func universeDeclarations(v *IRVisitor) {

	v.ir.module.AppendHeader("; Available external functions.")
	v.ir.module.AppendHeader("declare void @Print(i32)")
	v.ir.module.AppendHeader("declare void @PrintN(i32, ...)")
	v.ir.module.AppendHeader("declare void @Print_F(double)")
	v.ir.module.AppendHeader("declare void @PrintN_F(i32, ...)\n\n")

	/*
		pos := Position{Line: 0, Col: 0}
		token := Token{Type: TtIdentifier, Lexeme: "true", Pos: pos}
		id := &IdentifierNode{Tok: token}
		declareConstOrVar(v, id, llvm.NewLiteral(llvm.I1, "1"), nil)
		token = Token{TtIdentifier, "false", pos, pos}
		id = &IdentifierNode{Tok: token}
		declareConstOrVar(v, id, llvm.NewLiteral(llvm.I1, "0"), nil)
		//v.ir.module.Append(llvm.NewGlobal("true", llvm.I1, "1", true))
		//v.ir.module.Append(llvm.NewGlobal("false", llvm.I1, "0", true))
	*/
}

func declareConstOrVar(v *IRVisitor, id *IdentifierNode, init llvm.Operand, blockInitialize *llvm.Block) {
	obj, scope := resolveId(v, id)
	if obj == nil {
		return
	}
	ofType := ToIRType(obj.Type().String())
	if init == nil {
		init = llvm.NewLiteral(llvm.I32, "0") // NOTE
	}
	initValue := valueOrDefault(init)

	if isGlobalScope(&v.sp, scope) {
		name := obj.Name()
		_, isConst := obj.(*Const)
		mapIdToGlobalId[obj] = llvm.NewGlobalId(name, llvm.NewIRPointerType(ofType)) // NOTE: Take as init operand?
		v.ir.module.Append(llvm.NewGlobal(name, ofType, initValue, isConst))
	} else {
		reg := v.ir.blockAlloc.Append(llvm.NewAllocaInstr(ofType))
		mapIdToReg[obj] = reg
		if blockInitialize == nil {
			blockInitialize = v.ir.blockInitialize
		}
		blockInitialize.Append(llvm.NewStoreInstr(init /*llvm.NewLiteral(ofType, initValue)*/, reg))
	}
}
