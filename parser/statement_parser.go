package parser

import (
	"fmt"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/precedence"
	"github.com/Tramposo1312/pawn-parser/token"
)

// The entry point for parsing any statement
func (p *Parser) parseStatement() (ast.Statement, error) {
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

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil, fmt.Errorf("expected identifier after 'new', got %s", p.peekToken.Type)
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil, fmt.Errorf("expected '=' after identifier in let statement, got %s", p.peekToken.Type)
	}

	p.nextToken()

	var err error
	stmt.Value, err = p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression in let statement: %v", err)
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt, nil
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	var err error
	stmt.ReturnValue, err = p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, fmt.Errorf("failed to parse return value: %v", err)
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt, nil
}

func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	var err error
	stmt.Expression, err = p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression in expression statement: %v", err)
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt, nil
}

func (p *Parser) parseIfStatement() (*ast.IfStatement, error) {
	stmt := &ast.IfStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil, fmt.Errorf("expected ( after if")
	}

	p.nextToken()
	condition, err := p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, err
	}
	stmt.Condition = condition

	if !p.expectPeek(token.RPAREN) {
		return nil, fmt.Errorf("expected ) after if condition")
	}

	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected { after if condition")
	}

	consequence, err := p.parseBlockStatement()
	if err != nil {
		return nil, err // This will now catch the EOF error from parseBlockStatement
	}
	stmt.Consequence = consequence

	return stmt, nil
}

func (p *Parser) parseWhileStatement() (*ast.WhileStatement, error) {
	stmt := &ast.WhileStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil, fmt.Errorf("expected '(' after 'while', got %s", p.peekToken.Type)
	}

	p.nextToken()
	var err error
	stmt.Condition, err = p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, fmt.Errorf("failed to parse while condition: %v", err)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, fmt.Errorf("expected ')' after while condition, got %s", p.peekToken.Type)
	}

	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected '{' to start while block, got %s", p.peekToken.Type)
	}

	stmt.Body, err = p.parseBlockStatement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse while body: %v", err)
	}

	return stmt, nil
}

func (p *Parser) parseForStatement() (*ast.ForStatement, error) {
	stmt := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil, fmt.Errorf("expected '(' after 'for', got %s", p.peekToken.Type)
	}

	// Parse initialization
	p.nextToken()
	var err error
	stmt.Init, err = p.parseStatement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse for init statement: %v", err)
	}

	if !p.expectPeek(token.SEMICOLON) {
		return nil, fmt.Errorf("expected ';' after for init statement, got %s", p.peekToken.Type)
	}

	// Parse condition
	p.nextToken()
	stmt.Condition, err = p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, fmt.Errorf("failed to parse for condition: %v", err)
	}

	if !p.expectPeek(token.SEMICOLON) {
		return nil, fmt.Errorf("expected ';' after for condition, got %s", p.peekToken.Type)
	}

	// Parse update
	p.nextToken()
	stmt.Update, err = p.parseExpressionStatement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse for update statement: %v", err)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, fmt.Errorf("expected ')' after for clauses, got %s", p.peekToken.Type)
	}

	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected '{' to start for block, got %s", p.peekToken.Type)
	}

	stmt.Body, err = p.parseBlockStatement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse for body: %v", err)
	}

	return stmt, nil
}

func (p *Parser) parseTagDeclaration() (*ast.TagDeclaration, error) {
	decl := &ast.TagDeclaration{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil, fmt.Errorf("expected identifier after 'tag', got %s", p.peekToken.Type)
	}

	decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return decl, nil
}

func (p *Parser) parseEnumDeclaration() (*ast.EnumDeclaration, error) {
	decl := &ast.EnumDeclaration{Token: p.curToken}

	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected '{' to start enum block, got %s", p.peekToken.Type)
	}

	var err error
	decl.Members, err = p.parseEnumMembers()
	if err != nil {
		return nil, fmt.Errorf("failed to parse enum members: %v", err)
	}

	if !p.expectPeek(token.RBRACE) {
		return nil, fmt.Errorf("expected '}' to end enum block, got %s", p.peekToken.Type)
	}

	return decl, nil
}

func (p *Parser) parseEnumMembers() ([]*ast.EnumMember, error) {
	members := []*ast.EnumMember{}

	for !p.curTokenIs(token.RBRACE) {
		member := &ast.EnumMember{}

		if !p.expectPeek(token.IDENT) {
			return nil, fmt.Errorf("expected identifier for enum member, got %s", p.peekToken.Type)
		}

		member.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken() // consume '='
			p.nextToken() // move to the value
			var err error
			member.Value, err = p.parseExpression(precedence.LOWEST)
			if err != nil {
				return nil, fmt.Errorf("failed to parse enum member value: %v", err)
			}
		}

		members = append(members, member)

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil, fmt.Errorf("expected ',' or '}' after enum member, got %s", p.peekToken.Type)
		}
	}

	return members, nil
}

func (p *Parser) parseBlockStatement() (*ast.BlockStatement, error) {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		block.Statements = append(block.Statements, stmt)
		p.nextToken()
	}

	if p.curTokenIs(token.EOF) {
		return nil, fmt.Errorf("unexpected EOF, expected } to close if block")
	}

	return block, nil
}
