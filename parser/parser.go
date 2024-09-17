// parser/parser.go

package parser

import (
	"fmt"
	"strings"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/lexer"
	"github.com/Tramposo1312/pawn-parser/precedence"
	"github.com/Tramposo1312/pawn-parser/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (p *Parser) TokenPrecedence(tokenType token.TokenType) int {
	return precedence.GetPrecedence(tokenType)
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},

		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}

	// Read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.TRUE, p.parseBooleanLiteral)
	p.registerPrefix(token.FALSE, p.parseBooleanLiteral)
	p.registerPrefix(token.NULL, p.parseNullLiteral)
	p.registerPrefix(token.LBRACK, p.parseArrayLiteral)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.INT, p.parseLiteral)
	p.registerPrefix(token.FLOAT, p.parseLiteral)
	p.registerPrefix(token.STRING, p.parseLiteral)
	p.registerPrefix(token.TRUE, p.parseLiteral)
	p.registerPrefix(token.FALSE, p.parseLiteral)
	p.registerPrefix(token.NULL, p.parseLiteral)
	p.registerPrefix(token.LBRACK, p.parseLiteral)
	p.registerPrefix(token.FUNCTION, p.parseLiteral)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.QUO, p.parseInfixExpression)
	p.registerInfix(token.MUL, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NEQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	var errors []string

	for !p.curTokenIs(token.EOF) {
		stmt, err := p.parseStatement()
		if err != nil {
			errors = append(errors, err.Error())
		} else {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("parser errors:\n%s", strings.Join(errors, "\n"))
	}

	return program, nil
}

func (p *Parser) peekPrecedence() int {
	return precedence.GetPrecedence(p.peekToken.Type)
}

func (p *Parser) curPrecedence() int {
	return precedence.GetPrecedence(p.curToken.Type)
}
