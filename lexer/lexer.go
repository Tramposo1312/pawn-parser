package lexer

import (
	"fmt"

	"github.com/Tramposo1312/pawn-parser/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	column       int
	errors       []string
}

func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
		errors: []string{},
	}
	l.readChar()
	return l
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.EQL)
		} else {
			tok = l.makeToken(token.ASSIGN)
		}
	case '+':
		tok = l.handlePlusOperator()
	case '-':
		tok = l.handleMinusOperator()
	case '!':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.NEQ)
		} else {
			tok = l.makeToken(token.NOT)
		}
	case '/':
		if l.peekChar() == '/' {
			tok.Type = token.COMMENT
			tok.Literal = l.readLineComment()
			return tok
		} else if l.peekChar() == '*' {
			tok.Type = token.COMMENT
			tok.Literal = l.readBlockComment()
			return tok
		} else if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.QUO_ASSIGN)
		} else {
			tok = l.makeToken(token.QUO)
		}
	case '*':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.MUL_ASSIGN)
		} else {
			tok = l.makeToken(token.MUL)
		}
	case '<':
		tok = l.handleLessThanOperator()
	case '>':
		tok = l.handleGreaterThanOperator()
	case ';':
		tok = l.makeToken(token.SEMICOLON)
	case ':':
		tok = l.makeToken(token.COLON)
	case ',':
		tok = l.makeToken(token.COMMA)
	case '.':
		tok = l.makeToken(token.PERIOD)
	case '(':
		tok = l.makeToken(token.LPAREN)
	case ')':
		tok = l.makeToken(token.RPAREN)
	case '{':
		tok = l.makeToken(token.LBRACE)
	case '}':
		tok = l.makeToken(token.RBRACE)
	case '[':
		tok = l.makeToken(token.LBRACK)
	case ']':
		tok = l.makeToken(token.RBRACK)
	case '&':
		tok = l.handleAmpersandOperator()
	case '|':
		tok = l.handlePipeOperator()
	case '^':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.XOR_ASSIGN)
		} else {
			tok = l.makeToken(token.XOR)
		}
	case '%':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.REM_ASSIGN)
		} else {
			tok = l.makeToken(token.REM)
		}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '\'':
		tok.Type = token.CHAR
		tok.Literal = l.readCharLiteral()
	case '#':
		return l.readPreprocessorDirective()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			tok = l.makeToken(token.ILLEGAL)
			l.errors = append(l.errors, fmt.Sprintf("Unexpected character: %c at line %d, column %d", l.ch, l.line, l.column))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) Errors() []string {
	return l.errors
}
func (l *Lexer) makeToken(tokenType token.TokenType) token.Token {
	return token.Token{Type: tokenType, Literal: string(l.ch), Line: l.line, Column: l.column}
}

func (l *Lexer) makeTwoCharToken(tokenType token.TokenType) token.Token {
	ch := l.ch
	startColumn := l.column
	l.readChar()
	literal := string(ch) + string(l.ch)
	return token.Token{Type: tokenType, Literal: literal, Line: l.line, Column: startColumn}
}

func (l *Lexer) handlePlusOperator() token.Token {
	if l.peekChar() == '=' {
		return l.makeTwoCharToken(token.ADD_ASSIGN)
	} else if l.peekChar() == '+' {
		return l.makeTwoCharToken(token.INC)
	}
	return l.makeToken(token.PLUS)
}

func (l *Lexer) handleMinusOperator() token.Token {
	if l.peekChar() == '=' {
		return l.makeTwoCharToken(token.SUB_ASSIGN)
	} else if l.peekChar() == '-' {
		return l.makeTwoCharToken(token.DEC)
	}
	return l.makeToken(token.MINUS)
}
func (l *Lexer) handleLessThanOperator() token.Token {
	if l.peekChar() == '=' {
		return l.makeTwoCharToken(token.LEQ)
	} else if l.peekChar() == '<' {
		l.readChar()
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(token.SHL_ASSIGN)
		}
		return token.Token{Type: token.SHL, Literal: "<<", Line: l.line, Column: l.column - 1}
	}
	return l.makeToken(token.LSS)
}

