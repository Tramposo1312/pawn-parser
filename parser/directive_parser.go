package parser

import (
	"fmt"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/precedence"
	"github.com/Tramposo1312/pawn-parser/token"
)

func (p *Parser) parseIncludeDirective() (*ast.IncludeDirective, error) {
	directive := &ast.IncludeDirective{Token: p.curToken}

	if !p.expectPeek(token.LT) && !p.expectPeek(token.STRING) {
		return nil, fmt.Errorf("expected < or \" after #include, got %s", p.peekToken.Type)
	}

	if p.curTokenIs(token.LT) {
		p.nextToken() // consume
		var path string
		for !p.curTokenIs(token.GT) && !p.curTokenIs(token.EOF) {
			path += p.curToken.Literal
			p.nextToken()
		}
		if !p.curTokenIs(token.GT) {
			return nil, fmt.Errorf("expected > to close include path")
		}
		directive.Path = path
	} else {
		directive.Path = p.curToken.Literal
	}

	return directive, nil
}

func (p *Parser) parseDefineDirective() (*ast.DefineDirective, error) {
	directive := &ast.DefineDirective{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil, fmt.Errorf("expected identifier after #define, got %s", p.peekToken.Type)
	}

	directive.Name = p.curToken.Literal

	p.nextToken()
	value, err := p.parseExpression(precedence.LOWEST)
	if err != nil {
		return nil, err
	}

	directive.Value = value
	return directive, nil
}

func (p *Parser) parseIfDefDirective() (*ast.IfDefDirective, error) {
	directive := &ast.IfDefDirective{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil, fmt.Errorf("expected identifier after #ifdef, got %s", p.peekToken.Type)
	}

	directive.Condition = p.curToken.Literal

	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected { after #ifdef condition")
	}

	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}
	directive.Body = body.Statements

	if p.peekTokenIs(token.ELSE) {
		p.nextToken() // consume else
		if !p.expectPeek(token.LBRACE) {
			return nil, fmt.Errorf("expected { after else")
		}
		elseBody, err := p.parseBlockStatement()
		if err != nil {
			return nil, err
		}
		directive.ElseBody = elseBody.Statements
	}

	if !p.expectPeek(token.ENDIF) {
		return nil, fmt.Errorf("expected #endif, got %s", p.peekToken.Type)
	}

	return directive, nil
}

func (p *Parser) parseNativeFunctionDeclaration() (*ast.NativeFunctionDeclaration, error) {
	decl := &ast.NativeFunctionDeclaration{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil, fmt.Errorf("expected function name after 'native', got %s", p.peekToken.Type)
	}

	decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LPAREN) {
		return nil, fmt.Errorf("expected ( after function name, got %s", p.peekToken.Type)
	}

	params, err := p.parseFunctionParameters()
	if err != nil {
		return nil, err
	}
	decl.Parameters = params

	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume :
		p.nextToken() // move to return type
		returnType, err := p.parseExpression(precedence.LOWEST)
		if err != nil {
			return nil, err
		}
		decl.ReturnType = returnType
	}

	if !p.expectPeek(token.SEMICOLON) {
		return nil, fmt.Errorf("expected ; after native function declaration")
	}

	return decl, nil
}
func (p *Parser) parseFunctionDeclaration() (*ast.FunctionDeclaration, error) {
	decl := &ast.FunctionDeclaration{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil, fmt.Errorf("expected function name, got %s", p.peekToken.Type)
	}

	decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LPAREN) {
		return nil, fmt.Errorf("expected ( after function name, got %s", p.peekToken.Type)
	}

	params, err := p.parseFunctionParameters()
	if err != nil {
		return nil, err
	}
	decl.Parameters = params

	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected { to start function body, got %s", p.peekToken.Type)
	}

	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}
	decl.Body = body

	return decl, nil
}
func (p *Parser) parseStateDeclaration() (*ast.StateDeclaration, error) {
	decl := &ast.StateDeclaration{Token: p.curToken}

	// the state name
	if !p.expectPeek(token.IDENT) {
		return nil, fmt.Errorf("expected state name after 'state', got %s", p.peekToken.Type)
	}
	decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// opening brace
	if !p.expectPeek(token.LBRACE) {
		return nil, fmt.Errorf("expected '{' after state name, got %s", p.peekToken.Type)
	}

	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}
	decl.Body = body

	return decl, nil
}
func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, error) {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers, nil
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if p.peekTokenIs(token.COLON) {
		p.nextToken()
		ident.Value += ":" + p.curToken.Literal
	}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		if p.peekTokenIs(token.COLON) {
			p.nextToken()
			ident.Value += ":" + p.curToken.Literal
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, fmt.Errorf("expected ')' after function parameters")
	}

	return identifiers, nil
}
func (p *Parser) parseDirective() (ast.Statement, error) {
	switch p.peekToken.Literal {
	case "include":
		return p.parseIncludeDirective()
	case "define":
		return p.parseDefineDirective()
	case "ifdef":
		return p.parseIfDefDirective()
	default:
		return nil, fmt.Errorf("unknown directive: %s", p.peekToken.Literal)
	}
}
