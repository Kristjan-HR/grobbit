package semana

import (
	"Grobbit/common"
	"strings"
)

type VarKind uint8

const (
	_          VarKind = iota // (not meaningful)
	PackageVar                // a package-level variable
	LocalVar                  // a local variable
	ParamVar                  // a function parameter variable
)

type Object interface {
	Parent() *Scope
	Name() string
	Pos() common.Position
	Type() common.Type
	// Exported() bool
	//	Id() string
	String() string
	setType(common.Type)
	setParent(*Scope)
}

// An object implements the common parts of an Object.
type object struct {
	parent *Scope
	name   string
	pos    common.Position
	typ    common.Type
}

func (obj *object) Parent() *Scope       { return obj.parent }
func (obj *object) Name() string         { return obj.name }
func (obj *object) Pos() common.Position { return obj.pos }
func (obj *object) Type() common.Type    { return obj.typ }

// func (obj *object) Id() string              { return Id(obj.pkg, obj.name) }
func (obj *object) String() string          { panic("abstract") }
func (obj *object) setParent(parent *Scope) { obj.parent = parent }
func (obj *object) setType(typ common.Type) { obj.typ = typ }

type PkgName struct {
	object
	path string
}

func (p *PkgName) Path() string { return p.path }
func (p *PkgName) String() string {
	var str string
	if p.name != p.path {
		str = p.Name() + " (" + p.Path() + ")"
	} else {
		str = p.Name()
	}
	return "package " + str
}

func NewPkgName(parent *Scope, impNode *common.ImportSpecNode, typ common.Type) *PkgName {
	var name string
	if impNode.Tok.Type == common.TtIdentifier {
		name = impNode.Tok.Lexeme
	} else {
		name = strings.Trim(impNode.Path, "\"")
	}
	// NOTE: .  imports all methods (cannot do here) Disallow in Grobbit?
	return &PkgName{object{parent, name, impNode.Tok.Pos, typ}, impNode.Path}
}

type Const struct {
	object
}

func (v *Const) String() string { return "const " + v.Name() + " " + v.typ.String() }

func NewConst(parent *Scope, idNode *common.IdentifierNode, typ common.Type) *Const {
	return &Const{object{parent, idNode.Tok.Lexeme, idNode.Tok.Pos, typ}}
}

type Var struct {
	object
	kind VarKind
}

func (v *Var) String() string {
	typeStr := "nil"
	if !common.IsNil(v.typ) {
		typeStr = v.typ.String()
	}
	return /*v.kind.String()*/ "var" + " " + v.Name() + " " + typeStr // + " : " + v.Pos().String()
}

func NewVar(parent *Scope, kind VarKind, idNode *common.IdentifierNode, typ common.Type) *Var {
	return &Var{object: object{parent, idNode.Tok.Lexeme, idNode.Tok.Pos, typ}, kind: kind}
}

type Func struct {
	object
}

func (f *Func) String() string { return "func " + f.Name() + f.typ.String() }

func NewFunc(parent *Scope, typ *common.SignatureType, funcNode *common.FuncDeclNode) *Func {
	return &Func{object{parent, funcNode.Name.Tok.Lexeme, funcNode.Name.Tok.Pos, typ}}
}

type TypeName struct {
	object
}

func NewTypeName(parent *Scope, idNode *common.IdentifierNode, typ common.Type) *Const {
	return &Const{object{parent, idNode.Tok.Lexeme, idNode.Tok.Pos, typ}}
}

func (t *TypeName) String() string { return "type " + t.Name() + t.typ.String() }
