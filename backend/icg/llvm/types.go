package llvm

import (
	"fmt"
)

type IRType interface {
	_IRType_()
	Underlying() IRType
	String() string
}

type IRBuiltinType uint8

const (
	Void IRBuiltinType = iota
	I1
	I32
	I64
	Float
)

var builtinTypeNames = [...]string{
	Void:  "void",
	I1:    "i1",
	I32:   "i32",
	I64:   "i64",
	Float: "double",
}

func (t IRBuiltinType) _IRType_()          {}
func (t IRBuiltinType) Underlying() IRType { return t }
func (t IRBuiltinType) String() string {
	return builtinTypeNames[t]
}

func NewIRBuiltinType(typ IRBuiltinType) *IRBuiltinType {
	return &typ
}

type IRPointerType struct {
	Deref IRType
}

func (t IRPointerType) _IRType_()          {}
func (t IRPointerType) Underlying() IRType { return t.Deref }
func (t IRPointerType) String() string     { return t.Deref.String() + "*" }

func NewIRPointerType(typ IRType) *IRPointerType {
	return &IRPointerType{Deref: typ}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var counterRegLab uint = 0

func NextRegLabCounter() uint {
	counterRegLab++
	return counterRegLab
}

func ResetRegLabCounter() {
	counterRegLab = 0
}

type Label uint

func (l Label) String() string {
	return fmt.Sprintf("%%%d", uint(l))
}

func NewLabel() Label {
	return Label(NextRegLabCounter())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Operand interface {
	_Operand_()
	GetType() IRType
	GetValue() string
	String() string
}

type Literal struct {
	Typ  IRType
	Text string
}

func (l Literal) _Operand_()       {}
func (l Literal) GetType() IRType  { return l.Typ }
func (l Literal) GetValue() string { return l.Text }
func (l Literal) String() string   { return fmt.Sprintf("%s %s", l.GetType().String(), l.GetValue()) }

func NewLiteral(typ IRType, txt string) *Literal {
	return &Literal{Typ: typ, Text: txt}
}

type Register struct {
	Typ IRType
	Num uint
}

func (r Register) _Operand_()       {}
func (r Register) GetType() IRType  { return r.Typ }
func (r Register) GetValue() string { return fmt.Sprintf("%%%d", r.Num) }
func (r Register) String() string   { return fmt.Sprintf("%s %s", r.GetType().String(), r.GetValue()) }

func NewRegister(typ IRType) *Register {
	return &Register{Typ: typ, Num: NextRegLabCounter()}
}

type GlobalId struct {
	Typ  IRType
	Name string
}

func (g GlobalId) _Operand_()       {}
func (g GlobalId) GetType() IRType  { return g.Typ }
func (g GlobalId) GetValue() string { return fmt.Sprintf("@%s", g.Name) }
func (g GlobalId) String() string   { return fmt.Sprintf("%s %s", g.GetType().String(), g.GetValue()) }

func NewGlobalId(name string, typ IRType) *GlobalId {
	return &GlobalId{Typ: typ, Name: name}
}
