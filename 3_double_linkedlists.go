// Implementing a DoubleLinkedList data structure in go
// We'll implement Append, Prepend and Remove methods
// of a typical DoubleLinkedList


package main

import (
	"errors"
)

// LinkedList represents our list with it's properties
type DoubleLinkedList struct {
	Head   *DoubleNode
	Tail   *DoubleNode
	Length int
}

// Append adds a node to the end of the list
func (list *DoubleLinkedList) Append(newNode *DoubleNode) {

	if list.Length == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		lastNode := list.Tail

		lastNode.Next = newNode
		newNode.Before = lastNode

		list.Tail = newNode
	}

	list.Length++

}

// Prepend adds a node to the start of the list
func (list *DoubleLinkedList) Prepend(newNode *DoubleNode) {

	if list.Length == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		firstNode := list.Head
		list.Head = newNode
		newNode.Next = firstNode
	}

	list.Length++

}

// Remove removes a node from the list
// This has to iterate through the whole list
// to try find the node to remove
// Can be performance costly with large lists
// You might choose to use Arrays in this case
// where we can identify an element by Index

func (list *DoubleLinkedList) Remove(node *DoubleNode) {

	if list.Length == 0 {
		panic(errors.New("cannot remove element on an empty list"))
	}

	var previousPost *DoubleNode
	currentPost := list.Head

	for currentPost.Value != node.Value {
		if currentPost.Next == nil {
			panic(errors.New("no such element found with value"))
		}

		previousPost = currentPost
		currentPost = currentPost.Next
	}

	previousPost.Next = currentPost.Next

	list.Length--

}
