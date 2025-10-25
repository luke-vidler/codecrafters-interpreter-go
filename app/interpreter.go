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
	// For now, just return nil as placeholder
	// We'll implement this in later stages
	return nil
}

// VisitBinaryExpr evaluates a binary expression
func (i *Interpreter) VisitBinaryExpr(expr *Binary) interface{} {
	// For now, just return nil as placeholder
	// We'll implement this in later stages
	return nil
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
