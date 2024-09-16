package parser

import (
	"testing"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/lexer"
)

func TestParseVariableDeclaration(t *testing.T) {
	input := `new myVar = 10;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	// Ensure that the program has exactly one statement
	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got %T", program.Statements[0])
	}

	// Check the token and identifier
	if stmt.Token.Literal != "new" {
		t.Errorf("expected token literal 'new', got %s", stmt.Token.Literal)
	}

	if stmt.Name.Value != "myVar" {
		t.Errorf("expected identifier 'myVar', got %s", stmt.Name.Value)
	}

	// You can also extend this to verify the value, if applicable
}

func TestPawnParse(t *testing.T) {
	input := `
	new x = 5;
	new y = 10;
	new z = x + y;
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements, got %d", len(program.Statements))
	}

	// Check first statement
	stmt1, ok := program.Statements[0].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got %T", program.Statements[0])
	}
	if stmt1.Token.Literal != "new" {
		t.Errorf("expected token literal 'new', got %s", stmt1.Token.Literal)
	}
	if stmt1.Name.Value != "x" {
		t.Errorf("expected identifier 'x', got %s", stmt1.Name.Value)
	}

	// Check second statement
	stmt2, ok := program.Statements[1].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got %T", program.Statements[1])
	}
	if stmt2.Token.Literal != "new" {
		t.Errorf("expected token literal 'new', got %s", stmt2.Token.Literal)
	}
	if stmt2.Name.Value != "y" {
		t.Errorf("expected identifier 'y', got %s", stmt2.Name.Value)
	}

	// Check third statement
	stmt3, ok := program.Statements[2].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got %T", program.Statements[2])
	}
	if stmt3.Token.Literal != "new" {
		t.Errorf("expected token literal 'new', got %s", stmt3.Token.Literal)
	}
	if stmt3.Name.Value != "z" {
		t.Errorf("expected identifier 'z', got %s", stmt3.Name.Value)
	}
}
