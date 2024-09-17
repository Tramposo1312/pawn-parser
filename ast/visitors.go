package ast

type Visitor interface {
	VisitProgram(node *Program) interface{}
	VisitIdentifier(node *Identifier) interface{}
	VisitLetStatement(node *LetStatement) interface{}
	VisitReturnStatement(node *ReturnStatement) interface{}
	VisitExpressionStatement(node *ExpressionStatement) interface{}
	VisitBlockStatement(node *BlockStatement) interface{}
	VisitIfStatement(node *IfStatement) interface{}
	VisitWhileStatement(node *WhileStatement) interface{}
	VisitForStatement(node *ForStatement) interface{}
	VisitPrefixExpression(node *PrefixExpression) interface{}
	VisitInfixExpression(node *InfixExpression) interface{}
	VisitCallExpression(node *CallExpression) interface{}
	VisitIndexExpression(node *IndexExpression) interface{}
	VisitIntegerLiteral(node *IntegerLiteral) interface{}
	VisitFloatLiteral(node *FloatLiteral) interface{}
	VisitStringLiteral(node *StringLiteral) interface{}
	VisitBooleanLiteral(node *BooleanLiteral) interface{}
	VisitNullLiteral(node *NullLiteral) interface{}
	VisitArrayLiteral(node *ArrayLiteral) interface{}
	VisitFunctionLiteral(node *FunctionLiteral) interface{}
	VisitTypeName(node *TypeName) interface{}
	VisitArrayType(node *ArrayType) interface{}
	VisitFunctionType(node *FunctionType) interface{}
	VisitTaggedType(node *TaggedType) interface{}
	VisitTagDeclaration(node *TagDeclaration) interface{}
	VisitEnumDeclaration(node *EnumDeclaration) interface{}
}

func (p *Program) Accept(v Visitor) interface{} {
	return v.VisitProgram(p)
}

func (i *Identifier) Accept(v Visitor) interface{} {
	return v.VisitIdentifier(i)
}

func (ls *LetStatement) Accept(v Visitor) interface{} {
	return v.VisitLetStatement(ls)
}

func (rs *ReturnStatement) Accept(v Visitor) interface{} {
	return v.VisitReturnStatement(rs)
}

func (es *ExpressionStatement) Accept(v Visitor) interface{} {
	return v.VisitExpressionStatement(es)
}

func (bs *BlockStatement) Accept(v Visitor) interface{} {
	return v.VisitBlockStatement(bs)
}

func (is *IfStatement) Accept(v Visitor) interface{} {
	return v.VisitIfStatement(is)
}

func (ws *WhileStatement) Accept(v Visitor) interface{} {
	return v.VisitWhileStatement(ws)
}

func (fs *ForStatement) Accept(v Visitor) interface{} {
	return v.VisitForStatement(fs)
}

func (pe *PrefixExpression) Accept(v Visitor) interface{} {
	return v.VisitPrefixExpression(pe)
}

func (ie *InfixExpression) Accept(v Visitor) interface{} {
	return v.VisitInfixExpression(ie)
}

func (ce *CallExpression) Accept(v Visitor) interface{} {
	return v.VisitCallExpression(ce)
}

func (ie *IndexExpression) Accept(v Visitor) interface{} {
	return v.VisitIndexExpression(ie)
}

func (il *IntegerLiteral) Accept(v Visitor) interface{} {
	return v.VisitIntegerLiteral(il)
}

func (fl *FloatLiteral) Accept(v Visitor) interface{} {
	return v.VisitFloatLiteral(fl)
}

func (sl *StringLiteral) Accept(v Visitor) interface{} {
	return v.VisitStringLiteral(sl)
}

func (bl *BooleanLiteral) Accept(v Visitor) interface{} {
	return v.VisitBooleanLiteral(bl)
}

func (nl *NullLiteral) Accept(v Visitor) interface{} {
	return v.VisitNullLiteral(nl)
}

func (al *ArrayLiteral) Accept(v Visitor) interface{} {
	return v.VisitArrayLiteral(al)
}

func (fl *FunctionLiteral) Accept(v Visitor) interface{} {
	return v.VisitFunctionLiteral(fl)
}

func (tn *TypeName) Accept(v Visitor) interface{} {
	return v.VisitTypeName(tn)
}

func (at *ArrayType) Accept(v Visitor) interface{} {
	return v.VisitArrayType(at)
}

func (ft *FunctionType) Accept(v Visitor) interface{} {
	return v.VisitFunctionType(ft)
}

func (tt *TaggedType) Accept(v Visitor) interface{} {
	return v.VisitTaggedType(tt)
}

func (td *TagDeclaration) Accept(v Visitor) interface{} {
	return v.VisitTagDeclaration(td)
}

func (ed *EnumDeclaration) Accept(v Visitor) interface{} {
	return v.VisitEnumDeclaration(ed)
}
