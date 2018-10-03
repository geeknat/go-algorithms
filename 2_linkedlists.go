// Implementing a LinkedList data structure in go
// LinkedLists are a collection of several Nodes
// We'll implement Append, Prepend and Remove methods
// of a typical LinkedList
// LinkedLists have a Head as starting point and
// Tail as the end point which we can easily reference

package main

import (
	"errors"
)

// LinkedList represents our list with it's properties
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// Append adds a node to the end of the list
func (list *LinkedList) Append(newNode *Node) {

	if list.Length == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		lastNode := list.Tail
		lastNode.Next = newNode
		list.Tail = newNode
	}

	list.Length++

}

// Prepend adds a node to the start of the list
func (list *LinkedList) Prepend(newNode *Node) {

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

func (list *LinkedList) Remove(node *Node) {

	if list.Length == 0 {
		panic(errors.New("cannot remove element on an empty list"))
	}

	var previousPost *Node
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
