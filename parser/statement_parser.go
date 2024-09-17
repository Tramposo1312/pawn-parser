package parser

import (
	"fmt"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/token"
)

// The entry point for parsing any statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.NEW:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.FOR:
		return p.parseForStatement()
	case token.LBRACE:
		return p.parseBlockStatement()
	case token.TAG:
		return p.parseTagDeclaration()
	case token.ENUM:
		return p.parseEnumDeclaration()
	default:
		return p.parseExpressionStatement()
	}
}

// Parses a variable declaration statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		stmt.Alternative = p.parseBlockStatement()
	}

	return stmt
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse initialization
	p.nextToken()
	stmt.Init = p.parseStatement()

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	// Parse condition
	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	// Parse update
	p.nextToken()
	stmt.Update = p.parseExpressionStatement()

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseTagDeclaration() *ast.TagDeclaration {
	decl := &ast.TagDeclaration{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return decl
}

func (p *Parser) parseEnumDeclaration() *ast.EnumDeclaration {
	decl := &ast.EnumDeclaration{Token: p.curToken}

	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	decl.Members = p.parseEnumMembers()

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return decl
}

func (p *Parser) parseEnumMembers() []*ast.EnumMember {
	members := []*ast.EnumMember{}

	for !p.curTokenIs(token.RBRACE) {
		member := &ast.EnumMember{}

		if !p.expectPeek(token.IDENT) {
			return nil
		}

		member.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken() // consume '='
			p.nextToken() // move to the value
			member.Value = p.parseExpression(LOWEST)
		}

		members = append(members, member)

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	return members
}

// Checks if the next token is of the expected type
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// Adds an error to the parser's error list
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Checks if the next token is of the given type
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Checks if the current token is of the given type
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
