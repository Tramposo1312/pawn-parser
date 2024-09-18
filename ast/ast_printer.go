package ast

import (
	"fmt"
	"strings"
)

type AstPrinter struct {
	indentLevel int
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{indentLevel: 0}
}

func (ap *AstPrinter) Print(node Node) string {
	switch n := node.(type) {
	case *Program:
		return ap.visitProgram(n)
	case *Identifier:
		return ap.visitIdentifier(n)
	case *LetStatement:
		return ap.visitLetStatement(n)
	case *ReturnStatement:
		return ap.visitReturnStatement(n)
	case *ExpressionStatement:
		return ap.visitExpressionStatement(n)
	case *BlockStatement:
		return ap.visitBlockStatement(n)
	case *IfStatement:
		return ap.visitIfStatement(n)
	case *WhileStatement:
		return ap.visitWhileStatement(n)
	case *ForStatement:
		return ap.visitForStatement(n)
	case *PrefixExpression:
		return ap.visitPrefixExpression(n)
	case *InfixExpression:
		return ap.visitInfixExpression(n)
	case *CallExpression:
		return ap.visitCallExpression(n)
	case *IndexExpression:
		return ap.visitIndexExpression(n)
	case *IntegerLiteral:
		return ap.visitIntegerLiteral(n)
	case *FloatLiteral:
		return ap.visitFloatLiteral(n)
	case *StringLiteral:
		return ap.visitStringLiteral(n)
	case *BooleanLiteral:
		return ap.visitBooleanLiteral(n)
	case *NullLiteral:
		return ap.visitNullLiteral(n)
	case *ArrayLiteral:
		return ap.visitArrayLiteral(n)
	case *FunctionLiteral:
		return ap.visitFunctionLiteral(n)
	case *TypeName:
		return ap.visitTypeName(n)
	case *ArrayType:
		return ap.visitArrayType(n)
	case *FunctionType:
		return ap.visitFunctionType(n)
	case *TaggedType:
		return ap.visitTaggedType(n)
	case *TagDeclaration:
		return ap.visitTagDeclaration(n)
	case *EnumDeclaration:
		return ap.visitEnumDeclaration(n)
	case *IncludeDirective:
		return ap.visitIncludeDirective(n)
	case *DefineDirective:
		return ap.visitDefineDirective(n)
	case *IfDefDirective:
		return ap.visitIfDefDirective(n)
	case *NativeFunctionDeclaration:
		return ap.visitNativeFunctionDeclaration(n)
	case *StateDeclaration:
		return ap.visitStateDeclaration(n)
	case *FunctionDeclaration:
		return ap.visitFunctionDeclaration(n)
	default:
		return fmt.Sprintf("Unknown node type: %T", n)
	}
}

