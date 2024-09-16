// token/token.go

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"
	COMMENT           = "COMMENT"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	CHAR   = "CHAR"
	STRING = "STRING"

	// Operators
	ADD            = "+"
	SUB            = "-"
	MUL            = "*"
	QUO            = "/"
	REM            = "%"
	AND            = "&"
	OR             = "|"
	XOR            = "^"
	SHL            = "<<"
	SHR            = ">>"
	AND_NOT        = "&^"
	ADD_ASSIGN     = "+="
	SUB_ASSIGN     = "-="
	MUL_ASSIGN     = "*="
	QUO_ASSIGN     = "/="
	REM_ASSIGN     = "%="
	AND_ASSIGN     = "&="
	OR_ASSIGN      = "|="
	XOR_ASSIGN     = "^="
	SHL_ASSIGN     = "<<="
	SHR_ASSIGN     = ">>="
	AND_NOT_ASSIGN = "&^="
	LAND           = "&&"
	LOR            = "||"
	ARROW          = "<-"
	INC            = "++"
	DEC            = "--"
	EQL            = "=="
	LSS            = "<"
	GTR            = ">"
	ASSIGN         = "="
	NOT            = "!"
	NEQ            = "!="
	LEQ            = "<="
	GEQ            = ">="

	// Delimiters
	LPAREN    = "("
	LBRACK    = "["
	LBRACE    = "{"
	COMMA     = ","
	PERIOD    = "."
	RPAREN    = ")"
	RBRACK    = "]"
	RBRACE    = "}"
	SEMICOLON = ";"
	COLON     = ":"

	// Keywords
	BREAK    = "BREAK"
	CASE     = "CASE"
	CONST    = "CONST"
	CONTINUE = "CONTINUE"
	DEFAULT  = "DEFAULT"
	DO       = "DO"
	ELSE     = "ELSE"
	ENUM     = "ENUM"
	FOR      = "FOR"
	GOTO     = "GOTO"
	IF       = "IF"
	NEW      = "NEW"
	RETURN   = "RETURN"
	SIZEOF   = "SIZEOF"
	STATIC   = "STATIC"
	SWITCH   = "SWITCH"
	WHILE    = "WHILE"

	// Pawn keywords
	ASSERT   = "ASSERT"
	DEFINED  = "DEFINED"
	FORWARD  = "FORWARD"
	NATIVE   = "NATIVE"
	OPERATOR = "OPERATOR"
	PUBLIC   = "PUBLIC"
	STOCK    = "STOCK"
	TAGOF    = "TAGOF"

	// Preprocessor directives
	DIRECTIVE = "#"
)

var keywords = map[string]TokenType{
	"break":    BREAK,
	"case":     CASE,
	"const":    CONST,
	"continue": CONTINUE,
	"default":  DEFAULT,
	"do":       DO,
	"else":     ELSE,
	"enum":     ENUM,
	"for":      FOR,
	"goto":     GOTO,
	"if":       IF,
	"new":      NEW,
	"return":   RETURN,
	"sizeof":   SIZEOF,
	"static":   STATIC,
	"switch":   SWITCH,
	"while":    WHILE,
	"assert":   ASSERT,
	"defined":  DEFINED,
	"forward":  FORWARD,
	"native":   NATIVE,
	"operator": OPERATOR,
	"public":   PUBLIC,
	"stock":    STOCK,
	"tagof":    TAGOF,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
