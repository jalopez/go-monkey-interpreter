package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// ExpressionStatement statement
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (*ExpressionStatement) statementNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String string representation
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
