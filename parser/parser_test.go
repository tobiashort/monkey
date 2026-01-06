package parser_test

import (
	"reflect"
	"testing"

	"github.com/tobiashort/utils-go/strings"

	"github.com/tobiashort/monkey/ast"
	"github.com/tobiashort/monkey/lexer"
	"github.com/tobiashort/monkey/parser"
	"github.com/tobiashort/monkey/token"
)

func TestParse(t *testing.T) {
	input := `foobar;`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Expression: ast.IdentifierExpression{
				Identifier: token.Token{
					Type:    token.IDENT,
					Literal: "foobar",
					Line:    1,
					Column:  1,
				},
			},
		},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	p := parser.New(tokens)
	actualAst, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expectedAst, actualAst) {
		t.Fatalf(
			strings.Dedent(`
				           |Expected:
				           |%+v
				           |
				           |Got:
				           |%+v`),
			expectedAst,
			actualAst)
	}
}

func TestParse2(t *testing.T) {
	input := `1 + 2;`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Expression: ast.BinaryExpression{
				Left: ast.LiteralExpression{
					Literal: token.Token{
						Type:    token.INT,
						Literal: "1",
						Line:    1,
						Column:  1,
					},
				},
				Operator: token.Token{
					Type:    token.PLUS,
					Literal: "+",
					Line:    1,
					Column:  3,
				},
				Right: ast.LiteralExpression{
					Literal: token.Token{
						Type:    token.INT,
						Literal: "2",
						Line:    1,
						Column:  5,
					},
				},
			},
		},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	p := parser.New(tokens)
	actualAst, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expectedAst, actualAst) {
		t.Fatalf(
			strings.Dedent(`
				           |Expected:
				           |%+v
				           |
				           |Got:
				           |%+v`),
			expectedAst,
			actualAst)
	}
}

func TestParse3(t *testing.T) {
	input := `1 + 2 * 3;`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Expression: ast.BinaryExpression{
				Left: ast.LiteralExpression{
					Literal: token.Token{
						Type:    token.INT,
						Literal: "1",
						Line:    1,
						Column:  1,
					},
				},
				Operator: token.Token{
					Type:    token.PLUS,
					Literal: "+",
					Line:    1,
					Column:  3,
				},
				Right: ast.BinaryExpression{
					Left: ast.LiteralExpression{
						Literal: token.Token{
							Type:    token.INT,
							Literal: "2",
							Line:    1,
							Column:  5,
						},
					},
					Operator: token.Token{
						Type:    token.ASTERISK,
						Literal: "*",
						Line:    1,
						Column:  7,
					},
					Right: ast.LiteralExpression{
						Literal: token.Token{
							Type:    token.INT,
							Literal: "3",
							Line:    1,
							Column:  9,
						},
					},
				},
			},
		},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	p := parser.New(tokens)
	actualAst, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expectedAst, actualAst) {
		t.Fatalf(
			strings.Dedent(`
				           |Expected:
				           |%+v
				           |
				           |Got:
				           |%+v`),
			expectedAst,
			actualAst)
	}
}

func TestParse4(t *testing.T) {
	input := `1 * 2 + 3;`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Expression: ast.BinaryExpression{
				Left: ast.BinaryExpression{
					Left: ast.LiteralExpression{
						Literal: token.Token{
							Type:    token.INT,
							Literal: "1",
							Line:    1,
							Column:  1,
						},
					},
					Operator: token.Token{
						Type:    token.ASTERISK,
						Literal: "*",
						Line:    1,
						Column:  3,
					},
					Right: ast.LiteralExpression{
						Literal: token.Token{
							Type:    token.INT,
							Literal: "2",
							Line:    1,
							Column:  5,
						},
					},
				},
				Operator: token.Token{
					Type:    token.PLUS,
					Literal: "+",
					Line:    1,
					Column:  7,
				},
				Right: ast.LiteralExpression{
					Literal: token.Token{
						Type:    token.INT,
						Literal: "3",
						Line:    1,
						Column:  9,
					},
				},
			},
		},
	}

	l := lexer.New("", input)
	tokens, err := l.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	p := parser.New(tokens)
	actualAst, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expectedAst, actualAst) {
		t.Fatalf(
			strings.Dedent(`
				           |Expected:
				           |%+v
				           |
				           |Got:
				           |%+v`),
			expectedAst,
			actualAst)
	}
}
