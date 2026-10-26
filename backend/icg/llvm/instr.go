package llvm

import (
	"fmt"
	"reflect"
	"strings"
)

type Opcode uint8

const (
	Alloca Opcode = iota
	Load
	Store
	Branch
	Ret
	Call
	Select
	Phi
	IAdd
	ISub
	IMul
	IDiv
	IRem
	ICmp
	FAdd
	FSub
	FMul
	FDiv
	FCmp
)

var opcodeNames = [...]string{
	Alloca: "alloca",
	Load:   "load",
	Store:  "store",
	Branch: "br",
	Ret:    "ret",
	Call:   "call",
	Select: "select",
	Phi:    "phi",
	IAdd:   "add",
	ISub:   "sub",
	IMul:   "mul",
	IDiv:   "sdiv",
	IRem:   "srem",
	ICmp:   "icmp",
	FAdd:   "fadd",
	FSub:   "fsub",
	FMul:   "fmul",
	FDiv:   "fdiv",
	FCmp:   "fcmp",
}

func (o Opcode) String() string {
	return opcodeNames[o]
}

////////////////////////////////////////////////////////////////////////////////

type ICmpOperator uint16

const (
	Eq ICmpOperator = iota
	Ne
	Sgt
	Sge
	Slt
	Sle
)

var iCmpOperatorNames = [...]string{
	Eq:  "eq",
	Ne:  "ne",
	Sgt: "sgt",
	Sge: "sge",
	Slt: "slt",
	Sle: "sle",
}

func (o ICmpOperator) String() string {
	return iCmpOperatorNames[o]
}

type IArithmeticOperator uint16

const (
	Add IArithmeticOperator = iota
	Sub
	Mul
	Div
	Mod
	// No corresponding logical (||, &&, !) operators, implemented in program logic (e.g., short-circuit!)
)

var iArithmeticOperatorOpcode = [...]Opcode{
	Add: IAdd,
	Sub: ISub,
	Mul: IMul,
	Div: IDiv,
	Mod: IRem,
}

////////////////////////////////////////////////////////////////////////////////

type FCmpOperator uint16

const (
	FCmpEq FCmpOperator = iota
	FCmpNe
	FCmpGt
	FCmpGe
	FCmpLt
	FCmpLe
)

var FCmpOperatorNames = [...]string{
	FCmpEq: "oeq",
	FCmpNe: "une",
	FCmpGt: "ogt",
	FCmpGe: "sge",
	FCmpLt: "olt",
	FCmpLe: "ole",
}

func (o FCmpOperator) String() string {
	return FCmpOperatorNames[o]
}

type FArithmeticOperator uint16

const (
	FOpAdd FArithmeticOperator = iota
	FOpSub
	FOpMul
	FOpDiv
	// No corresponding mod or logical (||, &&, !) operators.
)

var fArithmeticOperatorOpcode = [...]Opcode{
	FOpAdd: FAdd,
	FOpSub: FSub,
	FOpMul: FMul,
	FOpDiv: FDiv,
}

////////////////////////////////////////////////////////////////////////////////

type (
	Instruction interface {
		String() string
		IsTerminal() bool
		Register() *Register
		GetO1() Operand
		GetO2() Operand
	}
	instr_ struct {
		op  Opcode
		o1  Operand
		o2  Operand
		ret *Register
	}
)

//func (ins *instr_) String() string { return opcodeNames[ins.op] }

func (ins *instr_) Register() *Register {
	return ins.ret
}
func (ins *instr_) GetOp() Opcode    { return ins.op }
func (ins *instr_) GetO1() Operand   { return ins.o1 }
func (ins *instr_) GetO2() Operand   { return ins.o2 }
func (ins *instr_) IsTerminal() bool { return ins.op == Branch || ins.op == Ret }

type (
	AllocaInstr struct {
		instr_
	}

	LoadInstr struct {
		instr_
	}

	StoreInstr struct {
		instr_
	}

	BranchInstr struct {
		instr_
		Lab Label
	}

	BranchCondInstr struct {
		instr_
		Cond     Operand
		TrueLab  Label
		FalseLab Label
	}

	ReturnInstr struct {
		instr_
	}

	ICmpInstr struct {
		instr_
		Operator ICmpOperator
	}

	FCmpInstr struct {
		instr_
		Operator FCmpOperator
	}

	IArithmeticInstr struct {
		instr_
		Operator IArithmeticOperator
	}

	FArithmeticInstr struct {
		instr_
		Operator FArithmeticOperator
	}

	SelectInstr struct {
		instr_
		Cond Operand
	}

	CallInstr struct {
		instr_
		Typ    IRType
		Prefix string
		Name   string
		Args   []Operand
	}

	ValueLabel struct {
		Op    Operand
		Label Label
	}

	PhiInstr struct {
		instr_
		options []ValueLabel
	}
)

