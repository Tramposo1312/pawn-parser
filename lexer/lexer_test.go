package lexer

import (
	"testing"

	"github.com/Tramposo1312/pawn-parser/token"
)

func TestNextToken(t *testing.T) {
	input := `new x = 5 + 10;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEW, "new"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.ADD, "+"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}