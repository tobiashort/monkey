package lexer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/tobiashort/monkey/token"
	"github.com/tobiashort/utils-go/option"
)

type Lexer struct {
	filename string
	input    string
	position int
	line     int
	column   int
}

func New(filename string, input string) *Lexer {
	return &Lexer{
		filename: filename,
		input:    input,
		position: 0,
		line:     1,
		column:   1,
	}
}

func (l *Lexer) Analyze() ([]token.Token, error) {
	var tokens []token.Token
	for {
		t := l.nextToken()
		tokens = append(tokens, t)
		if t.Type == token.EOF {
			return tokens, nil
		}
		if t.Type == token.ILLEGAL {
			return tokens, fmt.Errorf("%s:%d:%d illegal token %q", l.filename, l.line, l.column, t.Literal)
		}
	}
}

func (l *Lexer) nextToken() token.Token {
	var tok = option.None[token.Token]()

	if r := l.rune(); r == '\n' {
		l.position++
		l.line++
		l.column = 1
		return l.nextToken()
	}

	if r := l.rune(); unicode.IsSpace(r) {
		l.position++
		l.column++
		return l.nextToken()
	}

	switch r := l.rune(); r {
	case '!':
		tok = option.Some(token.Token{
			Type:    token.BANG,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
		if nr := l.rune(); nr == '=' {
			tok = option.Some(token.Token{
				Type:    token.NOT_EQUAL,
				Literal: string(r) + string(nr),
				Line:    l.line,
				Column:  l.column - 1,
			})
			l.position++
			l.column++
		}
	case '=':
		tok = option.Some(token.Token{
			Type:    token.ASSIGN,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
		if nr := l.rune(); nr == '=' {
			tok = option.Some(token.Token{
				Type:    token.EQUAL,
				Literal: string(r) + string(nr),
				Line:    l.line,
				Column:  l.column - 1,
			})
			l.position++
			l.column++
		}
	case '+':
		tok = option.Some(token.Token{
			Type:    token.PLUS,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case '-':
		tok = option.Some(token.Token{
			Type:    token.MINUS,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case '/':
		tok = option.Some(token.Token{
			Type:    token.SLASH,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case '*':
		tok = option.Some(token.Token{
			Type:    token.ASTERISK,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case '<':
		tok = option.Some(token.Token{
			Type:    token.LT,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
		if nr := l.rune(); nr == '=' {
			tok = option.Some(token.Token{
				Type:    token.LEQT,
				Literal: string(r) + string(nr),
				Line:    l.line,
				Column:  l.column - 1,
			})
			l.position++
			l.column++
		}
	case '>':
		tok = option.Some(token.Token{
			Type:    token.GT,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
		if nr := l.rune(); nr == '=' {
			tok = option.Some(token.Token{
				Type:    token.GEQT,
				Literal: string(r) + string(nr),
				Line:    l.line,
				Column:  l.column - 1,
			})
			l.position++
			l.column++
		}
	case ',':
		tok = option.Some(token.Token{
			Type:    token.COMMA,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case ';':
		tok = option.Some(token.Token{
			Type:    token.SEMICOLON,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case '(':
		tok = option.Some(token.Token{
			Type:    token.LPAREN,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case ')':
		tok = option.Some(token.Token{
			Type:    token.RPAREN,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case '{':
		tok = option.Some(token.Token{
			Type:    token.LBRACE,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case '}':
		tok = option.Some(token.Token{
			Type:    token.RBRACE,
			Literal: string(r),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	case 0:
		tok = option.Some(token.Token{
			Type:    token.EOF,
			Literal: "",
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	}

	if !tok.None {
		return tok.Val
	}

	if r := l.rune(); unicode.IsLetter(r) {
		switch f := l.field(); f {
		case "let":
			tok = option.Some(token.Token{
				Type:    token.LET,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		case "fn":
			tok = option.Some(token.Token{
				Type:    token.FUNCTION,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		case "if":
			tok = option.Some(token.Token{
				Type:    token.IF,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		case "else":
			tok = option.Some(token.Token{
				Type:    token.ELSE,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		case "false":
			tok = option.Some(token.Token{
				Type:    token.FALSE,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		case "true":
			tok = option.Some(token.Token{
				Type:    token.TRUE,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		case "return":
			tok = option.Some(token.Token{
				Type:    token.RETURN,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		default:
			if f != "" {
				tok = option.Some(token.Token{
					Type:    token.IDENT,
					Literal: f,
					Line:    l.line,
					Column:  l.column,
				})
				l.position += len(f)
				l.column += len(f)
			}
		}
	}

	if !tok.None {
		return tok.Val
	}

	if r := l.rune(); unicode.IsDigit(r) {
		f := l.field()
		if _, err := strconv.ParseInt(f, 10, 64); err == nil {
			tok = option.Some(token.Token{
				Type:    token.INT,
				Literal: f,
				Line:    l.line,
				Column:  l.column,
			})
			l.position += len(f)
			l.column += len(f)
		}
	}

	if tok.None {
		tok = option.Some(token.Token{
			Type:    token.ILLEGAL,
			Literal: string(l.rune()),
			Line:    l.line,
			Column:  l.column,
		})
		l.position++
		l.column++
	}

	return tok.Val
}

func (l *Lexer) rune() rune {
	if l.position >= len(l.input) {
		return 0
	}
	return []rune(l.input)[l.position]
}

func (l *Lexer) field() string {
	fields := strings.FieldsFunc(l.input[l.position:], func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsSymbol(r) || unicode.IsPunct(r)
	})
	if len(fields) > 0 {
		return fields[0]
	}
	return ""
}
