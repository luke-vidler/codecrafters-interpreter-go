package main

// Stmt is the interface for all statement types
type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

// StmtVisitor is the visitor interface for statements
type StmtVisitor interface {
	VisitPrintStmt(stmt *Print) interface{}
	VisitExpressionStmt(stmt *Expression) interface{}
	VisitVarStmt(stmt *Var) interface{}
	VisitBlockStmt(stmt *Block) interface{}
}

// Print represents a print statement
type Print struct {
	Expression Expr
}

func (p *Print) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitPrintStmt(p)
}

// Expression represents an expression statement
type Expression struct {
	Expression Expr
}

func (e *Expression) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(e)
}

// Var represents a variable declaration statement
type Var struct {
	Name        Token
	Initializer Expr
}

func (v *Var) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVarStmt(v)
}

// Block represents a block statement
type Block struct {
	Statements []Stmt
}

func (b *Block) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitBlockStmt(b)
}
