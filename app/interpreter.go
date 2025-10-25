package main

import (
	"fmt"
	"strconv"
)

// Interpreter evaluates expressions
type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
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
		// Negation: convert to number and negate
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
		// Check if both operands are strings (not number strings)
		leftStr, leftIsString := left.(string)
		rightStr, rightIsString := right.(string)

		if leftIsString && rightIsString {
			// Check if they are numeric strings (from scanner)
			_, leftErr := strconv.ParseFloat(leftStr, 64)
			_, rightErr := strconv.ParseFloat(rightStr, 64)

			// If both can be parsed as numbers, treat as numeric addition
			if leftErr == nil && rightErr == nil {
				leftNum := i.toNumber(left)
				rightNum := i.toNumber(right)
				return leftNum + rightNum
			}

			// Otherwise, it's string concatenation
			return leftStr + rightStr
		}

		// Numeric addition (for float64 values or mixed types)
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum + rightNum
	case MINUS:
		// Subtraction
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum - rightNum
	case STAR:
		// Multiplication
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum * rightNum
	case SLASH:
		// Division
		leftNum := i.toNumber(left)
		rightNum := i.toNumber(right)
		return leftNum / rightNum
	}

	// Unreachable
	return nil
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
