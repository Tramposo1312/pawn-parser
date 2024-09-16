package parser

import (
	"testing"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/lexer"
)

func TestFunctionDefinition(t *testing.T) {
	input := `
    function add(x, y) {
        return x + y;
    }
    `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.FunctionDefinition)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.FunctionDefinition. got=%T",
			program.Statements[0])
	}

	if stmt.Name.Value != "add" {
		t.Errorf("function name wrong. expected 'add', got=%q", stmt.Name.Value)
	}

	if len(stmt.Parameters) != 2 {
		t.Fatalf("function parameters wrong. expected 2, got=%d",
			len(stmt.Parameters))
	}

	if stmt.Parameters[0].Value != "x" {
		t.Errorf("parameter[0] is not 'x'. got=%q", stmt.Parameters[0].Value)
	}

	if stmt.Parameters[1].Value != "y" {
		t.Errorf("parameter[1] is not 'y'. got=%q", stmt.Parameters[1].Value)
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
