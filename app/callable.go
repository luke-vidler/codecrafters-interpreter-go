package main

import (
	"fmt"
	"time"
)

// LoxCallable is the interface for all callable objects (functions, native functions, etc.)
type LoxCallable interface {
	Arity() int
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}

// ClockNative implements the native clock() function
type ClockNative struct{}

func (c *ClockNative) Arity() int {
	return 0
}

func (c *ClockNative) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	// Return Unix timestamp as float64
	return float64(time.Now().Unix())
}

// LoxFunction represents a user-defined function
type LoxFunction struct {
	declaration *Function
}

func NewLoxFunction(declaration *Function) *LoxFunction {
	return &LoxFunction{declaration: declaration}
}

func (f *LoxFunction) Arity() int {
	return len(f.declaration.Params)
}

func (f *LoxFunction) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	// Create a new environment for the function execution
	environment := NewEnclosedEnvironment(interpreter.environment)

	// Bind parameters to arguments
	for i, param := range f.declaration.Params {
		environment.Define(param.Lexeme, arguments[i])
	}

	// Execute the function body
	interpreter.executeBlock(f.declaration.Body, environment)

	return nil
}

func (f *LoxFunction) String() string {
	return fmt.Sprintf("<fn %s>", f.declaration.Name.Lexeme)
}