// Check if an interface is nil
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Ptr && v.IsNil()
}

func ensureNonNil(operands ...Operand) {
	for _, o := range operands {
		if isNil(o) {
			panic("Error: Operand to instruction is nil!")
		}
	}
}

func (instr *AllocaInstr) String() string {
	return fmt.Sprintf("%s = %s %s", instr.Register().GetValue(), instr.GetOp().String(),
		instr.Register().GetType().Underlying().String())
}

func NewAllocaInstr(typ IRType) *AllocaInstr {
	return &AllocaInstr{instr_: instr_{op: Alloca, ret: NewRegister(NewIRPointerType(typ))}}
}

func (instr *LoadInstr) String() string {
	return fmt.Sprintf("%s = %s %s, %s", instr.Register().GetValue(), instr.GetOp().String(),
		instr.Register().GetType().String(), instr.GetO1().String())
}

func NewLoadInstr(src Operand) *LoadInstr {
	ensureNonNil(src)
	o1 := src
	return &LoadInstr{instr_: instr_{op: Load, ret: NewRegister(o1.GetType().Underlying()), o1: o1, o2: nil}}
}

func (instr *StoreInstr) String() string {
	return fmt.Sprintf("%s %s, %s", instr.GetOp().String(), instr.GetO1().String(), instr.GetO2().String())
}

func NewStoreInstr(src Operand, dest Operand) *StoreInstr {
	ensureNonNil(src, dest)
	o1, o2 := src, dest
	return &StoreInstr{instr_: instr_{op: Store, ret: nil, o1: o1, o2: o2}}
}

func (instr *BranchInstr) String() string {
	return fmt.Sprintf("%s label %s", instr.GetOp().String(), instr.Lab.String())
}

func NewBranchInstr(lab Label) *BranchInstr {
	return &BranchInstr{instr_: instr_{op: Branch}, Lab: lab}
}

func (instr *BranchCondInstr) String() string {
	return fmt.Sprintf("%s %s, label %s, label %s", instr.GetOp().String(),
		instr.Cond.String(), instr.TrueLab.String(), instr.FalseLab.String())
}

func NewBranchCondInstr(cond Operand, labTrue Label, labFalse Label) *BranchCondInstr {
	ensureNonNil(cond)
	return &BranchCondInstr{instr_: instr_{op: Branch}, Cond: cond, TrueLab: labTrue, FalseLab: labFalse}
}

func (instr *ReturnInstr) String() string {
	var str string
	if instr.GetO1() == nil {
		str = "void"
	} else {
		str = instr.GetO1().String()
	}
	return fmt.Sprintf("%s %s", instr.GetOp().String(), str)
}

func NewReturnInstr(src Operand) *ReturnInstr {
	ensureNonNil(src)
	o1 := src
	return &ReturnInstr{instr_: instr_{op: Ret, ret: nil, o1: o1, o2: nil}}
}

func (instr *ICmpInstr) String() string {
	return fmt.Sprintf("%s = %s %s %s, %s", instr.Register().GetValue(), instr.GetOp().String(),
		instr.Operator.String(), instr.GetO1().String(), instr.GetO2().GetValue())
}

func NewICmpInstr(op ICmpOperator, lhs, rhs Operand) *ICmpInstr {
	ensureNonNil(lhs, rhs)
	o1, o2 := lhs, rhs
	return &ICmpInstr{instr_: instr_{op: ICmp, ret: NewRegister(I1), o1: o1, o2: o2}, Operator: op}
}

func (instr *IArithmeticInstr) String() string {
	return fmt.Sprintf("%s = %s %s, %s", instr.Register().GetValue(), instr.GetOp().String(),
		instr.GetO1().String(), instr.GetO2().GetValue())
}

