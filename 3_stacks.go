// Implementing a Stack data structure in go using interfaces
// You can use any specific data type or implement generics using available libraries
// Stacks are a collection in which data is added (Push) and removed (Pop)
// in a Last In, First Out principle - LIFO
// Stacks can be useful in implementing UNDO actions, storing user history,
// and performing Postfix calculations
// We'll implement the Push, Pop, and Peek methods

package main

import "errors"

// Container holds the stack items
type Container struct {
	items []interface{}
}

// NewStack creates a new container
func (stack *Container) NewContainer() *Container {
	stack.items = []interface{}{}
	return stack
}

// Push adds an item to the top of the stack
// You can add lock to the stack during operations
func (stack *Container) Push(value interface{}) {
	stack.items = append(stack.items, value)
}

// Pop removes an item from the top of the stack
// and returns the item
// You can add lock to the stack during operations
func (stack *Container) Pop() interface{} {

	if len(stack.items) == 0 {
		panic(errors.New("cannot pop from an empty stack"))
	}

	item := stack.items[len(stack.items)-1]

	stack.items = stack.items[0 : len(stack.items)-1]

	return item

}

// Peek returns the top item in the stack
func (stack *Container) Peek() interface{} {

	if len(stack.items) == 0 {
		panic(errors.New("cannot peek from an empty stack"))
	}

	return stack.items[len(stack.items)-1]

}
