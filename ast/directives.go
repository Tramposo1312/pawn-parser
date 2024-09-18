package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Tramposo1312/pawn-parser/token"
)

type IncludeDirective struct {
	Token token.Token //  '#include'
	Path  string
}

type DefineDirective struct {
	Token token.Token //  '#define'
	Name  string
	Value Expression
}

type IfDefDirective struct {
	Token     token.Token //  '#ifdef'
	Condition string
	Body      []Statement
	ElseBody  []Statement
}

type NativeFunctionDeclaration struct {
	Token      token.Token //  'native'
	Name       *Identifier
	Parameters []*Identifier
	ReturnType Expression
}

type StateDeclaration struct {
	Token token.Token //  'state'
	Name  *Identifier
	Body  *BlockStatement
}
type FunctionDeclaration struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

// ==== String()

func (id *IncludeDirective) String() string {
	return "#include " + id.Path
}

func (dd *DefineDirective) String() string {
	var out bytes.Buffer
	out.WriteString("#define ")
	out.WriteString(dd.Name)
	out.WriteString(" ")
	out.WriteString(dd.Value.String())
	return out.String()
}

func (idd *IfDefDirective) String() string {
	var out bytes.Buffer
	out.WriteString("#ifdef ")
	out.WriteString(idd.Condition)
	out.WriteString(" {\n")
	for _, s := range idd.Body {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	if idd.ElseBody != nil {
		out.WriteString("} else {\n")
		for _, s := range idd.ElseBody {
			out.WriteString(s.String())
			out.WriteString("\n")
		}
	}
	out.WriteString("}")
	return out.String()
}

func (nfd *NativeFunctionDeclaration) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range nfd.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("native ")
	out.WriteString(nfd.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	if nfd.ReturnType != nil {
		out.WriteString(": ")
		out.WriteString(nfd.ReturnType.String())
	}
	out.WriteString(";")
	return out.String()
}

func (sd *StateDeclaration) String() string {
	return fmt.Sprintf("state %s %s", sd.Name.String(), sd.Body.String())
}
func (fd *FunctionDeclaration) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fd.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fd.TokenLiteral() + " ")
	out.WriteString(fd.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fd.Body.String())
	return out.String()
}

// ====
func (id *IncludeDirective) statementNode()       {}
func (id *IncludeDirective) TokenLiteral() string { return id.Token.Literal }

func (dd *DefineDirective) statementNode()       {}
func (dd *DefineDirective) TokenLiteral() string { return dd.Token.Literal }

func (idd *IfDefDirective) statementNode()       {}
func (idd *IfDefDirective) TokenLiteral() string { return idd.Token.Literal }

func (nfd *NativeFunctionDeclaration) statementNode()       {}
func (nfd *NativeFunctionDeclaration) TokenLiteral() string { return nfd.Token.Literal }

func (sd *StateDeclaration) statementNode()       {}
func (sd *StateDeclaration) TokenLiteral() string { return sd.Token.Literal }

func (fd *FunctionDeclaration) statementNode()       {}
func (fd *FunctionDeclaration) TokenLiteral() string { return fd.Token.Literal }
