// parser/parser_test.go

package parser

import (
	"fmt"
	"testing"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/lexer"
)

func TestVariableDeclarations(t *testing.T) {
	input := `
    new x = 5;
    new y = 10.5;
    new str[] = "Hello, Pawn!";
    new ch = 'A';
    new flag = true;
    new empty = null;
    `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 6 {
		t.Fatalf("program.Statements does not contain 6 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedValue      interface{}
		expectedIsArray    bool
	}{
		{"x", int64(5), false},
		{"y", 10.5, false},
		{"str", "Hello, Pawn!", true},
		{"ch", byte('A'), false},
		{"flag", true, false},
		{"empty", nil, false},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVariableDeclaration(t, stmt, tt.expectedIdentifier, tt.expectedIsArray) {
			return
		}

		val := stmt.(*ast.VariableDeclaration).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func testVariableDeclaration(t *testing.T, s ast.Statement, name string, isArray bool) bool {
	if s.TokenLiteral() != "new" {
		t.Errorf("s.TokenLiteral not 'new'. got=%q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VariableDeclaration)
	if !ok {
		t.Errorf("s not *ast.VariableDeclaration. got=%T", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value not '%s'. got=%s", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("varStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, varStmt.Name.TokenLiteral())
		return false
	}

	if varStmt.IsArray != isArray {
		t.Errorf("varStmt.IsArray not %v. got=%v", isArray, varStmt.IsArray)
		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case float64:
		return testFloatLiteral(t, exp, v)
	case string:
		return testStringLiteral(t, exp, v)
	case byte:
		return testCharLiteral(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	case nil:
		return testNullLiteral(t, exp)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func TestArithmeticOperations(t *testing.T) {
	input := `
    new result = 10 + 5 * 2;
    new complex = (20 - 5) / 3 + 2 * 4;
    `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}
}

func TestFunctionDeclaration(t *testing.T) {
	input := `
    function add(a, b) {
        return a + b;
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

	stmt, ok := program.Statements[0].(*ast.FunctionDeclaration)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.FunctionDeclaration. got=%T",
			program.Statements[0])
	}

	if stmt.Name.Value != "add" {
		t.Errorf("function name was not 'add'. got=%q", stmt.Name.Value)
	}

	if len(stmt.Parameters) != 2 {
		t.Fatalf("function parameters wrong. want 2, got=%d",
			len(stmt.Parameters))
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

// ======================= HELPERS
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}
	return true
}

func testFloatLiteral(t *testing.T, fl ast.Expression, value float64) bool {
	float, ok := fl.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("fl not *ast.FloatLiteral. got=%T", fl)
		return false
	}
	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}
	return true
}

func testStringLiteral(t *testing.T, sl ast.Expression, value string) bool {
	str, ok := sl.(*ast.StringLiteral)
	if !ok {
		t.Errorf("sl not *ast.StringLiteral. got=%T", sl)
		return false
	}
	if str.Value != value {
		t.Errorf("str.Value not %q. got=%q", value, str.Value)
		return false
	}
	return true
}

func testCharLiteral(t *testing.T, cl ast.Expression, value byte) bool {
	char, ok := cl.(*ast.CharLiteral)
	if !ok {
		t.Errorf("cl not *ast.CharLiteral. got=%T", cl)
		return false
	}
	if char.Value != value {
		t.Errorf("char.Value not %q. got=%q", value, char.Value)
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, bl ast.Expression, value bool) bool {
	bool, ok := bl.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("bl not *ast.BooleanLiteral. got=%T", bl)
		return false
	}
	if bool.Value != value {
		t.Errorf("bool.Value not %t. got=%t", value, bool.Value)
		return false
	}
	return true
}

func testNullLiteral(t *testing.T, nl ast.Expression) bool {
	_, ok := nl.(*ast.NullLiteral)
	if !ok {
		t.Errorf("nl not *ast.NullLiteral. got=%T", nl)
		return false
	}
	return true
}
