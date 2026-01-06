package lexer_test

import (
	"reflect"
	"testing"

	"github.com/tobiashort/monkey/lexer"
	"github.com/tobiashort/monkey/token"

	"github.com/tobiashort/utils-go/strings"
)

func TestAnalyze(t *testing.T) {
	input := `=+(){},;`

	expectedTokens := []token.Token{
		{Type: token.ASSIGN, Literal: "=", File: "", Line: 1, Column: 1},
		{Type: token.PLUS, Literal: "+", File: "", Line: 1, Column: 2},
		{Type: token.LPAREN, Literal: "(", File: "", Line: 1, Column: 3},
		{Type: token.RPAREN, Literal: ")", File: "", Line: 1, Column: 4},
		{Type: token.LBRACE, Literal: "{", File: "", Line: 1, Column: 5},
		{Type: token.RBRACE, Literal: "}", File: "", Line: 1, Column: 6},
		{Type: token.COMMA, Literal: ",", File: "", Line: 1, Column: 7},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 1, Column: 8},
		{Type: token.EOF, Literal: "", File: "", Line: 1, Column: 9},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i, expectedToken := range expectedTokens {
		actualToken := tokens[i]
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

func TestAnalyze2(t *testing.T) {
	input := strings.Dedent(`let five = 5;
                            |let ten = 10;
                            |let add = fn(x, y) {
                            |    x + y;
                            |};
                            |let result = add(five, ten);
                            |`)

	expectedTokens := []token.Token{
		{Type: token.LET, Literal: "let", File: "", Line: 1, Column: 1},
		{Type: token.IDENT, Literal: "five", File: "", Line: 1, Column: 5},
		{Type: token.ASSIGN, Literal: "=", File: "", Line: 1, Column: 10},
		{Type: token.INT, Literal: "5", File: "", Line: 1, Column: 12},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 1, Column: 13},
		{Type: token.LET, Literal: "let", File: "", Line: 2, Column: 1},
		{Type: token.IDENT, Literal: "ten", File: "", Line: 2, Column: 5},
		{Type: token.ASSIGN, Literal: "=", File: "", Line: 2, Column: 9},
		{Type: token.INT, Literal: "10", File: "", Line: 2, Column: 11},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 2, Column: 13},
		{Type: token.LET, Literal: "let", File: "", Line: 3, Column: 1},
		{Type: token.IDENT, Literal: "add", File: "", Line: 3, Column: 5},
		{Type: token.ASSIGN, Literal: "=", File: "", Line: 3, Column: 9},
		{Type: token.FUNCTION, Literal: "fn", File: "", Line: 3, Column: 11},
		{Type: token.LPAREN, Literal: "(", File: "", Line: 3, Column: 13},
		{Type: token.IDENT, Literal: "x", File: "", Line: 3, Column: 14},
		{Type: token.COMMA, Literal: ",", File: "", Line: 3, Column: 15},
		{Type: token.IDENT, Literal: "y", File: "", Line: 3, Column: 17},
		{Type: token.RPAREN, Literal: ")", File: "", Line: 3, Column: 18},
		{Type: token.LBRACE, Literal: "{", File: "", Line: 3, Column: 20},
		{Type: token.IDENT, Literal: "x", File: "", Line: 4, Column: 5},
		{Type: token.PLUS, Literal: "+", File: "", Line: 4, Column: 7},
		{Type: token.IDENT, Literal: "y", File: "", Line: 4, Column: 9},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 4, Column: 10},
		{Type: token.RBRACE, Literal: "}", File: "", Line: 5, Column: 1},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 5, Column: 2},
		{Type: token.LET, Literal: "let", File: "", Line: 6, Column: 1},
		{Type: token.IDENT, Literal: "result", File: "", Line: 6, Column: 5},
		{Type: token.ASSIGN, Literal: "=", File: "", Line: 6, Column: 12},
		{Type: token.IDENT, Literal: "add", File: "", Line: 6, Column: 14},
		{Type: token.LPAREN, Literal: "(", File: "", Line: 6, Column: 17},
		{Type: token.IDENT, Literal: "five", File: "", Line: 6, Column: 18},
		{Type: token.COMMA, Literal: ",", File: "", Line: 6, Column: 22},
		{Type: token.IDENT, Literal: "ten", File: "", Line: 6, Column: 24},
		{Type: token.RPAREN, Literal: ")", File: "", Line: 6, Column: 27},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 6, Column: 28},
		{Type: token.EOF, Literal: "", File: "", Line: 7, Column: 1},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i, expectedToken := range expectedTokens {
		actualToken := tokens[i]
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

func TestAnalyze3(t *testing.T) {
	input := `==!!=-/*<><=>=`

	expectedTokens := []token.Token{
		{Type: token.EQUAL, Literal: "==", File: "", Line: 1, Column: 1},
		{Type: token.BANG, Literal: "!", File: "", Line: 1, Column: 3},
		{Type: token.NOT_EQUAL, Literal: "!=", File: "", Line: 1, Column: 4},
		{Type: token.MINUS, Literal: "-", File: "", Line: 1, Column: 6},
		{Type: token.SLASH, Literal: "/", File: "", Line: 1, Column: 7},
		{Type: token.ASTERISK, Literal: "*", File: "", Line: 1, Column: 8},
		{Type: token.LT, Literal: "<", File: "", Line: 1, Column: 9},
		{Type: token.GT, Literal: ">", File: "", Line: 1, Column: 10},
		{Type: token.LEQT, Literal: "<=", File: "", Line: 1, Column: 11},
		{Type: token.GEQT, Literal: ">=", File: "", Line: 1, Column: 13},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i, expectedToken := range expectedTokens {
		actualToken := tokens[i]
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

func TestAnalyze4(t *testing.T) {
	input :=
		strings.Dedent(
			`if (5 < 10) {
            |    return true;
            |} else {
            |    return false;
            |}
            |`)

	expectedTokens := []token.Token{
		{Type: token.IF, Literal: "if", File: "", Line: 1, Column: 1},
		{Type: token.LPAREN, Literal: "(", File: "", Line: 1, Column: 4},
		{Type: token.INT, Literal: "5", File: "", Line: 1, Column: 5},
		{Type: token.LT, Literal: "<", File: "", Line: 1, Column: 7},
		{Type: token.INT, Literal: "10", File: "", Line: 1, Column: 9},
		{Type: token.RPAREN, Literal: ")", File: "", Line: 1, Column: 11},
		{Type: token.LBRACE, Literal: "{", File: "", Line: 1, Column: 13},
		{Type: token.RETURN, Literal: "return", File: "", Line: 2, Column: 5},
		{Type: token.TRUE, Literal: "true", File: "", Line: 2, Column: 12},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 2, Column: 16},
		{Type: token.RBRACE, Literal: "}", File: "", Line: 3, Column: 1},
		{Type: token.ELSE, Literal: "else", File: "", Line: 3, Column: 3},
		{Type: token.LBRACE, Literal: "{", File: "", Line: 3, Column: 8},
		{Type: token.RETURN, Literal: "return", File: "", Line: 4, Column: 5},
		{Type: token.FALSE, Literal: "false", File: "", Line: 4, Column: 12},
		{Type: token.SEMICOLON, Literal: ";", File: "", Line: 4, Column: 17},
		{Type: token.RBRACE, Literal: "}", File: "", Line: 5, Column: 1},
		{Type: token.EOF, Literal: "", File: "", Line: 6, Column: 1},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i, expectedToken := range expectedTokens {
		actualToken := tokens[i]
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

func TestAnalyze5(t *testing.T) {
	input :=
		strings.Dedent(
			`"foo" 0
			|12.3`)

	expectedTokens := []token.Token{
		{Type: token.STRING, Literal: "\"foo\"", File: "", Line: 1, Column: 1},
		{Type: token.INT, Literal: "0", File: "", Line: 1, Column: 7},
		{Type: token.FLOAT, Literal: "12.3", File: "", Line: 2, Column: 1},
		{Type: token.EOF, Literal: "", File: "", Line: 2, Column: 5},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i, expectedToken := range expectedTokens {
		actualToken := tokens[i]
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
