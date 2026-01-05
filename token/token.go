package token

import (
	"fmt"
	"runtime"
)

type TokenType = string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	EQUAL     = "=="
	NOT_EQUAL = "!="
	BANG      = "!"
	SLASH     = "/"
	ASTERISK  = "*"
	LT        = "<"
	GT        = ">"
	LEQT      = "<="
	GEQT      = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

func BindingPower(t Token) (int, error) {
	switch t.Type {
	case SEMICOLON:
		return 0, nil
	case EQUAL, NOT_EQUAL:
		return 1, nil
	case LT, GT, LEQT, GEQT:
		return 2, nil
	case PLUS, MINUS:
		return 3, nil
	case ASTERISK, SLASH:
		return 4, nil
	case BANG:
		return 5, nil
	default:
		_, file, line, _ := runtime.Caller(0)
		return -1, fmt.Errorf("%s:%d: illegal token type %q", file, line, t.Type)
	}
}
