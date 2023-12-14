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

// ToJSON to json
func (il *IntegerLiteral) ToJSON() string {
	return `{"type":"integer","value":` + il.Token.Literal + `}`
}

// Boolean literal with boolean value
type Boolean struct {
	Token token.Token
	Value bool
}

func (*Boolean) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// String string representation
func (b *Boolean) String() string { return b.Token.Literal }

// ToJSON to json
func (b *Boolean) ToJSON() string {
	return `{"type":"boolean","value":` + b.Token.Literal + `}`
}

// StringLiteral literal with string value
type StringLiteral struct {
	Token token.Token
	Value string
}

func (*StringLiteral) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// String string representation
func (sl *StringLiteral) String() string { return sl.Token.Literal }

// ToJSON to json
func (sl *StringLiteral) ToJSON() string {
	return `{"type":"string","value":` + sl.Token.Literal + `}`
}