func (l *Lexer) handleGreaterThanOperator() token.Token {
	if l.peekChar() == '=' {
		return l.makeTwoCharToken(token.GEQ)
	} else if l.peekChar() == '>' {
		l.readChar()
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(token.SHR_ASSIGN)
		}
		return token.Token{Type: token.SHR, Literal: ">>", Line: l.line, Column: l.column - 1}
	}
	return l.makeToken(token.GTR)
}

func (l *Lexer) handleAmpersandOperator() token.Token {
	if l.peekChar() == '&' {
		return l.makeTwoCharToken(token.LAND)
	} else if l.peekChar() == '=' {
		return l.makeTwoCharToken(token.AND_ASSIGN)
	} else if l.peekChar() == '^' {
		l.readChar()
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(token.AND_NOT_ASSIGN)
		}
		return token.Token{Type: token.AND_NOT, Literal: "&^", Line: l.line, Column: l.column - 1}
	}
	return l.makeToken(token.AND)
}

func (l *Lexer) handlePipeOperator() token.Token {
	if l.peekChar() == '|' {
		return l.makeTwoCharToken(token.LOR)
	} else if l.peekChar() == '=' {
		return l.makeTwoCharToken(token.OR_ASSIGN)
	}
	return l.makeToken(token.OR)
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == ':' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() token.Token {
	startPosition := l.position
	startColumn := l.column
	isFloat := false

	if l.ch == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X') {
		// Hexadecimal
		l.readChar() // consume '0'
		l.readChar() // consume 'x' or 'X'
		for isHexDigit(l.ch) {
			l.readChar()
		}
	} else if l.ch == '0' && (l.peekChar() == 'b' || l.peekChar() == 'B') {
		// Binary
		l.readChar() // consume '0'
		l.readChar() // consume 'b' or 'B'
		for l.ch == '0' || l.ch == '1' {
			l.readChar()
		}
	} else {
		for isDigit(l.ch) {
			l.readChar()
		}
		if l.ch == '.' && isDigit(l.peekChar()) {
			isFloat = true
			l.readChar() // consume '.'
			for isDigit(l.ch) {
				l.readChar()
			}
		}
	}

	var tokenType token.TokenType
	if isFloat {
		tokenType = token.FLOAT
	} else {
		tokenType = token.INT
	}

	return token.Token{
		Type:    tokenType,
		Literal: l.input[startPosition:l.position],
		Line:    l.line,
		Column:  startColumn,
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1 // Start after the opening quote
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position] // Don't include the closing quote
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++
}

func (l *Lexer) readCharLiteral() string {
	position := l.position
	for {
		l.readChar()
		if l.ch == '\'' || l.ch == 0 {
			break
		}
	}
	l.readChar() // read closing quote
	return l.input[position:l.position]
}

func (l *Lexer) readLineComment() string {
	position := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readBlockComment() string {
	position := l.position
	for {
		if l.ch == 0 {
			break
		}
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar()
			l.readChar()
			break
		}
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readPreprocessorDirective() token.Token {
	startPosition := l.position
	l.readChar() // consume '#'

	directive := l.readIdentifier()

	switch directive {
	case "include":
		return token.Token{Type: token.INCLUDE, Literal: "#include", Line: l.line, Column: l.column - (l.position - startPosition)}
	case "define":
		return token.Token{Type: token.DEFINE, Literal: "#define", Line: l.line, Column: l.column - (l.position - startPosition)}
	case "ifdef":
		return token.Token{Type: token.IFDEF, Literal: "#ifdef", Line: l.line, Column: l.column - (l.position - startPosition)}
	case "endif":
		return token.Token{Type: token.ENDIF, Literal: "#endif", Line: l.line, Column: l.column - (l.position - startPosition)}
	default:
		return token.Token{Type: token.DIRECTIVE, Literal: "#" + directive, Line: l.line, Column: l.column - (l.position - startPosition)}
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isHexDigit(ch byte) bool {
	return isDigit(ch) || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}