func (ap *AstPrinter) visitProgram(program *Program) string {
	var out strings.Builder
	out.WriteString("Program\n")
	ap.indentLevel++
	for _, stmt := range program.Statements {
		out.WriteString(ap.indent())
		out.WriteString(ap.Print(stmt))
		out.WriteString("\n")
	}
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitIdentifier(id *Identifier) string {
	return fmt.Sprintf("Identifier(%s)", id.Value)
}

func (ap *AstPrinter) visitLetStatement(ls *LetStatement) string {
	return fmt.Sprintf("LetStatement(Name: %s, Value: %s)", ap.Print(ls.Name), ap.Print(ls.Value))
}

func (ap *AstPrinter) visitReturnStatement(rs *ReturnStatement) string {
	return fmt.Sprintf("ReturnStatement(Value: %s)", ap.Print(rs.ReturnValue))
}

func (ap *AstPrinter) visitExpressionStatement(es *ExpressionStatement) string {
	return fmt.Sprintf("ExpressionStatement(%s)", ap.Print(es.Expression))
}

func (ap *AstPrinter) visitBlockStatement(bs *BlockStatement) string {
	var out strings.Builder
	out.WriteString("BlockStatement\n")
	ap.indentLevel++
	for _, stmt := range bs.Statements {
		out.WriteString(ap.indent())
		out.WriteString(ap.Print(stmt))
		out.WriteString("\n")
	}
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitIfStatement(is *IfStatement) string {
	var out strings.Builder
	out.WriteString("IfStatement\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Condition: ")
	out.WriteString(ap.Print(is.Condition))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Consequence: ")
	out.WriteString(ap.Print(is.Consequence))
	if is.Alternative != nil {
		out.WriteString("\n")
		out.WriteString(ap.indent())
		out.WriteString("Alternative: ")
		out.WriteString(ap.Print(is.Alternative))
	}
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitWhileStatement(ws *WhileStatement) string {
	var out strings.Builder
	out.WriteString("WhileStatement\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Condition: ")
	out.WriteString(ap.Print(ws.Condition))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Body: ")
	out.WriteString(ap.Print(ws.Body))
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitForStatement(fs *ForStatement) string {
	var out strings.Builder
	out.WriteString("ForStatement\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Init: ")
	out.WriteString(ap.Print(fs.Init))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Condition: ")
	out.WriteString(ap.Print(fs.Condition))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Update: ")
	out.WriteString(ap.Print(fs.Update))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Body: ")
	out.WriteString(ap.Print(fs.Body))
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitPrefixExpression(pe *PrefixExpression) string {
	return fmt.Sprintf("PrefixExpression(Operator: %s, Right: %s)", pe.Operator, ap.Print(pe.Right))
}

func (ap *AstPrinter) visitInfixExpression(ie *InfixExpression) string {
	return fmt.Sprintf("InfixExpression(Left: %s, Operator: %s, Right: %s)", ap.Print(ie.Left), ie.Operator, ap.Print(ie.Right))
}

func (ap *AstPrinter) visitCallExpression(ce *CallExpression) string {
	var out strings.Builder
	out.WriteString("CallExpression\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Function: ")
	out.WriteString(ap.Print(ce.Function))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Arguments:\n")
	ap.indentLevel++
	for _, arg := range ce.Arguments {
		out.WriteString(ap.indent())
		out.WriteString(ap.Print(arg))
		out.WriteString("\n")
	}
	ap.indentLevel -= 2
	return out.String()
}

func (ap *AstPrinter) visitIndexExpression(ie *IndexExpression) string {
	return fmt.Sprintf("IndexExpression(Left: %s, Index: %s)", ap.Print(ie.Left), ap.Print(ie.Index))
}

func (ap *AstPrinter) visitIntegerLiteral(il *IntegerLiteral) string {
	return fmt.Sprintf("IntegerLiteral(%d)", il.Value)
}

func (ap *AstPrinter) visitFloatLiteral(fl *FloatLiteral) string {
	return fmt.Sprintf("FloatLiteral(%f)", fl.Value)
}

func (ap *AstPrinter) visitStringLiteral(sl *StringLiteral) string {
	return fmt.Sprintf("StringLiteral(%s)", sl.Value)
}

func (ap *AstPrinter) visitBooleanLiteral(bl *BooleanLiteral) string {
	return fmt.Sprintf("BooleanLiteral(%t)", bl.Value)
}

func (ap *AstPrinter) visitNullLiteral(nl *NullLiteral) string {
	return "NullLiteral"
}

func (ap *AstPrinter) visitArrayLiteral(al *ArrayLiteral) string {
	var out strings.Builder
	out.WriteString("ArrayLiteral\n")
	ap.indentLevel++
	for _, elem := range al.Elements {
		out.WriteString(ap.indent())
		out.WriteString(ap.Print(elem))
		out.WriteString("\n")
	}
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitFunctionLiteral(fl *FunctionLiteral) string {
	var out strings.Builder
	out.WriteString("FunctionLiteral\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Parameters: ")
	for i, param := range fl.Parameters {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(ap.Print(param))
	}
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Body: ")
	out.WriteString(ap.Print(fl.Body))
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitTypeName(tn *TypeName) string {
	return fmt.Sprintf("TypeName(%s)", tn.Name)
}

func (ap *AstPrinter) visitArrayType(at *ArrayType) string {
	return fmt.Sprintf("ArrayType(ElementType: %s)", ap.Print(at.ElementType))
}

func (ap *AstPrinter) visitFunctionType(ft *FunctionType) string {
	var out strings.Builder
	out.WriteString("FunctionType(")
	for i, param := range ft.Parameters {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(ap.Print(param))
	}
	out.WriteString(") -> ")
	out.WriteString(ap.Print(ft.ReturnType))
	return out.String()
}

func (ap *AstPrinter) visitTaggedType(tt *TaggedType) string {
	return fmt.Sprintf("TaggedType(Tag: %s, Type: %s)", ap.Print(tt.Tag), ap.Print(tt.Type))
}

func (ap *AstPrinter) visitTagDeclaration(td *TagDeclaration) string {
	return fmt.Sprintf("TagDeclaration(%s)", ap.Print(td.Name))
}

func (ap *AstPrinter) visitEnumDeclaration(ed *EnumDeclaration) string {
	var out strings.Builder
	out.WriteString("EnumDeclaration\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Name: ")
	out.WriteString(ap.Print(ed.Name))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Members:\n")
	ap.indentLevel++
	for _, member := range ed.Members {
		out.WriteString(ap.indent())
		out.WriteString(ap.Print(member.Name))
		if member.Value != nil {
			out.WriteString(" = ")
			out.WriteString(ap.Print(member.Value))
		}
		out.WriteString("\n")
	}
	ap.indentLevel -= 2
	return out.String()
}

func (ap *AstPrinter) visitIncludeDirective(id *IncludeDirective) string {
	return fmt.Sprintf("IncludeDirective(%s)", id.Path)
}

func (ap *AstPrinter) visitDefineDirective(dd *DefineDirective) string {
	return fmt.Sprintf("DefineDirective(Name: %s, Value: %s)", dd.Name, ap.Print(dd.Value))
}

func (ap *AstPrinter) visitIfDefDirective(idd *IfDefDirective) string {
	var out strings.Builder
	out.WriteString("IfDefDirective\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Condition: ")
	out.WriteString(idd.Condition)
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Body:\n")
	ap.indentLevel++
	for _, stmt := range idd.Body {
		out.WriteString(ap.indent())
		out.WriteString(ap.Print(stmt))
		out.WriteString("\n")
	}
	ap.indentLevel--
	if idd.ElseBody != nil {
		out.WriteString(ap.indent())
		out.WriteString("ElseBody:\n")
		ap.indentLevel++
		for _, stmt := range idd.ElseBody {
			out.WriteString(ap.indent())
			out.WriteString(ap.Print(stmt))
			out.WriteString("\n")
		}
		ap.indentLevel--
	}
	ap.indentLevel--
	return out.String()
}

func (ap *AstPrinter) visitNativeFunctionDeclaration(nfd *NativeFunctionDeclaration) string {
	var out strings.Builder
	out.WriteString("NativeFunctionDeclaration\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Name: ")
	out.WriteString(ap.Print(nfd.Name))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Parameters: ")
	for i, param := range nfd.Parameters {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(ap.Print(param))
	}
	if nfd.ReturnType != nil {
		out.WriteString("\n")
		out.WriteString(ap.indent())
		out.WriteString("ReturnType: ")
		out.WriteString(ap.Print(nfd.ReturnType))
	}
	ap.indentLevel--
	return out.String()
}
func (ap *AstPrinter) visitStateDeclaration(sd *StateDeclaration) string {
	var out strings.Builder
	out.WriteString("StateDeclaration\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Name: ")
	out.WriteString(ap.Print(sd.Name))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Body: ")
	out.WriteString(ap.Print(sd.Body))
	ap.indentLevel--
	return out.String()
}
func (ap *AstPrinter) visitFunctionDeclaration(fd *FunctionDeclaration) string {
	var out strings.Builder
	out.WriteString("FunctionDeclaration\n")
	ap.indentLevel++
	out.WriteString(ap.indent())
	out.WriteString("Name: ")
	out.WriteString(ap.Print(fd.Name))
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Parameters: ")
	for i, param := range fd.Parameters {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(ap.Print(param))
	}
	out.WriteString("\n")
	out.WriteString(ap.indent())
	out.WriteString("Body: ")
	out.WriteString(ap.Print(fd.Body))
	ap.indentLevel--
	return out.String()
}
func (ap *AstPrinter) indent() string {
	return strings.Repeat("  ", ap.indentLevel)
}
