package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (*IfExpression) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String string representation
func (ie *IfExpression) String() string {
	out := "if"
	out += ie.Condition.String()
	out += " "
	out += ie.Consequence.String()

	if ie.Alternative != nil {
		out += "else "
		out += ie.Alternative.String()
	}

	return out
}
