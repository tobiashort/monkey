package lexer_test

import (
	"reflect"
	"testing"

	"github.com/tobiashort/monkey/lexer"
	"github.com/tobiashort/monkey/token"

	"github.com/tobiashort/utils-go/strings"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	expectedTokens := []token.Token{
		{Type: token.ASSIGN, Literal: "=", Line: 1, Column: 1},
		{Type: token.PLUS, Literal: "+", Line: 1, Column: 2},
		{Type: token.LPAREN, Literal: "(", Line: 1, Column: 3},
		{Type: token.RPAREN, Literal: ")", Line: 1, Column: 4},
		{Type: token.LBRACE, Literal: "{", Line: 1, Column: 5},
		{Type: token.RBRACE, Literal: "}", Line: 1, Column: 6},
		{Type: token.COMMA, Literal: ",", Line: 1, Column: 7},
		{Type: token.SEMICOLON, Literal: ";", Line: 1, Column: 8},
		{Type: token.EOF, Literal: "", Line: 1, Column: 9},
	}

	l := lexer.New(input)

	for _, expectedToken := range expectedTokens {
		actualToken := l.NextToken()
		if !reflect.DeepEqual(expectedToken, actualToken) {
			t.Fatalf(
				strings.Dedent(`
				               |Expected: %+v
				               |Got:      %+v`),
				expectedToken,
				actualToken)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := strings.Dedent(`let five = 5;
                            |let ten = 10;
                            |let add = fn(x, y) {
                            |    x + y;
                            |};
                            |let result = add(five, ten);
                            |`)

	expectedTokens := []token.Token{
		{Type: token.LET, Literal: "let", Line: 1, Column: 1},
		{Type: token.IDENT, Literal: "five", Line: 1, Column: 5},
		{Type: token.ASSIGN, Literal: "=", Line: 1, Column: 10},
		{Type: token.INT, Literal: "5", Line: 1, Column: 12},
		{Type: token.SEMICOLON, Literal: ";", Line: 1, Column: 13},
		{Type: token.LET, Literal: "let", Line: 2, Column: 1},
		{Type: token.IDENT, Literal: "ten", Line: 2, Column: 5},
		{Type: token.ASSIGN, Literal: "=", Line: 2, Column: 9},
		{Type: token.INT, Literal: "10", Line: 2, Column: 11},
		{Type: token.SEMICOLON, Literal: ";", Line: 2, Column: 13},
		{Type: token.LET, Literal: "let", Line: 3, Column: 1},
		{Type: token.IDENT, Literal: "add", Line: 3, Column: 5},
		{Type: token.ASSIGN, Literal: "=", Line: 3, Column: 9},
		{Type: token.FUNCTION, Literal: "fn", Line: 3, Column: 11},
		{Type: token.LPAREN, Literal: "(", Line: 3, Column: 13},
		{Type: token.IDENT, Literal: "x", Line: 3, Column: 14},
		{Type: token.COMMA, Literal: ",", Line: 3, Column: 15},
		{Type: token.IDENT, Literal: "y", Line: 3, Column: 17},
		{Type: token.RPAREN, Literal: ")", Line: 3, Column: 18},
		{Type: token.LBRACE, Literal: "{", Line: 3, Column: 20},
		{Type: token.IDENT, Literal: "x", Line: 4, Column: 5},
		{Type: token.PLUS, Literal: "+", Line: 4, Column: 7},
		{Type: token.IDENT, Literal: "y", Line: 4, Column: 9},
		{Type: token.SEMICOLON, Literal: ";", Line: 4, Column: 10},
		{Type: token.RBRACE, Literal: "}", Line: 5, Column: 1},
		{Type: token.SEMICOLON, Literal: ";", Line: 5, Column: 2},
		{Type: token.LET, Literal: "let", Line: 6, Column: 1},
		{Type: token.IDENT, Literal: "result", Line: 6, Column: 5},
		{Type: token.ASSIGN, Literal: "=", Line: 6, Column: 12},
		{Type: token.IDENT, Literal: "add", Line: 6, Column: 14},
		{Type: token.LPAREN, Literal: "(", Line: 6, Column: 17},
		{Type: token.IDENT, Literal: "five", Line: 6, Column: 18},
		{Type: token.COMMA, Literal: ",", Line: 6, Column: 22},
		{Type: token.IDENT, Literal: "ten", Line: 6, Column: 24},
		{Type: token.RPAREN, Literal: ")", Line: 6, Column: 27},
		{Type: token.SEMICOLON, Literal: ";", Line: 6, Column: 28},
		{Type: token.EOF, Literal: "", Line: 7, Column: 1},
	}

	l := lexer.New(input)

	for _, expectedToken := range expectedTokens {
		actualToken := l.NextToken()
		if !reflect.DeepEqual(expectedToken, actualToken) {
			t.Fatalf(
				strings.Dedent(`
				               |Expected: %+v
				               |Got:      %+v`),
				expectedToken,
				actualToken)
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `==!!=-/*<><=>=`

	expectedTokens := []token.Token{
		{Type: token.EQUAL, Literal: "==", Line: 1, Column: 1},
		{Type: token.BANG, Literal: "!", Line: 1, Column: 3},
		{Type: token.NOT_EQUAL, Literal: "!=", Line: 1, Column: 4},
		{Type: token.MINUS, Literal: "-", Line: 1, Column: 6},
		{Type: token.SLASH, Literal: "/", Line: 1, Column: 7},
		{Type: token.ASTERISK, Literal: "*", Line: 1, Column: 8},
		{Type: token.LT, Literal: "<", Line: 1, Column: 9},
		{Type: token.GT, Literal: ">", Line: 1, Column: 10},
		{Type: token.LEQT, Literal: "<=", Line: 1, Column: 11},
		{Type: token.GEQT, Literal: ">=", Line: 1, Column: 13},
	}

	l := lexer.New(input)

	for _, expectedToken := range expectedTokens {
		actualToken := l.NextToken()
		if !reflect.DeepEqual(expectedToken, actualToken) {
			t.Fatalf(
				strings.Dedent(`
				               |Expected: %+v
				               |Got:      %+v`),
				expectedToken,
				actualToken)
		}
	}
}

func TestNextToken4(t *testing.T) {
	input :=
		strings.Dedent(
			`if (5 < 10) {
            |    return true;
            |} else {
            |    return false;
            |}
            |`)

	expectedTokens := []token.Token{
		{Type: token.IF, Literal: "if", Line: 1, Column: 1},
		{Type: token.LPAREN, Literal: "(", Line: 1, Column: 4},
		{Type: token.INT, Literal: "5", Line: 1, Column: 5},
		{Type: token.LT, Literal: "<", Line: 1, Column: 7},
		{Type: token.INT, Literal: "10", Line: 1, Column: 9},
		{Type: token.RPAREN, Literal: ")", Line: 1, Column: 11},
		{Type: token.LBRACE, Literal: "{", Line: 1, Column: 13},
		{Type: token.RETURN, Literal: "return", Line: 2, Column: 5},
		{Type: token.TRUE, Literal: "true", Line: 2, Column: 12},
		{Type: token.SEMICOLON, Literal: ";", Line: 2, Column: 16},
		{Type: token.RBRACE, Literal: "}", Line: 3, Column: 1},
		{Type: token.ELSE, Literal: "else", Line: 3, Column: 3},
		{Type: token.LBRACE, Literal: "{", Line: 3, Column: 8},
		{Type: token.RETURN, Literal: "return", Line: 4, Column: 5},
		{Type: token.FALSE, Literal: "false", Line: 4, Column: 12},
		{Type: token.SEMICOLON, Literal: ";", Line: 4, Column: 17},
		{Type: token.RBRACE, Literal: "}", Line: 5, Column: 1},
		{Type: token.EOF, Literal: "", Line: 6, Column: 1},
	}

	l := lexer.New(input)

	for _, expectedToken := range expectedTokens {
		actualToken := l.NextToken()
		if !reflect.DeepEqual(expectedToken, actualToken) {
			t.Fatalf(
				strings.Dedent(`
				               |Expected: %+v
				               |Got:      %+v`),
				expectedToken,
				actualToken)
		}
	}
}
