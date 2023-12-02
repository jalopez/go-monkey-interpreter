package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// LetStatement "let" statement
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (*LetStatement) statementNode() {}

// TokenLiteral token literal
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier IDENT token
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (*Identifier) expressionNode() {}

// TokenLiteral token literal
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
