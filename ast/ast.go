package ast

import "github.com/tobiashort/monkey/token"

type Node any
type Ast []Node

type ExpressionStatement struct {
	Expression Node
}

type BinaryExpression struct {
	Left     Node
	Operator token.Token
	Right    Node
}

type IdentifierExpression struct {
	Identifier token.Token
}

type LiteralExpression struct {
	Literal token.Token
}
