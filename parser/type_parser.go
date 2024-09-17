package parser

import (
	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/token"
)

func (p *Parser) parseType() ast.Expression {
	switch p.curToken.Type {
	case token.IDENT:
		return p.parseTypeName()
	case token.LBRACK:
		return p.parseArrayType()
	case token.FUNCTION:
		return p.parseFunctionType()
	default:
		return p.parseTaggedType()
	}
}

func (p *Parser) parseTypeName() ast.Expression {
	return &ast.TypeName{Token: p.curToken, Name: p.curToken.Literal}
}

func (p *Parser) parseArrayType() ast.Expression {
	arrayType := &ast.ArrayType{Token: p.curToken}

	if !p.expectPeek(token.RBRACK) {
		return nil
	}

	p.nextToken() // Move past ']'

	// Parse the element type
	arrayType.ElementType = p.parseType()

	return arrayType
}

func (p *Parser) parseFunctionType() ast.Expression {
	funcType := &ast.FunctionType{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	funcType.Parameters = p.parseFunctionTypeParameters()

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	// I guess in Pawn function types don't typically specify a return type

	return funcType
}

func (p *Parser) parseFunctionTypeParameters() []ast.Expression {
	var params []ast.Expression

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken() // consume ')'
		return params
	}

	p.nextToken() // consume '('

	params = append(params, p.parseType())

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // consume ','
		p.nextToken() // move to next parameter
		params = append(params, p.parseType())
	}

	return params
}

func (p *Parser) parseTaggedType() ast.Expression {
	taggedType := &ast.TaggedType{Token: p.curToken}

	taggedType.Tag = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken() // move past ':'

	taggedType.Type = p.parseType()

	return taggedType
}
