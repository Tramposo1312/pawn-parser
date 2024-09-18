package lexer

import (
	"testing"

	"github.com/Tramposo1312/pawn-parser/token"
)

func TestNextToken(t *testing.T) {
	input := `
#include <a_samp>
#define MAX_PLAYERS 50

native SetPlayerPos(playerid, Float:x, Float:y, Float:z);

public OnGameModeInit()
{
    // This is a line comment
    print("Game mode initialized!");
    return 1;
}

/* This is a
   multi-line comment */

stock Float:GetDistance(Float:x1, Float:y1, Float:z1, Float:x2, Float:y2, Float:z2)
{
    return floatsqroot((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1));
}

#ifdef DEBUG
    printf("Debug mode enabled");
#endif
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INCLUDE, "#include"},
		{token.LT, "<"},
		{token.IDENT, "a_samp"},
		{token.GT, ">"},
		{token.DEFINE, "#define"},
		{token.IDENT, "MAX_PLAYERS"},
		{token.INT, "50"},
		{token.NATIVE, "native"},
		{token.IDENT, "SetPlayerPos"},
		{token.LPAREN, "("},
		{token.IDENT, "playerid"},
		{token.COMMA, ","},
		{token.IDENT, "Float:x"},
		{token.COMMA, ","},
		{token.IDENT, "Float:y"},
		{token.COMMA, ","},
		{token.IDENT, "Float:z"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.PUBLIC, "public"},
		{token.IDENT, "OnGameModeInit"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.COMMENT, "// This is a line comment"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.STRING, "Game mode initialized!"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.COMMENT, "/* This is a\n   multi-line comment */"},
		{token.STOCK, "stock"},
		{token.IDENT, "Float:GetDistance"},
		{token.LPAREN, "("},
		{token.IDENT, "Float:x1"},
		{token.COMMA, ","},
		{token.IDENT, "Float:y1"},
		{token.COMMA, ","},
		{token.IDENT, "Float:z1"},
		{token.COMMA, ","},
		{token.IDENT, "Float:x2"},
		{token.COMMA, ","},
		{token.IDENT, "Float:y2"},
		{token.COMMA, ","},
		{token.IDENT, "Float:z2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "floatsqroot"},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.IDENT, "x2"},
		{token.MINUS, "-"},
		{token.IDENT, "x1"},
		{token.RPAREN, ")"},
		{token.MUL, "*"},
		{token.LPAREN, "("},
		{token.IDENT, "x2"},
		{token.MINUS, "-"},
		{token.IDENT, "x1"},
		{token.RPAREN, ")"},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.IDENT, "y2"},
		{token.MINUS, "-"},
		{token.IDENT, "y1"},
		{token.RPAREN, ")"},
		{token.MUL, "*"},
		{token.LPAREN, "("},
		{token.IDENT, "y2"},
		{token.MINUS, "-"},
		{token.IDENT, "y1"},
		{token.RPAREN, ")"},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.IDENT, "z2"},
		{token.MINUS, "-"},
		{token.IDENT, "z1"},
		{token.RPAREN, ")"},
		{token.MUL, "*"},
		{token.LPAREN, "("},
		{token.IDENT, "z2"},
		{token.MINUS, "-"},
		{token.IDENT, "z1"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IFDEF, "#ifdef"},
		{token.IDENT, "DEBUG"},
		{token.IDENT, "printf"},
		{token.LPAREN, "("},
		{token.STRING, "Debug mode enabled"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.ENDIF, "#endif"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestPawnSpecificTokens(t *testing.T) {
	input := `
#define MAX_PLAYERS 100
#if defined SOME_CONSTANT
new Float:myVariable = 10.5;
@emit inc eax
#endif
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEFINE, "#define"},
		{token.IDENT, "MAX_PLAYERS"},
		{token.INT, "100"},
		{token.DIRECTIVE, "#if"},
		{token.DEFINED, "defined"},
		{token.IDENT, "SOME_CONSTANT"},
		{token.NEW, "new"},
		{token.IDENT, "Float:myVariable"},
		{token.ASSIGN, "="},
		{token.FLOAT, "10.5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "@emit"},
		{token.IDENT, "inc"},
		{token.IDENT, "eax"},
		{token.ENDIF, "#endif"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
