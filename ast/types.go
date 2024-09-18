package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Tramposo1312/pawn-parser/token"
)

// Represents a named type.
type TypeName struct {
	Token token.Token
	Name  string
}

func (tn *TypeName) expressionNode()      {}
func (tn *TypeName) TokenLiteral() string { return tn.Token.Literal }
func (tn *TypeName) String() string       { return tn.Name }

type ArrayType struct {
	Token       token.Token // the '[' token
	ElementType Expression
}

func (at *ArrayType) expressionNode()      {}
func (at *ArrayType) TokenLiteral() string { return at.Token.Literal }
func (at *ArrayType) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	out.WriteString(at.ElementType.String())
	out.WriteString("]")
	return out.String()
}

type FunctionType struct {
	Token      token.Token // the 'function' token
	Parameters []Expression
	ReturnType Expression
}

func (ft *FunctionType) expressionNode()      {}
func (ft *FunctionType) TokenLiteral() string { return ft.Token.Literal }
func (ft *FunctionType) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range ft.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("function(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	if ft.ReturnType != nil {
		out.WriteString(" : ")
		out.WriteString(ft.ReturnType.String())
	}
	return out.String()
}

// Represents a tagged type in Pawn
type TaggedType struct {
	Token token.Token //  'tag'
	Tag   *Identifier
	Type  Expression
}

func (tt *TaggedType) expressionNode()      {}
func (tt *TaggedType) TokenLiteral() string { return tt.Token.Literal }
func (tt *TaggedType) String() string {
	return fmt.Sprintf("%s:%s", tt.Tag.String(), tt.Type.String())
}
