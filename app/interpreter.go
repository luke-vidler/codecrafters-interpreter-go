package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Interpreter evaluates expressions
type Interpreter struct {
	hadRuntimeError bool
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		hadRuntimeError: false,
	}
}

// Evaluate evaluates an expression and returns its value
func (i *Interpreter) Evaluate(expr Expr) interface{} {
	return expr.Accept(i)
}

// VisitLiteralExpr evaluates a literal expression
func (i *Interpreter) VisitLiteralExpr(expr *Literal) interface{} {
	return expr.Value
}

// VisitGroupingExpr evaluates a grouping expression
func (i *Interpreter) VisitGroupingExpr(expr *Grouping) interface{} {
	return i.Evaluate(expr.Expression)
}

// VisitUnaryExpr evaluates a unary expression
func (i *Interpreter) VisitUnaryExpr(expr *Unary) interface{} {
	right := i.Evaluate(expr.Right)

	switch expr.Operator.Type {
	case MINUS:
		// Negation: check if operand is a number
		if !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operand must be a number.")
			return nil
		}
		num := i.toNumber(right)
		return -num
	case BANG:
		// Logical not: invert truthiness
		return !i.isTruthy(right)
	}

	// Unreachable
	return nil
}

// VisitBinaryExpr evaluates a binary expression
func (i *Interpreter) VisitBinaryExpr(expr *Binary) interface{} {
	left := i.Evaluate(expr.Left)
	right := i.Evaluate(expr.Right)

	switch expr.Operator.Type {
	case PLUS:
		// Addition or string concatenation
		leftIsNum := i.isNumber(left)
		rightIsNum := i.isNumber(right)

		// Both are numbers - numeric addition
		if leftIsNum && rightIsNum {
			leftNum := i.toNumber(left)
			rightNum := i.toNumber(right)
			return leftNum + rightNum
		}

		// Check if both are non-numeric strings
		leftStr, leftIsString := left.(string)
		rightStr, rightIsString := right.(string)

		// Both are strings and neither is a number - string concatenation
		if leftIsString && rightIsString && !leftIsNum && !rightIsNum {
			return leftStr + rightStr
		}

		// If we get here, operands are not compatible (mixed types)
		i.runtimeError(expr.Operator, "Operands must be two numbers or two strings.")
		return nil
	case MINUS:
		// Subtraction - check if both operands are numbers
		if !i.isNumber(left) || !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operands must be numbers.")
			return nil
		}
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum - rightNum
	case STAR:
		// Multiplication - check if both operands are numbers
		if !i.isNumber(left) || !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operands must be numbers.")
			return nil
		}
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum * rightNum
	case SLASH:
		// Division - check if both operands are numbers
		if !i.isNumber(left) || !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operands must be numbers.")
			return nil
		}
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum / rightNum
	case GREATER:
		// Greater than - check if both operands are numbers
		if !i.isNumber(left) || !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operands must be numbers.")
			return nil
		}
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum > rightNum
	case GREATER_EQUAL:
		// Greater than or equal - check if both operands are numbers
		if !i.isNumber(left) || !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operands must be numbers.")
			return nil
		}
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum >= rightNum
	case LESS:
		// Less than - check if both operands are numbers
		if !i.isNumber(left) || !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operands must be numbers.")
			return nil
		}
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum < rightNum
	case LESS_EQUAL:
		// Less than or equal - check if both operands are numbers
		if !i.isNumber(left) || !i.isNumber(right) {
			i.runtimeError(expr.Operator, "Operands must be numbers.")
			return nil
		}
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum <= rightNum
	case EQUAL_EQUAL:
		// Equality
		return i.isEqual(left, right)
	case BANG_EQUAL:
		// Inequality
		return !i.isEqual(left, right)
	}

	// Unreachable
	return nil
}

