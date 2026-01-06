package parser

import (
	"fmt"

	"github.com/tobiashort/monkey/ast"
	"github.com/tobiashort/monkey/token"
)

type Parser struct {
	position int
	tokens   []token.Token
	ast      ast.Ast
}

func New(tokens []token.Token) *Parser {
	return &Parser{
		position: 0,
		tokens:   tokens,
		ast:      make(ast.Ast, 0),
	}
}

func (p *Parser) Parse() (ast.Ast, error) {
	for p.token().Type != token.EOF {
		switch p.token().Type {
		case token.LET:
			if err := p.parseLetStatement(); err != nil {
				return p.ast, err
			}
		case token.RETURN:
			if err := p.parseReturnStatement(); err != nil {
				return p.ast, err
			}
		case token.LPAREN:
			fallthrough
		case token.INT:
			fallthrough
		case token.FLOAT:
			fallthrough
		case token.STRING:
			fallthrough
		case token.IDENT:
			if err := p.parseExpressionStatement(); err != nil {
				return p.ast, err
			}
		default:
			return p.ast, fmt.Errorf("%s:%d:%d: illegal token type %q", p.token().File, p.token().Line, p.token().Column, p.token().Type)
		}
	}
	return p.ast, nil
}

func (p *Parser) parseLetStatement() error {
	if err := p.expect(token.LET); err != nil {
		return err
	}
	p.nextToken()
	node := ast.LetStatement{
		Type: ast.LET,
	}
	if err := p.expect(token.IDENT); err != nil {
		return err
	}
	node.Identifier = p.token()
	p.nextToken()
	p.expect(token.ASSIGN)
	p.nextToken()
	if expr, err := p.parseExpression(0); err != nil {
		return err
	} else {
		node.Expression = expr
	}
	p.ast = append(p.ast, node)
	p.nextToken()
	if err := p.expect(token.SEMICOLON); err != nil {
		return err
	}
	p.nextToken()
	return nil
}

func (p *Parser) parseReturnStatement() error {
	if err := p.expect(token.RETURN); err != nil {
		return err
	}
	p.nextToken()
	if expr, err := p.parseExpression(0); err != nil {
		return err
	} else {
		p.ast = append(p.ast, ast.ReturnStatement{
			Type:       ast.RETURN,
			Expression: expr,
		})
	}
	p.nextToken()
	if err := p.expect(token.SEMICOLON); err != nil {
		return err
	}
	p.nextToken()
	return nil
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
	node = ast.ExpressionStatement{
		Type:       ast.EXPR,
		Expression: node,
	}
	p.ast = append(p.ast, node)
	return nil
}

func (p *Parser) parseExpression(bindingPower int) (ast.Node, error) {
	var left ast.Node
	switch p.token().Type {
	case token.LPAREN:
		p.nextToken()
		var err error
		left, err = p.parseExpression(0)
		if err != nil {
			return nil, err
		}
		p.nextToken()
		if err := p.expect(token.RPAREN); err != nil {
			return nil, err
		}
	case token.IDENT:
		left = ast.IdentifierExpression{
			Type:       ast.IDENT,
			Identifier: p.token(),
		}
	case token.STRING:
		fallthrough
	case token.FLOAT:
		fallthrough
	case token.INT:
		left = ast.LiteralExpression{
			Type:    ast.LITERAL,
			Literal: p.token(),
		}
	default:
		return nil, fmt.Errorf("%s:%d:%d: illegal token type %q", p.token().File, p.token().Line, p.token().Column, p.token().Type)
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
			Type:     ast.BINARY,
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
		return fmt.Errorf("parse error %s %d:%d: got %q, expected %q", t.File, t.Line, t.Column, t.Type, tokenType)
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
