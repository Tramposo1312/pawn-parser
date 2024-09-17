package ast

import (
	"bytes"
	"strings"

	"github.com/Tramposo1312/pawn-parser/precedence"
	"github.com/Tramposo1312/pawn-parser/token"
)

type PrecedenceProvider interface {
	TokenPrecedence(token.TokenType) int
}

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // The operator token e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	leftParen := false
	rightParen := false

	if left, ok := ie.Left.(*InfixExpression); ok {
		leftParen = precedence.GetPrecedenceFromString(ie.Operator) > precedence.GetPrecedenceFromString(left.Operator)
	}

	if right, ok := ie.Right.(*InfixExpression); ok {
		rightParen = precedence.GetPrecedenceFromString(ie.Operator) >= precedence.GetPrecedenceFromString(right.Operator)
	}

	if leftParen {
		out.WriteString("(")
	}
	out.WriteString(ie.Left.String())
	if leftParen {
		out.WriteString(")")
	}

	out.WriteString(" " + ie.Operator + " ")

	if rightParen {
		out.WriteString("(")
	}
	out.WriteString(ie.Right.String())
	if rightParen {
		out.WriteString(")")
	}

	return out.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}
