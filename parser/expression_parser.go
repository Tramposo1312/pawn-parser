package parser

import (
	"fmt"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s found", t))
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)
	if expression.Right == nil {
		p.errors = append(p.errors, fmt.Sprintf("Expected expression after prefix operator %s", expression.Operator))
		return nil
	}

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	if expression.Right == nil {
		p.errors = append(p.errors, fmt.Sprintf("Expected expression after operator %s", expression.Operator))
		return nil
	}

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	expr := p.parseExpression(LOWEST)
	if expr == nil {
		p.errors = append(p.errors, "Expected expression in list")
		return nil
	}
	list = append(list, expr)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		expr := p.parseExpression(LOWEST)
		if expr == nil {
			p.errors = append(p.errors, "Expected expression after comma in list")
			return nil
		}
		list = append(list, expr)
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}
