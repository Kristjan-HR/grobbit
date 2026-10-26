package semana

import . "Grobbit/common"

func DeclarePkgName(scope *Scope, impNode *ImportSpecNode) {
	obj := NewPkgName(nil, impNode, nil)
	if alt := scope.Insert(obj); alt != nil {
		// NOTE: Already declared (note short var)
	}
}

func DeclareConst(v *SAVisitor, scope *Scope, typ Type, idNode *IdentifierNode) {
	const msgRedefined = "%s redeclared in this block"
	obj := NewConst(nil, idNode, typ) // NOTE  pkg/scope
	if alt := scope.Insert(obj); alt != nil {
		v.logError(idNode.Tok.Pos, msgRedefined, idNode.Tok.Lexeme)
	}
}

func DeclareConsts(v *SAVisitor, scope *Scope, typ Type, ids ...*IdentifierNode) {
	for _, id := range ids {
		DeclareConst(v, scope, typ, id)
	}
}

func DeclareVar(v *SAVisitor, scope *Scope, kind VarKind, typ Type, idNode *IdentifierNode) {
	const msgRedefined = "%s redeclared in this block"
	obj := NewVar(nil, kind, idNode, typ)
	if alt := scope.Insert(obj); alt != nil {
		v.logError(idNode.Tok.Pos, msgRedefined, idNode.Tok.Lexeme)
	}
}

func DeclareVars(v *SAVisitor, scope *Scope, kind VarKind, typ Type, ids ...*IdentifierNode) {
	for _, id := range ids {
		DeclareVar(v, scope, kind, typ, id)
	}
}

func DeclareValueSpec(v *SAVisitor, scope *Scope, kind VarKind, node *ValueSpecNode, flags SAFlags) {
	const msgSpecMismatch1 = "assignment mismatch: %d variable(s) but %d value(s)"
	const msgSpecMismatch2 = "extra init expr %d"

	if len(node.Names) > len(node.Values) && len(node.Values) > 0 {
		v.logError(node.Token.Pos, msgSpecMismatch1, len(node.Names), len(node.Values))
	} else if len(node.Names) < len(node.Values) {
		v.logError(node.Token.Pos, msgSpecMismatch2, len(node.Values))
	}
	var typ Type
	for i, idNode := range node.Names {
		if !IsNil(node.Type) {
			typ = node.Type.GetType()
		} else if i < len(node.Values) {
			typ = node.Values[i].GetType()
		} else {
			typ = NewInvalidType()
		}
		if flags&InConstDecl != 0 {
			DeclareConst(v, scope, typ, idNode)
		} else {
			DeclareVar(v, scope, kind, typ, idNode)
		}
	}
}

func DeclareShortHand(v *SAVisitor, scope *Scope, node *AssignStmtNode) {
	const msgAssignMismatch1 = "assignment mismatch: %d variable(s) but %d value(s)"
	const msgDecl = "non-name on left side of :="
	if len(node.Lhs) != len(node.Rhs) {
		v.logError(node.Tok.Pos, msgAssignMismatch1, len(node.Lhs), len(node.Rhs))
	}
	for i, exprNode := range node.Lhs {
		var typ Type
		if i < len(node.Rhs) {
			typ = node.Rhs[i].GetType()
		} else {
			typ = NewInvalidType()
		}
		if idNode, ok := exprNode.(*IdentifierNode); ok {
			DeclareVar(v, scope, LocalVar, typ, idNode)
		} else {
			v.logError(node.Tok.Pos, msgDecl)
		}
	}
}

func DeclareFunc(scope *Scope, node *FuncDeclNode) {
	var paramsType, resultsType []Type
	for _, param := range node.Type.Params {
		paramsType = append(paramsType, param.Type.GetType())
	}
	if node.Type.Type != nil {
		resultsType = append(resultsType, node.Type.Type.GetType())
	}
	scope.Insert(NewFunc(nil, NewSignatureType(paramsType, resultsType), node))
}

func DeclareTypeName(v *SAVisitor, scope *Scope, typ Type, idNode *IdentifierNode) {
	const msgRedefined = "%s redeclared in this block"
	obj := NewTypeName(nil, idNode, typ)
	if alt := scope.Insert(obj); alt != nil {
		v.logError(idNode.Tok.Pos, msgRedefined, idNode.Tok.Lexeme)
	}
}

func posInScope(idDecl, idLook Position) bool {
	if idLook.Line == 0 && idLook.Col == 0 {
		return true
	}
	if idLook.Line < idDecl.Line {
		return false
	}
	if idLook.Line == idDecl.Line {
		return idLook.Col >= idDecl.Col
	}
	return true
}

func Resolve(scope *Scope, idNode *IdentifierNode, onlyGivenScope bool) (Object, *Scope) {
	for s := scope; s != nil; s = s.Parent() {
		obj := s.Lookup(idNode.Tok.Lexeme)
		if obj != nil && posInScope(obj.Pos(), idNode.Tok.Pos) {

			return obj, s
		}
		if onlyGivenScope {
			break
		}
	}
	return nil, nil
}

func NearestEnclosingScope(scope *Scope, comment string) *Scope {
	for s := scope; s != nil; s = s.Parent() {
		if s.comment == comment {
			return s
		}
	}
	return nil
}

func PopulateScope(v *SAVisitor, scope *Scope) {
	DeclareTypeName(v, scope, NewIntType(), NewIdentifier(TtIdentifier, "int"))
	DeclareTypeName(v, scope, NewBoolType(), NewIdentifier(TtIdentifier, "bool"))
	DeclareTypeName(v, scope, NewFloatType(), NewIdentifier(TtIdentifier, "float64"))
	//DeclareTypeName(v, scope, NewStringType(), NewIdentifier(TtIdentifier, "string"))
	//DeclareTypeName(v, scope, NewRuneType(), NewIdentifier(TtIdentifier, "rune"))
	//DeclareConst(v, scope, NewBoolType(), NewIdentifier(TtIdentifier, "true"))
	//DeclareConst(v, scope, NewBoolType(), NewIdentifier(TtIdentifier, "false"))
	//DeclareConst(v, scope, NewPointerType(nil), NewIdentifier(TtIdentifier, "nil"))

}

func IsConst(scope *Scope, idNode *IdentifierNode) bool {
	obj, _ := Resolve(scope, idNode, false)
	if obj != nil {
		_, ok := obj.(*Const)
		return ok
	}
	return false
}

func IsVariable(scope *Scope, idNode *IdentifierNode) bool {
	obj, _ := Resolve(scope, idNode, false)
	if obj != nil {
		_, ok := obj.(*Var)
		return ok
	}
	return false
}

func IsFunction(scope *Scope, idNode *IdentifierNode) bool {
	obj, _ := Resolve(scope, idNode, false)
	if obj != nil {
		_, ok := obj.(*Func)
		return ok
	}
	return false
}
func IsTypeName(scope *Scope, idNode *IdentifierNode) bool {
	obj, _ := Resolve(scope, idNode, false)
	if obj != nil {
		_, ok := obj.(*TypeName)
		return ok
	}
	return false
}

func IsAddressable(scope *Scope, node ExprNode) bool {
	switch nt := node.(type) {
	case *IdentifierNode:
		return IsVariable(scope, nt)
	default:
		return false
	}
}

func IsAssignable(scope *Scope, node ExprNode) bool {
	return IsAddressable(scope, node)
}
