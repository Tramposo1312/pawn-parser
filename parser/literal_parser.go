package parser

import (
	"fmt"
	"strconv"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/token"
)

// The entry point for parsing any literal
func (p *Parser) parseLiteral() (ast.Expression, error) {
	switch p.curToken.Type {
	case token.INT:
		return p.parseIntegerLiteral()
	case token.FLOAT:
		return p.parseFloatLiteral()
	case token.STRING:
		return p.parseStringLiteral()
	case token.TRUE, token.FALSE:
		return p.parseBooleanLiteral()
	case token.NULL:
		return p.parseNullLiteral()
	case token.LBRACK:
		return p.parseArrayLiteral()
	case token.FUNCTION:
		return p.parseFunctionLiteral()
	default:
		p.errors = append(p.errors, fmt.Sprintf("unexpected token %s while parsing literal", p.curToken.Type))
		return nil, nil
	}
}

func (p *Parser) parseIntegerLiteral() (ast.Expression, error) {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as integer", p.curToken.Literal)
	}

	lit.Value = value
	return lit, nil
}

func (p *Parser) parseFloatLiteral() (ast.Expression, error) {
	lit := &ast.FloatLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as float", p.curToken.Literal)
	}

	lit.Value = value
	return lit, nil
}

func (p *Parser) parseStringLiteral() (ast.Expression, error) {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}, nil
}

func (p *Parser) parseBooleanLiteral() (ast.Expression, error) {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curToken.Type == token.TRUE}, nil
}

func (p *Parser) parseNullLiteral() (ast.Expression, error) {
	return &ast.NullLiteral{Token: p.curToken}, nil
}

func (p *Parser) parseArrayLiteral() (ast.Expression, error) {
	array := &ast.ArrayLiteral{Token: p.curToken}
	elements, err := p.parseExpressionList(token.RBRACK)
	if err != nil {
		return nil, err
	}
	array.Elements = elements
	return array, nil
}

func (p *Parser) parseFunctionLiteral() (ast.Expression, error) {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil, fmt.Errorf("expected ( after function keyword")
	}

	params, err := p.parseFunctionParameters()
	if err != nil {
		return nil, err
	}
	lit.Parameters = params

	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected { after function parameters")
	}

	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}
	lit.Body = body

	return lit, nil
}

func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, error) {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers, nil
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, fmt.Errorf("expected ) after function parameters")
	}

	return identifiers, nil
}
