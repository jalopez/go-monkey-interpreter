package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// IntegerLiteral literal with integer value
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (*IntegerLiteral) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// String string representation
func (il *IntegerLiteral) String() string { return il.Token.Literal }
