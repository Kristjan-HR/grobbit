package common

import (
	"fmt"
	"strings"
)

type Type interface {
	String() string
}

type TypeFlags uint16

const (
	_ TypeFlags = 1 << iota
	TypeInt
	TypeFloat
	TypeBool
	TypeString
	TypeRune
	TypePointer
	TypeSignature
	TypeInvalid // Good to have to be able to continue with type-checking even tough a type-error occurs.
)

func TypeFlag(t Type) TypeFlags {
	if t == nil {
		fmt.Println("<nil>, should not happen")
		return 0
	}
	switch t.(type) {
	case *IntType:
		return TypeInt
	case *FloatType:
		return TypeFloat
	case *BoolType:
		return TypeBool
	case *StringType:
		return TypeString
	case *RuneType:
		return TypeRune
	case *PointerType:
		return TypePointer
	case *SignatureType:
		return TypeSignature
	case *InvalidType:
		return TypeInvalid
	default:
		panic("unreachable")
	}
}

func TypeIsOneOf(t Type, flags TypeFlags) bool {
	return TypeFlag(t)&flags != 0
}

func TypeIsNilPointer(t Type) bool {
	return (TypeFlag(t)&TypePointer != 0) && t.(*PointerType).GetBaseType() == nil
}

func TypeIsOrderable(t Type) bool {
	return TypeFlag(t)&(TypeInt|TypeFloat|TypeString|TypeRune) != 0
}

func TypesAreIdentical(t1 Type, t2 Type) bool {
	return t1.String() == t2.String()
}

func TypesAreComparable(t1 Type, t2 Type) bool {
	if TypeIsOneOf(t1, TypePointer) && TypeIsNilPointer(t2) {
		return true
	}
	if TypeIsOneOf(t2, TypePointer) && TypeIsNilPointer(t1) {
		return true
	}
	return TypesAreIdentical(t1, t2)
}

type typeStruct struct{ name string }

func (t *typeStruct) String() string     { return t.name }
func (t *typeStruct) Same(typ Type) bool { return t.String() == typ.String() }

type IntType struct{ typeStruct }

func NewIntType() *IntType { return &IntType{typeStruct{name: "int"}} }

type FloatType struct{ typeStruct }

func NewFloatType() *FloatType { return &FloatType{typeStruct{name: "float64"}} }

type BoolType struct{ typeStruct }

func NewBoolType() *BoolType { return &BoolType{typeStruct{name: "bool"}} }

type StringType struct{ typeStruct }

func NewStringType() *StringType { return &StringType{typeStruct{name: "string"}} }

type RuneType struct{ typeStruct }

func NewRuneType() *RuneType { return &RuneType{typeStruct{name: "rune"}} }

type PointerType struct {
	typeStruct
	baseType Type
}

func (t *PointerType) GetBaseType() Type { return t.baseType }

func NewPointerType(typ Type) *PointerType {
	var name string
	if IsNil(typ) {
		name = "nil"
	} else {
		name = "*" + typ.String()
	}
	return &PointerType{typeStruct{name: name}, typ}
}

type SignatureType struct {
	typeStruct
	Params  []Type
	Results []Type
}

func TypesAsString(types []Type) string {
	var builder strings.Builder
	if len(types) > 0 {
		for i, t := range types {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(t.String())
		}
	}
	return builder.String()
}

func NewSignatureType(params []Type, results []Type) *SignatureType {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(TypesAsString(params))
	builder.WriteString(")")
	if len(results) > 0 {
		builder.WriteString(" ")
		if len(results) == 1 {
			builder.WriteString(results[0].String())
		} else {
			builder.WriteString("(")
			builder.WriteString(TypesAsString(results))
			builder.WriteString(")")
		}
	}
	return &SignatureType{typeStruct{name: builder.String()}, params, results}
}

type InvalidType struct {
	typeStruct
}

func NewInvalidType() *InvalidType { return &InvalidType{typeStruct{name: "invalid type"}} }
