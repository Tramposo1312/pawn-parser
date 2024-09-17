// precedence/precedence.go

package precedence

import "github.com/Tramposo1312/pawn-parser/token"

const (
	LOWEST int = iota
	ASSIGN
	TERNARY
	LOGICAL_OR
	LOGICAL_AND
	BIT_OR
	BIT_XOR
	BIT_AND
	EQUALS
	LESSGREATER
	SHIFT
	SUM
	PRODUCT
	PREFIX
	POSTFIX
	CALL
	INDEX
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:   ASSIGN,
	token.QUESTION: TERNARY,
	token.LOR:      LOGICAL_OR,
	token.LAND:     LOGICAL_AND,
	token.OR:       BIT_OR,
	token.XOR:      BIT_XOR,
	token.AND:      BIT_AND,
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.LTE:      LESSGREATER,
	token.GTE:      LESSGREATER,
	token.SHL:      SHIFT,
	token.SHR:      SHIFT,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.MULTIPLY: PRODUCT,
	token.DIVIDE:   PRODUCT,
	token.MODULO:   PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACK:   INDEX,
	token.INC:      POSTFIX,
	token.DEC:      POSTFIX,
}

func GetPrecedence(tokenType token.TokenType) int {
	if p, ok := precedences[tokenType]; ok {
		return p
	}
	return LOWEST
}

func GetPrecedenceFromString(operator string) int {
	var tokenType token.TokenType
	switch operator {
	case "=":
		tokenType = token.ASSIGN
	case "||":
		tokenType = token.LOR
	case "&&":
		tokenType = token.LAND
	case "|":
		tokenType = token.OR
	case "^":
		tokenType = token.XOR
	case "&":
		tokenType = token.AND
	case "==", "!=":
		tokenType = token.EQ
	case "<", ">", "<=", ">=":
		tokenType = token.LT
	case "<<", ">>":
		tokenType = token.SHL
	case "+", "-":
		tokenType = token.PLUS
	case "*", "/", "%":
		tokenType = token.MULTIPLY
	default:
		return LOWEST
	}
	return GetPrecedence(tokenType)
}
