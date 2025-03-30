package evaluator

import (
	"github.com/dallinja/monkey-interpreter-go/ast"
	"github.com/dallinja/monkey-interpreter-go/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// Statements
	case *ast.Program:
		// fmt.Println("*ast.Program")
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		// fmt.Println("*ast.ExpressionStatement")
		return Eval(node.Expression)

	// Expressions
	case *ast.PrefixExpression:
		// fmt.Println("*ast.PrefixExpression")
		right := Eval(node.Right)
		// fmt.Printf("Eval(node.Right) = %s\n", right.Inspect())
		v := evalPrefixExpression(node.Operator, right)
		// fmt.Printf("evalPrefixExpression(%s, %s) = %s\n", node.Operator, right.Inspect(), v.Inspect())
		return v

	case *ast.InfixExpression:
		// fmt.Println("*ast.PrefixExpression")
		left := Eval(node.Left)
		right := Eval(node.Right)
		// fmt.Printf("Eval(node.Right) = %s\n", right.Inspect())
		v := evalInfixExpression(node.Operator, left, right)
		// fmt.Printf("evalPrefixExpression(%s, %s) = %s\n", node.Operator, right.Inspect(), v.Inspect())
		return v

	case *ast.IntegerLiteral:
		// fmt.Println("*ast.IntegerLiteral")
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		// fmt.Println("*ast.Boolean")
		v := nativeBoolToBooleanObject(node.Value)
		// fmt.Printf("nativeBoolToBooleanObject(%t) = %s\n", node.Value, v.Inspect())
		return v
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
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

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
