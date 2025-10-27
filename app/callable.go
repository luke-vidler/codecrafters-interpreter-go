package main

import "time"

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
