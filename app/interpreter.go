package main

import "fmt"

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

	// For numbers, format them properly
	if num, ok := value.(float64); ok {
		return fmt.Sprintf("%v", num)
	}

	// For strings, return as-is
	if str, ok := value.(string); ok {
		return str
	}

	// Default formatting
	return fmt.Sprintf("%v", value)
}