// isEqual checks if two values are equal
func (i *Interpreter) isEqual(left, right interface{}) bool {
	// Handle nil cases
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	// Try to determine if values are numbers (either string or float64)
	leftStr, leftIsString := left.(string)
	rightStr, rightIsString := right.(string)
	leftNum, leftIsNum := left.(float64)
	rightNum, rightIsNum := right.(float64)

	// Check if left is a numeric string (from NUMBER token, has decimal point)
	leftIsNumericString := false
	if leftIsString && strings.Contains(leftStr, ".") {
		var err error
		leftNum, err = strconv.ParseFloat(leftStr, 64)
		if err == nil {
			leftIsNumericString = true
			leftIsNum = true
		}
	}

	// Check if right is a numeric string (from NUMBER token, has decimal point)
	rightIsNumericString := false
	if rightIsString && strings.Contains(rightStr, ".") {
		var err error
		rightNum, err = strconv.ParseFloat(rightStr, 64)
		if err == nil {
			rightIsNumericString = true
			rightIsNum = true
		}
	}

	// If both are numbers (either float64 or numeric strings), compare as numbers
	if leftIsNum && rightIsNum {
		return leftNum == rightNum
	}

	// If both are non-numeric strings, compare as strings
	if leftIsString && rightIsString && !leftIsNumericString && !rightIsNumericString {
		return leftStr == rightStr
	}

	// Check if both are booleans
	leftBool, leftIsBool := left.(bool)
	rightBool, rightIsBool := right.(bool)
	if leftIsBool && rightIsBool {
		return leftBool == rightBool
	}

	// Different types are not equal
	return false
}

// isTruthy determines the truthiness of a value
// false and nil are falsy, everything else is truthy
func (i *Interpreter) isTruthy(value interface{}) bool {
	if value == nil {
		return false
	}
	if b, ok := value.(bool); ok {
		return b
	}
	return true
}

// toNumber converts a value to a float64
func (i *Interpreter) toNumber(value interface{}) float64 {
	// If it's already a float64, return it
	if num, ok := value.(float64); ok {
		return num
	}

	// If it's a string (from scanner), parse it
	if str, ok := value.(string); ok {
		if num, err := strconv.ParseFloat(str, 64); err == nil {
			return num
		}
	}

	// Default to 0 if can't convert
	return 0
}

// isNumber checks if a value is a number (float64 or numeric string)
func (i *Interpreter) isNumber(value interface{}) bool {
	// Check if it's a float64
	if _, ok := value.(float64); ok {
		return true
	}

	// Check if it's a numeric string (has decimal point)
	if str, ok := value.(string); ok {
		if strings.Contains(str, ".") {
			if _, err := strconv.ParseFloat(str, 64); err == nil {
				return true
			}
		}
	}

	return false
}

// HasRuntimeError returns true if a runtime error occurred
func (i *Interpreter) HasRuntimeError() bool {
	return i.hadRuntimeError
}

// runtimeError reports a runtime error
func (i *Interpreter) runtimeError(token Token, message string) {
	i.hadRuntimeError = true
	fmt.Fprintf(os.Stderr, "%s\n[line 1]\n", message)
}

// Stringify converts a value to its string representation for output
func (i *Interpreter) Stringify(value interface{}) string {
	if value == nil {
		return "nil"
	}

	// For booleans, return "true" or "false"
	if b, ok := value.(bool); ok {
		return fmt.Sprintf("%t", b)
	}

	// For strings that represent numbers (from scanner), parse and format them
	if str, ok := value.(string); ok {
		// Try to parse as float to see if it's a number
		if num, err := strconv.ParseFloat(str, 64); err == nil {
			// It's a number - format with minimum decimal places
			formatted := strconv.FormatFloat(num, 'f', -1, 64)
			return formatted
		}
		// Not a number, return string as-is (for string literals)
		return str
	}

	// For numbers stored as float64, format them properly
	if num, ok := value.(float64); ok {
		formatted := strconv.FormatFloat(num, 'f', -1, 64)
		return formatted
	}

	// Default formatting
	return fmt.Sprintf("%v", value)
}
