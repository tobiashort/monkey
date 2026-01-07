package ast

import "github.com/tobiashort/monkey/token"

type NodeType = string

const (
	LET     = "LET"
	RETURN  = "RETURN"
	EXPR    = "EXPR"
	UNARY   = "UNARY"
	BINARY  = "BINARY"
	IDENT   = "IDENT"
	LITERAL = "LITERAL"
	IF      = "IF"
	BLOCK   = "BLOCK"
)

type Node any
type Ast []Node

type Block struct {
	Type NodeType
	Ast  Ast
}

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

type UnaryExpression struct {
	Type     NodeType
	Operator token.Token
	Right    Node
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

type IfExpression struct {
	Type        NodeType
	Condition   Node
	Consequence Node
	Alternative Node
}