func NewIArithmeticInstr(op IArithmeticOperator, lhs, rhs Operand) *IArithmeticInstr {
	ensureNonNil(lhs, rhs)
	o1, o2 := lhs, rhs
	typ := o1.GetType()
	if typ != o2.GetType() {
		fmt.Println("Ops, types of both operands in NewIArithmeticInstr must match.")
	}
	return &IArithmeticInstr{instr_: instr_{op: iArithmeticOperatorOpcode[op],
		ret: NewRegister(typ), o1: o1, o2: o2}, Operator: op}
}

func (instr *FCmpInstr) String() string {
	return fmt.Sprintf("%s = %s %s %s, %s", instr.Register().GetValue(), instr.GetOp().String(),
		instr.Operator.String(), instr.GetO1().String(), instr.GetO2().GetValue())
}

func NewFCmpInstr(op FCmpOperator, lhs, rhs Operand) *FCmpInstr {
	ensureNonNil(lhs, rhs)
	o1, o2 := lhs, rhs
	return &FCmpInstr{instr_: instr_{op: FCmp, ret: NewRegister(I1), o1: o1, o2: o2}, Operator: op}
}

func (instr *FArithmeticInstr) String() string {
	return fmt.Sprintf("%s = %s %s, %s", instr.Register().GetValue(), instr.GetOp().String(),
		instr.GetO1().String(), instr.GetO2().GetValue())
}

func NewFArithmeticInstr(op FArithmeticOperator, lhs, rhs Operand) *FArithmeticInstr {
	ensureNonNil(lhs, rhs)
	o1, o2 := lhs, rhs
	typ := o1.GetType()
	if typ != o2.GetType() {
		fmt.Println("Ops, types of both operands in NewFArithmeticInstr must match.")
	}
	return &FArithmeticInstr{instr_: instr_{op: fArithmeticOperatorOpcode[op],
		ret: NewRegister(typ), o1: o1, o2: o2}, Operator: op}
}

func (instr *SelectInstr) String() string {
	reg := instr.Register()
	str := fmt.Sprintf("%s = %s %s %s, %s, %s",
		reg.GetValue(),
		instr.GetOp().String(), I1.String(), instr.Cond.GetValue(),
		instr.GetO1().String(), instr.GetO2().String())
	return str
}

func NewSelectInstr(cond, tr, fa Operand) *SelectInstr {
	ensureNonNil(cond, tr, fa)
	o1, o2 := tr, fa
	return &SelectInstr{instr_: instr_{op: Select, ret: NewRegister(o1.GetType()), o1: o1, o2: o2}, Cond: cond}
}

func (instr *CallInstr) String() string {
	argsStr := ""
	for i, arg := range instr.Args {
		if i > 0 {
			argsStr += ", "
		}
		argsStr += arg.GetType().String() + " " + arg.GetValue()
	}
	prefixStr := ""
	if instr.Typ.String() != "void" {
		prefixStr = fmt.Sprintf("%s = ", instr.Register().GetValue())
	}
	return fmt.Sprintf("%s%s %s %s @%s(%s)", prefixStr, instr.GetOp().String(),
		instr.Typ.String(), instr.Prefix, instr.Name, argsStr)
}

func NewCallInstr(typ IRType, name string, args ...Operand) *CallInstr {
	// A quick-and-dirty hack to allow the variadic number of args by @PrintN.
	prefix := ""
	if name == "PrintN" || name == "PrintN_F" {
		prefix = "(i32, ...)"
	}
	return &CallInstr{instr_: instr_{op: Call, ret: NewRegister(typ)}, Typ: typ, Prefix: prefix, Name: name, Args: args}
}

func (instr *PhiInstr) String() string {
	var sb strings.Builder
	str := fmt.Sprintf("%s = %s %s ",
		instr.Register().GetValue(), instr.GetOp().String(), instr.Register().GetType())
	sb.WriteString(str)
	for i, option := range instr.options {
		if i > 0 {
			sb.WriteString(", ")
		}
		optionStr := fmt.Sprintf("[ %s, %s ]", option.Op.GetValue(), option.Label.String())
		sb.WriteString(optionStr)
	}
	return sb.String()
}

func NewPhiInstr(typ IRType, options ...ValueLabel) *PhiInstr {
	return &PhiInstr{instr_: instr_{op: Phi, ret: NewRegister(typ)}, options: options}
}
