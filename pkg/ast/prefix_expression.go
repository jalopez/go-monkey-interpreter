package ast

import (
	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

// PrefixExpression prefix expression
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (*PrefixExpression) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String string representation
func (pe *PrefixExpression) String() string {
	return "(" + pe.Operator + pe.Right.String() + ")"
}

// ToJSON to json
func (pe *PrefixExpression) ToJSON() string {
	return `{"type":"` + pe.Operator + `","value":` + pe.Right.ToJSON() + `}`
}
