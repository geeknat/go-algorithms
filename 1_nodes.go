// Implementing Nodes in go
// Nodes are the basic building blocks of a data structure
// Nodes have two properties
// 1. Ability to store a value
// 2. Provides a mechanism to link up with another node through a reference pointer
// called the Next pointer
// Nodes join up to form a NodeChain which forms basis of what we'll look next
// called LinkedLists
// DoubleNodes have two reference pointers Previous & Next pointers

package main

// Node represent a single Node
type Node struct {
	Value interface{}
	Next  *Node
}

// DoubleNode represents a double node used to create a DoublyLinkedList
type DoubleNode struct {
	Value  interface{}
	Next   *DoubleNode
	Before *DoubleNode
}

