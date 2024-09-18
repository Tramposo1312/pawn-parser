package parser

import (
	"fmt"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/precedence"
	"github.com/Tramposo1312/pawn-parser/token"
)

type (
	prefixParseFn func() (ast.Expression, error)
	infixParseFn  func(ast.Expression) (ast.Expression, error)
)

func (p *Parser) parseIdentifier() (ast.Expression, error) {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}, nil
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) error {
	return fmt.Errorf("no prefix parse function for %s found", t)
}

func (p *Parser) parsePrefixExpression() (ast.Expression, error) {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	right, err := p.parseExpression(precedence.PREFIX)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression after prefix operator %s: %v", expression.Operator, err)
	}
	expression.Right = right

	return expression, nil
}

func (p *Parser) parseGroupedExpression() (ast.Expression, error) {
	p.nextToken()

	exp, err := p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, err
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, fmt.Errorf("expected ), got %s", p.peekToken.Type)
	}

	return exp, nil
}

func (p *Parser) parseCallExpression(function ast.Expression) (ast.Expression, error) {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	args, err := p.parseExpressionList(token.RPAREN)
	if err != nil {
		return nil, err
	}
	exp.Arguments = args
	return exp, nil
}

func (p *Parser) parseExpressionList(end token.TokenType) ([]ast.Expression, error) {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list, nil
	}

	p.nextToken()
	expr, err := p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression in list: %v", err)
	}
	list = append(list, expr)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		expr, err := p.parseExpression(precedence.LOWEST)
		if err != nil {
			return nil, fmt.Errorf("failed to parse expression after comma in list: %v", err)
		}
		list = append(list, expr)
	}

	if !p.expectPeek(end) {
		return nil, fmt.Errorf("expected %s, got %s", end, p.peekToken.Type)
	}

	return list, nil
}

func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil, p.noPrefixParseFnError(p.curToken.Type)
	}
	leftExp, err := prefix()
	if err != nil {
		return nil, err
	}

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp, nil
		}
		p.nextToken()
		leftExp, err = infix(leftExp)
		if err != nil {
			return nil, err
		}
	}

	return leftExp, nil
}
func (p *Parser) parseInfixExpression(left ast.Expression) (ast.Expression, error) {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	right, err := p.parseExpression(precedence)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression after operator %s: %v", expression.Operator, err)
	}
	expression.Right = right

	return expression, nil
}
