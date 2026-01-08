package token

import "github.com/tobiashort/utils-go/errors"

type TokenType = string

type Token struct {
	Type    TokenType
	Literal string
	File    string
	Line    int
	Column  int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

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
	BAND      = "&"
	LAND      = "&&"
	BOR       = "|"
	LOR       = "||"

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
	case SEMICOLON, RPAREN, LBRACE:
		return 0, nil
	case LOR:
		return 1, nil
	case LAND:
		return 2, nil
	case BOR:
		return 3, nil
	case BAND:
		return 4, nil
	case EQUAL, NOT_EQUAL:
		return 5, nil
	case LT, GT, LEQT, GEQT:
		return 6, nil
	case PLUS, MINUS:
		return 7, nil
	case ASTERISK, SLASH:
		return 8, nil
	case BANG:
		return 9, nil
	default:
		return -1, errors.WithCtxf("%s:%d:%d: illegal token type %q", t.File, t.Line, t.Column, t.Type)
	}
}
