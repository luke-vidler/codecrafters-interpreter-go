package main

import "fmt"

// AstPrinter implements the ExprVisitor interface to print expressions
type AstPrinter struct{}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

// Print converts an expression to its string representation
func (p *AstPrinter) Print(expr Expr) string {
	result := expr.Accept(p)
	return result.(string)
}

// VisitLiteralExpr formats a literal expression
func (p *AstPrinter) VisitLiteralExpr(expr *Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}

	// For booleans, return "true" or "false"
	if b, ok := expr.Value.(bool); ok {
		return fmt.Sprintf("%t", b)
	}

	// For other values, use default formatting
	return fmt.Sprintf("%v", expr.Value)
}

// VisitGroupingExpr formats a grouping expression
func (p *AstPrinter) VisitGroupingExpr(expr *Grouping) interface{} {
	innerExpr := expr.Expression.Accept(p).(string)
	return fmt.Sprintf("(group %s)", innerExpr)
}
