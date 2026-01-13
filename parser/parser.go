package parser

import (
	"github.com/tobiashort/monkey/ast"
	"github.com/tobiashort/monkey/token"
	"github.com/tobiashort/utils-go/errors"
	"github.com/tobiashort/utils-go/slices"
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
		case token.LBRACE:
			if block, err := p.parseBlock(); err != nil {
				return p.ast, err
			} else {
				p.ast = append(p.ast, block)
			}
		case token.LET:
			if err := p.parseLetStatement(); err != nil {
				return p.ast, err
			}
		case token.RETURN:
			if err := p.parseReturnStatement(); err != nil {
				return p.ast, err
			}
		case token.YIELD:
			if err := p.parseYieldStatement(); err != nil {
				return p.ast, err
			}
		case token.IF:
			if err := p.parseIfStatement(); err != nil {
				return p.ast, err
			}
		case token.FUNCTION:
			if err := p.parseFunction(); err != nil {
				return p.ast, err
			}
		case token.LPAREN, token.INT, token.FLOAT, token.STRING, token.IDENT:
			if err := p.parseExpressionStatement(); err != nil {
				return p.ast, err
			}
		default:
			return p.ast, errors.WithCtxf("%s:%d:%d: illegal token type %q", p.token().File, p.token().Line, p.token().Column, p.token().Type)
		}
		p.nextToken()
	}
	return p.ast, nil
}

func (p *Parser) parseBlock() (ast.Node, error) {
	if err := p.expect(token.LBRACE); err != nil {
		return nil, err
	}
	depth := 1
	lbrace := p.token()
	tokens := make([]token.Token, 0)
	for p.token().Type != token.EOF {
		t := p.nextToken()
		if t.Type == token.LBRACE {
			depth++
			tokens = append(tokens, t)
		} else if t.Type == token.RBRACE {
			depth--
			if depth == 0 {
				break
			} else {
				tokens = append(tokens, t)
			}
		} else {
			tokens = append(tokens, t)
		}
	}
	if depth != 0 {
		return nil, errors.WithCtxf("%s:%d:%d: unclosed block", lbrace.File, lbrace.Line, lbrace.Column)
	}
	tokens = append(tokens, token.Token{
		Type:    token.EOF,
		Literal: "",
		File:    tokens[len(tokens)-1].File,
		Line:    tokens[len(tokens)-1].Line,
		Column:  tokens[len(tokens)-1].Column + 1,
	})
	np := New(tokens)
	nast, err := np.Parse()
	if err != nil {
		return nast, err
	}
	return ast.Block{
		Type: ast.BLOCK,
		Ast:  nast,
	}, nil
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
	return nil
}

func (p *Parser) parseYieldStatement() error {
	if err := p.expect(token.YIELD); err != nil {
		return err
	}
	p.nextToken()
	if expr, err := p.parseExpression(0); err != nil {
		return err
	} else {
		p.ast = append(p.ast, ast.YieldStatement{
			Type:       ast.YIELD,
			Expression: expr,
		})
	}
	p.nextToken()
	if err := p.expect(token.SEMICOLON); err != nil {
		return err
	}
	return nil
}

func (p *Parser) parseIfStatement() error {
	if err := p.expect(token.IF); err != nil {
		return err
	}
	stmt := ast.IfStatement{
		Type: ast.IF,
	}
	p.nextToken()
	if cond, err := p.parseExpression(0); err != nil {
		return err
	} else {
		stmt.Condition = cond
		p.nextToken()
	}
	if cons, err := p.parseBlock(); err != nil {
		return err
	} else {
		stmt.Consequence = cons
	}
	if p.hasNext() && p.peekToken().Type == token.ELSE {
		p.nextToken()
		p.nextToken()
		if alt, err := p.parseBlock(); err != nil {
			return err
		} else {
			stmt.Alternative = alt
		}
	}
	p.ast = append(p.ast, stmt)
	return nil
}

