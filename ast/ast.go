package ast

import "github.com/tobiashort/monkey/token"

type NodeType = string

const (
	LET     = "LET"
	RETURN  = "RETURN"
	EXPR    = "EXPR"
	BINARY  = "BINARY"
	IDENT   = "IDENT"
	LITERAL = "LITERAL"
)

type Node any
type Ast []Node

type LetStatement struct {
	Type       NodeType
	Identifier token.Token
	Expression Node
}

type ReturnStatement struct {
	Type       NodeType
	Expression Node
}

type ExpressionStatement struct {
	Type       NodeType
	Expression Node
}

type BinaryExpression struct {
	Type     NodeType
	Left     Node
	Operator token.Token
	Right    Node
}

type IdentifierExpression struct {
	Type       NodeType
	Identifier token.Token
}

type LiteralExpression struct {
	Type    NodeType
	Literal token.Token
}
