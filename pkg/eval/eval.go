package eval

import (
	"github.com/jalopez/go-monkey-interpreter/pkg/ast"
	"github.com/jalopez/go-monkey-interpreter/pkg/eval/object"
	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

var (
	// NULL null
	NULL = &object.Null{}
	// TRUE true
	TRUE = &object.Boolean{Value: true}
	// FALSE false
	FALSE = &object.Boolean{Value: false}
)

// Eval evaluates an AST node
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node)

	case *ast.IfExpression:
		return evalIfExpression(node)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: val}

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	condition = maybeIntegerToBoolean(condition)

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	}

	if ie.Alternative != nil {
		return Eval(ie.Alternative)
	}

	return NULL
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			// TODO add error if more statements after return
			return returnValue.Value
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement)

		if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
			// TODO add error if more statements after return
			return result
		}
	}

	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	if operator == token.BANG {
		return evalBangOperatorExpression(right)
	}
	if operator == token.MINUS {
		return evalMinusPrefixOperatorExpression(right)
	}

	return NULL
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == token.EQ:
		return nativeBoolToBooleanObject(left == right)
	case operator == token.NOTEQ:
		return nativeBoolToBooleanObject(left != right)
	}

	return NULL
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case token.PLUS:
		return &object.Integer{Value: leftVal + rightVal}
	case token.MINUS:
		return &object.Integer{Value: leftVal - rightVal}
	case token.ASTERISK:
		return &object.Integer{Value: leftVal * rightVal}
	case token.SLASH:
		return &object.Integer{Value: leftVal / rightVal}
	case token.LT:
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case token.GT:
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case token.EQ:
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case token.NOTEQ:
		return nativeBoolToBooleanObject(leftVal != rightVal)
	}
	return NULL
}

func evalBangOperatorExpression(right object.Object) object.Object {
	right = maybeIntegerToBoolean(right)

	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func maybeIntegerToBoolean(input object.Object) object.Object {
	obj, ok := input.(*object.Integer)

	if !ok {
		return input
	}

	if obj.Value == 0 {
		return FALSE
	}

	return TRUE
}

// nolint:revive
func nativeBoolToBooleanObject(input bool) object.Object {
	if input {
		return TRUE
	}

	return FALSE
}
