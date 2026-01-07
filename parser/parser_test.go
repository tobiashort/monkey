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

func test(t *testing.T, input string, expectedAst ast.Ast) {
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

func TestParse(t *testing.T) {
	input := strings.Dedent(`foobar;
							|"foobar";
							|42;
							|42.0;`)

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.IdentifierExpression{
				Type: ast.IDENT,
				Identifier: token.Token{
					Type:    token.IDENT,
					Literal: "foobar",
					Line:    1,
					Column:  1,
				},
			},
		},
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.LiteralExpression{
				Type: ast.LITERAL,
				Literal: token.Token{
					Type:    token.STRING,
					Literal: "\"foobar\"",
					Line:    2,
					Column:  1,
				},
			},
		},
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.LiteralExpression{
				Type: ast.LITERAL,
				Literal: token.Token{
					Type:    token.INT,
					Literal: "42",
					Line:    3,
					Column:  1,
				},
			},
		},
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.LiteralExpression{
				Type: ast.LITERAL,
				Literal: token.Token{
					Type:    token.FLOAT,
					Literal: "42.0",
					Line:    4,
					Column:  1,
				},
			},
		},
	}

	test(t, input, expectedAst)
}

func TestParse2(t *testing.T) {
	input := `1 + 2;`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.BinaryExpression{
				Type: ast.BINARY,
				Left: ast.LiteralExpression{
					Type: ast.LITERAL,
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
					Type: ast.LITERAL,
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

	test(t, input, expectedAst)
}

func TestParse3(t *testing.T) {
	input := `1 + 2 * 3;`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.BinaryExpression{
				Type: ast.BINARY,
				Left: ast.LiteralExpression{
					Type: ast.LITERAL,
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
					Type: ast.BINARY,
					Left: ast.LiteralExpression{
						Type: ast.LITERAL,
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
						Type: ast.LITERAL,
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

	test(t, input, expectedAst)
}

func TestParse4(t *testing.T) {
	input := `1 * 2 + 3;`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.BinaryExpression{
				Type: ast.BINARY,
				Left: ast.BinaryExpression{
					Type: ast.BINARY,
					Left: ast.LiteralExpression{
						Type: ast.LITERAL,
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
						Type: ast.LITERAL,
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
					Type: ast.LITERAL,
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

	test(t, input, expectedAst)
}

func TestParse5(t *testing.T) {
	input := `(1 * (2 + 3));`

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.BinaryExpression{
				Type: ast.BINARY,
				Left: ast.LiteralExpression{
					Type: ast.LITERAL,
					Literal: token.Token{
						Type:    token.INT,
						Literal: "1",
						Line:    1,
						Column:  2,
					},
				},
				Operator: token.Token{
					Type:    token.ASTERISK,
					Literal: "*",
					Line:    1,
					Column:  4,
				},
				Right: ast.BinaryExpression{
					Type: ast.BINARY,
					Left: ast.LiteralExpression{
						Type: ast.LITERAL,
						Literal: token.Token{
							Type:    token.INT,
							Literal: "2",
							Line:    1,
							Column:  7,
						},
					},
					Operator: token.Token{
						Type:    token.PLUS,
						Literal: "+",
						Line:    1,
						Column:  9,
					},
					Right: ast.LiteralExpression{
						Type: ast.LITERAL,
						Literal: token.Token{
							Type:    token.INT,
							Literal: "3",
							Line:    1,
							Column:  11,
						},
					},
				},
			},
		},
	}

	test(t, input, expectedAst)
}

func TestParse6(t *testing.T) {
	input := "let a = 42;"

	expectedAst := ast.Ast{
		ast.LetStatement{
			Type: ast.LET,
			Identifier: token.Token{
				Type:    token.IDENT,
				Literal: "a",
				File:    "",
				Line:    1,
				Column:  5,
			},
			Expression: ast.LiteralExpression{
				Type: ast.LITERAL,
				Literal: token.Token{
					Type:    token.INT,
					Literal: "42",
					File:    "",
					Line:    1,
					Column:  9,
				},
			},
		},
	}

	test(t, input, expectedAst)
}

func TestParse7(t *testing.T) {
	input := "return a + b;"

	expectedAst := ast.Ast{
		ast.ReturnStatement{
			Type: ast.RETURN,
			Expression: ast.BinaryExpression{
				Type: ast.BINARY,
				Left: ast.IdentifierExpression{
					Type: ast.IDENT,
					Identifier: token.Token{
						Type:    token.IDENT,
						Literal: "a",
						File:    "",
						Line:    1,
						Column:  8,
					},
				},
				Operator: token.Token{
					Type:    token.PLUS,
					Literal: "+",
					File:    "",
					Line:    1,
					Column:  10,
				},
				Right: ast.IdentifierExpression{
					Type: ast.IDENT,
					Identifier: token.Token{
						Type:    token.IDENT,
						Literal: "b",
						File:    "",
						Line:    1,
						Column:  12,
					},
				},
			},
		},
	}

	test(t, input, expectedAst)
}

func TestParse8(t *testing.T) {
	input := strings.Dedent(`1 - -2;
		                    |foo && !bar;`)

	expectedAst := ast.Ast{
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.BinaryExpression{
				Type: ast.BINARY,
				Left: ast.LiteralExpression{
					Type: ast.LITERAL,
					Literal: token.Token{
						Type:    token.INT,
						Literal: "1",
						File:    "",
						Line:    1,
						Column:  1,
					},
				},
				Operator: token.Token{
					Type:    token.MINUS,
					Literal: "-",
					File:    "",
					Line:    1,
					Column:  3,
				},
				Right: ast.UnaryExpression{
					Type: ast.UNARY,
					Operator: token.Token{
						Type:    token.MINUS,
						Literal: "-",
						File:    "",
						Line:    1,
						Column:  5,
					},
					Right: ast.LiteralExpression{
						Type: ast.LITERAL,
						Literal: token.Token{
							Type:    token.INT,
							Literal: "2",
							File:    "",
							Line:    1,
							Column:  6,
						},
					},
				},
			},
		},
		ast.ExpressionStatement{
			Type: ast.EXPR,
			Expression: ast.BinaryExpression{
				Type: ast.BINARY,
				Left: ast.IdentifierExpression{
					Type: ast.IDENT,
					Identifier: token.Token{
						Type:    token.IDENT,
						Literal: "foo",
						File:    "",
						Line:    2,
						Column:  1,
					},
				},
				Operator: token.Token{
					Type:    token.LAND,
					Literal: "&&",
					File:    "",
					Line:    2,
					Column:  5,
				},
				Right: ast.UnaryExpression{
					Type: ast.UNARY,
					Operator: token.Token{
						Type:    token.BANG,
						Literal: "!",
						File:    "",
						Line:    2,
						Column:  8,
					},
					Right: ast.IdentifierExpression{
						Type: ast.IDENT,
						Identifier: token.Token{
							Type:    token.IDENT,
							Literal: "bar",
							File:    "",
							Line:    2,
							Column:  9,
						},
					},
				},
			},
		},
	}

	test(t, input, expectedAst)
}
