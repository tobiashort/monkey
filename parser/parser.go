package parser

import (
	"fmt"
	"runtime"

	"github.com/tobiashort/monkey/ast"
	"github.com/tobiashort/monkey/lexer"
	"github.com/tobiashort/monkey/token"
)

type Parser struct {
	filename string
	lexer    *lexer.Lexer
	tokens   []token.Token
	position int
	ast      ast.Ast
}

func New(filename string, input string) *Parser {
	return &Parser{
		filename: filename,
		lexer:    lexer.New(filename, input),
		tokens:   make([]token.Token, 0),
		position: 0,
		ast:      make(ast.Ast, 0),
	}
}

func (p *Parser) Parse() (ast.Ast, error) {
	tokens, err := p.lexer.Analyze()
	if err != nil {
		return p.ast, err
	}
	p.tokens = tokens

	for p.token().Type != token.EOF {
		switch p.token().Type {
		case token.INT:
			fallthrough
		case token.IDENT:
			if err := p.parseExpressionStatement(); err != nil {
				return p.ast, err
			}
		default:
			_, file, line, _ := runtime.Caller(0)
			return p.ast, fmt.Errorf("%s:%d: illegal token type %q", file, line, p.token().Type)
		}
	}
	return p.ast, nil
}

func (p *Parser) parseExpressionStatement() error {
	node, err := p.parseExpression(0)
	if err != nil {
		return err
	}
	p.nextToken()
	if err := p.expect(token.SEMICOLON); err != nil {
		return err
	}
	p.nextToken()
	node = ast.ExpressionStatement{Expression: node}
	p.ast = append(p.ast, node)
	return nil
}

func (p *Parser) parseExpression(bindingPower int) (ast.Node, error) {
	var left ast.Node
	switch p.token().Type {
	case token.IDENT:
		left = ast.IdentifierExpression{
			Identifier: p.token(),
		}
	case token.INT:
		left = ast.LiteralExpression{
			Literal: p.token(),
		}
	default:
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("%s:%d: illegal token type %q", file, line, p.token().Type)
	}

	for {
		nextBindingPower, err := token.BindingPower(p.peekToken())
		if err != nil {
			return nil, err
		}
		if nextBindingPower <= bindingPower {
			break
		}
		operator := p.nextToken()
		p.nextToken()
		right, err := p.parseExpression(nextBindingPower)
		if err != nil {
			return nil, err
		}
		left = ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left, nil
}

func (p *Parser) expect(tokenType token.TokenType) error {
	t := p.token()
	if t.Type != tokenType {
		return fmt.Errorf("parse error %s %d:%d: got %q, expected %q", p.filename, t.Line, t.Column, t.Type, tokenType)
	}
	return nil
}

func (p *Parser) token() token.Token {
	return p.tokens[p.position]
}

func (p *Parser) peekToken() token.Token {
	return p.tokens[p.position+1]
}

func (p *Parser) nextToken() token.Token {
	p.position++
	return p.tokens[p.position]
}
