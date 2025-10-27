package main

import (
	"fmt"
)

// Environment stores variable bindings
type Environment struct {
	values map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{
		values: make(map[string]interface{}),
	}
}

// Define adds a new variable to the environment
func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

// Get retrieves a variable's value from the environment
func (e *Environment) Get(name Token) (interface{}, error) {
	if value, ok := e.values[name.Lexeme]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("Undefined variable '%s'.", name.Lexeme)
}

// Assign updates an existing variable's value in the environment
func (e *Environment) Assign(name Token, value interface{}) error {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return nil
	}
	return fmt.Errorf("Undefined variable '%s'.", name.Lexeme)
}
