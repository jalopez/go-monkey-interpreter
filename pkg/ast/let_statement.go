package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// LetStatement "let" statement
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (*LetStatement) statementNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	return ls.TokenLiteral() + " " +
		ls.Name.String() + " = " +
		ls.Value.String() + ";"
}

// Identifier IDENT token
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (*Identifier) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}