// Run examples here
package main

import (
	"fmt"
)

func main() {

	container := Container{}
	stack := container.NewContainer()
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push("*")
	stack.Push("+")
	stack.Push(1)
	stack.Push("-")

	fmt.Println(PostFixCalculator(container))
}

func stackExample() {

	container := Container{}

	stack := container.NewContainer()

	stack.Push(6)

	// Prints 6
	fmt.Println(stack.Peek())

	stack.Push(19)

	stack.Push(20)

	// Print 20
	fmt.Println(stack.Peek())

	stack.Pop()

	// Print 19
	fmt.Println(stack.Peek())
}

func linkedListExample() {

	newList := LinkedList{}

	newList.Append(&Node{Value: 10})

	fmt.Println(newList.Tail.Value)

	newList.Append(&Node{Value: 12})

	fmt.Println(newList.Head.Value)

}

func nodeExample() {

	firstNode := Node{Value: 3}

	secondNode := Node{Value: 5}

	thirdNode := Node{Value: 7}

	fourthNode := Node{Value: 9}

	firstNode.Next = &secondNode

	secondNode.Next = &thirdNode

	thirdNode.Next = &fourthNode

	printList(firstNode)

}

func printList(node Node) {

	defer func() {
		if recover() != nil {
			fmt.Println("Dead")
		}
	}()

	for node.Value != 0 {

		_, err := fmt.Println(node)

		if err != nil {
			break
		}

		if node.Next == nil {
			break
		}

		node = *node.Next
	}

}
