// token/token.go

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"
	COMMENT           = "COMMENT"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"
	FLOAT  = "FLOAT"
	CHAR   = "CHAR" // 'a'
	STRING = "STRING"

	// Operators
	PLUS           = "+"
	MINUS          = "-"
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
	INC            = "++"
	DEC            = "--"
	EQL            = "=="
	LSS            = "<"
	GTR            = ">"
	ASSIGN         = "="
	NOT            = "!"
	ELLIPSIS       = "..."
	TILDE          = "~"

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
	FUNCTION = "function"
	BREAK    = "break"
	CASE     = "case"
	CONST    = "const"
	CONTINUE = "continue"
	DEFAULT  = "default"
	DO       = "do"
	ELSE     = "else"
	ENUM     = "enum"
	FOR      = "for"
	GOTO     = "goto"
	IF       = "if"
	NEW      = "new"
	RETURN   = "return"
	SIZEOF   = "sizeof"
	STATIC   = "static"
	SWITCH   = "switch"
	WHILE    = "while"

	// Pawn-specific keywords
	ASSERT     = "assert"
	DEFINED    = "defined"
	FORWARD    = "forward"
	NATIVE     = "native"
	OPERATOR   = "operator"
	STRUCT     = "struct"
	TAG        = "tag"
	PUBLIC     = "public"
	STOCK      = "stock"
	TAGOF      = "tagof"
	CHAR_      = "char"
	FLOAT_     = "float"
	BOOL       = "bool"
	VOID       = "void"
	TRUE       = "true"
	FALSE      = "false"
	NULL       = "null"
	FOREACH    = "foreach"
	SLEEP      = "sleep"
	STATE      = "state"
	EXIT       = "exit"
	TIMER      = "timer"
	ITERFUNC   = "iterfunc"
	HOOK       = "hook"
	INLINE     = "inline"
	MASTER     = "master"
	TASK       = "task"
	PTASK      = "ptask"
	FOREIGN    = "foreign"
	GLOBAL     = "global"
	REMOTEFUNC = "remotefunc"
	USING      = "using"
	YIELD      = "yield"
	LOADTEXT   = "loadtext"

	// Preprocessor directives
	DIRECTIVE = "#"

	// Comparison
	EQ  = "=="
	NEQ = "!="
	LT  = "<"
	GT  = ">"
	LEQ = "<="
	GEQ = ">="
)

var keywords = map[string]TokenType{
	"break":      BREAK,
	"case":       CASE,
	"const":      CONST,
	"continue":   CONTINUE,
	"default":    DEFAULT,
	"do":         DO,
	"else":       ELSE,
	"enum":       ENUM,
	"for":        FOR,
	"goto":       GOTO,
	"if":         IF,
	"new":        NEW,
	"return":     RETURN,
	"sizeof":     SIZEOF,
	"static":     STATIC,
	"switch":     SWITCH,
	"while":      WHILE,
	"assert":     ASSERT,
	"defined":    DEFINED,
	"forward":    FORWARD,
	"native":     NATIVE,
	"operator":   OPERATOR,
	"public":     PUBLIC,
	"stock":      STOCK,
	"tagof":      TAGOF,
	"char":       CHAR_,
	"float":      FLOAT_,
	"bool":       BOOL,
	"void":       VOID,
	"true":       TRUE,
	"false":      FALSE,
	"null":       NULL,
	"foreach":    FOREACH,
	"sleep":      SLEEP,
	"state":      STATE,
	"exit":       EXIT,
	"timer":      TIMER,
	"iterfunc":   ITERFUNC,
	"hook":       HOOK,
	"inline":     INLINE,
	"master":     MASTER,
	"task":       TASK,
	"ptask":      PTASK,
	"foreign":    FOREIGN,
	"global":     GLOBAL,
	"remotefunc": REMOTEFUNC,
	"using":      USING,
	"yield":      YIELD,
	"loadtext":   LOADTEXT,
	"function":   FUNCTION,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