func (p *Parser) parseIfExpr() (ast.Node, error) {
	if err := p.expect(token.IF); err != nil {
		return nil, err
	}
	expr := ast.IfExpression{
		Type: ast.IFEXPR,
	}
	p.nextToken()
	if cond, err := p.parseExpression(0); err != nil {
		return nil, err
	} else {
		expr.Condition = cond
		p.nextToken()
	}
	if cons, err := p.parseBlock(); err != nil {
		return nil, err
	} else {
		expr.Consequence = cons
	}
	if p.hasNext() && p.peekToken().Type == token.ELSE {
		p.nextToken()
		p.nextToken()
		if alt, err := p.parseBlock(); err != nil {
			return nil, err
		} else {
			expr.Alternative = alt
		}
	}
	return expr, nil
}

func (p *Parser) parseFunction() error {
	if err := p.expect(token.FUNCTION); err != nil {
		return err
	}
	f := ast.Function{Type: ast.FUNCTION}
	p.nextToken()
	if err := p.expect(token.IDENT); err != nil {
		return err
	}
	f.Identifier = p.token()
	p.nextToken()
	if params, err := p.parseParameters(); err != nil {
		return err
	} else {
		f.Parameters = params
	}
	p.nextToken()
	if err := p.expect(token.LBRACE); err != nil {
		return err
	}
	if block, err := p.parseBlock(); err != nil {
		return err
	} else {
		f.Block = block
	}
	p.ast = append(p.ast, f)
	return nil
}

func (p *Parser) parseParameters() ([]ast.Node, error) {
	if err := p.expect(token.LPAREN); err != nil {
		return nil, err
	}
	startToken := p.token()
	paramTokens := make([]token.Token, 0)
	depth := 1
	for {
		p.nextToken()
		if p.token().Type == token.EOF {
			break
		} else if p.token().Type == token.RPAREN {
			depth--
			if depth == 0 {
				break
			} else {
				paramTokens = append(paramTokens, p.token())
			}
		} else if p.token().Type == token.LPAREN {
			depth++
			paramTokens = append(paramTokens, p.token())
		} else {
			paramTokens = append(paramTokens, p.token())
		}
	}
	if depth != 0 {
		return nil, errors.WithCtxf("%s:%d:%d: unclosed parameters", startToken.File, startToken.Line, startToken.Column)
	}
	paramTokensSplit := slices.Split(paramTokens, func(t token.Token) bool { return t.Type == token.COMMA })
	params := make([]ast.Node, 0)
	for _, ts := range paramTokensSplit {
		if expr, err := New(ts).parseExpression(0); err != nil {
			return params, err
		} else {
			params = append(params, expr)
		}
	}
	return params, nil
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
	case token.MINUS, token.BANG:
		operator := p.token()
		p.nextToken()
		right, err := p.parseExpression(0)
		if err != nil {
			return nil, err
		}
		left = ast.UnaryExpression{
			Type:     ast.UNARY,
			Operator: operator,
			Right:    right,
		}
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
		if p.hasNext() && p.peekToken().Type == token.LPAREN {
			call := ast.CallExpression{
				Type:       ast.CALL,
				Identifier: p.token(),
			}
			p.nextToken()
			if params, err := p.parseParameters(); err != nil {
				return nil, err
			} else {
				call.Parameters = params
			}
			left = call
		} else {
			left = ast.IdentifierExpression{
				Type:       ast.IDENT,
				Identifier: p.token(),
			}
		}
	case token.IF:
		if expr, err := p.parseIfExpr(); err != nil {
			return nil, err
		} else {
			left = expr
		}
	case token.STRING, token.FLOAT, token.INT:
		left = ast.LiteralExpression{
			Type:    ast.LITERAL,
			Literal: p.token(),
		}
	default:
		return nil, errors.WithCtxf("%s:%d:%d: illegal token type %q", p.token().File, p.token().Line, p.token().Column, p.token().Type)
	}

	for p.hasNext() {
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
		return errors.WithCtxf("parse error %s %d:%d: got %q, expected %q", t.File, t.Line, t.Column, t.Type, tokenType)
	}
	return nil
}

func (p *Parser) token() token.Token {
	return p.tokens[p.position]
}

func (p *Parser) hasNext() bool {
	return p.position < len(p.tokens)-1
}

func (p *Parser) peekToken() token.Token {
	return p.tokens[p.position+1]
}

func (p *Parser) nextToken() token.Token {
	p.position++
	return p.tokens[p.position]
}
