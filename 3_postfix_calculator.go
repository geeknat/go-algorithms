package main

import "C"
import (
	"fmt"
	"time"
)

// Implementing a PostFix calculator in Go

func PostFixCalculator(container Container) int {

	var emptyStack Container

	t1 := time.Now().Unix()

	for _, item := range container.items {

		switch v := item.(type) {
		default:
			fmt.Printf("unexpected type %T", v)
		case int:

			//push it into the stack
			emptyStack.Push(item)

		case string:

			rightValue := emptyStack.Pop()
			leftValue := emptyStack.Pop()

			result := 0

			if item.(string) == "+" {
				result = leftValue.(int) + rightValue.(int)
			}

			if item.(string) == "*" {
				result = leftValue.(int) * rightValue.(int)
			}

			if item.(string) == "-" {
				result = leftValue.(int) - rightValue.(int)
			}

			if item.(string) == "%" {
				result = leftValue.(int) % rightValue.(int)
			}

			fmt.Println(leftValue, rightValue, item, result)

			emptyStack.Push(result)
		}

	}

	t2 := time.Now().Unix()

	elapsedTime := t2 - t1

	fmt.Println("Finished in seconds", elapsedTime)

	return emptyStack.Peek().(int)
}
