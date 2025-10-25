package main

// Expr is the interface for all expression types
type Expr interface {
	Accept(visitor ExprVisitor) interface{}
}

// ExprVisitor is the visitor interface for expressions
type ExprVisitor interface {
	VisitLiteralExpr(expr *Literal) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
}

// Literal represents a literal value expression
type Literal struct {
	Value interface{}
}

func (l *Literal) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(l)
}

// Grouping represents a parenthesized expression
type Grouping struct {
	Expression Expr
}

func (g *Grouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(g)
}
