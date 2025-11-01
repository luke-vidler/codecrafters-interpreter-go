package main

import (
	"fmt"
	"time"
)

// ReturnValue is used to propagate return values up the call stack
type ReturnValue struct {
	Value interface{}
}

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
	closure     *Environment
}

func NewLoxFunction(declaration *Function, closure *Environment) *LoxFunction {
	return &LoxFunction{
		declaration: declaration,
		closure:     closure,
	}
}

func (f *LoxFunction) Arity() int {
	return len(f.declaration.Params)
}

func (f *LoxFunction) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	// Create a new environment for the function execution
	// Use the closure environment as the parent, not the current environment
	environment := NewEnclosedEnvironment(f.closure)

	// Bind parameters to arguments
	for i, param := range f.declaration.Params {
		environment.Define(param.Lexeme, arguments[i])
	}

	// Use defer/recover to catch return values
	var returnValue interface{}
	func() {
		defer func() {
			if r := recover(); r != nil {
				if ret, ok := r.(*ReturnValue); ok {
					returnValue = ret.Value
				} else {
					// Re-panic if it's not a return value
					panic(r)
				}
			}
		}()

		// Execute the function body
		interpreter.executeBlock(f.declaration.Body, environment)
	}()

	return returnValue
}

func (f *LoxFunction) String() string {
	return fmt.Sprintf("<fn %s>", f.declaration.Name.Lexeme)
}

// Bind creates a bound method with a specific instance as "this"
func (f *LoxFunction) Bind(instance *LoxInstance) *LoxFunction {
	// For now, we'll just return the function as-is
	// In a later stage with "this", we'll bind the instance to a "this" variable
	return f
}

// LoxClass represents a user-defined class
type LoxClass struct {
	name    string
	methods map[string]*LoxFunction
}

func NewLoxClass(name string, methods map[string]*LoxFunction) *LoxClass {
	return &LoxClass{
		name:    name,
		methods: methods,
	}
}

// FindMethod looks up a method by name
func (c *LoxClass) FindMethod(name string) *LoxFunction {
	if method, ok := c.methods[name]; ok {
		return method
	}
	return nil
}

func (c *LoxClass) String() string {
	return c.name
}

// Arity returns the number of arguments the class constructor takes
func (c *LoxClass) Arity() int {
	return 0
}

// Call creates a new instance of the class
func (c *LoxClass) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	instance := NewLoxInstance(c)
	return instance
}

// LoxInstance represents an instance of a class
type LoxInstance struct {
	class  *LoxClass
	fields map[string]interface{}
}

func NewLoxInstance(class *LoxClass) *LoxInstance {
	return &LoxInstance{
		class:  class,
		fields: make(map[string]interface{}),
	}
}

func (i *LoxInstance) String() string {
	return fmt.Sprintf("%s instance", i.class.name)
}

// Get retrieves a property or method from the instance
func (i *LoxInstance) Get(name Token) interface{} {
	// First check for fields
	if value, ok := i.fields[name.Lexeme]; ok {
		return value
	}

	// Then check for methods
	method := i.class.FindMethod(name.Lexeme)
	if method != nil {
		return method.Bind(i)
	}

	// Property doesn't exist - this will be handled by the interpreter
	return nil
}

// Set sets a property on the instance
func (i *LoxInstance) Set(name Token, value interface{}) {
	i.fields[name.Lexeme] = value
}
