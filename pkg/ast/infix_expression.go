package ast

import (
	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

// InfixExpression prefix expression
type InfixExpression struct {
	Token    token.Token // The infix token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (*InfixExpression) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String string representation
func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Operator + " " + ie.Right.String() + ")"
}

// ToJSON to json
func (ie *InfixExpression) ToJSON() string {
	return `{"type":"` + ie.Operator + `","left":` + ie.Left.ToJSON() + `,"right":` + ie.Right.ToJSON() + `}`
}
