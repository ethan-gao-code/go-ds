// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ethan-gao-code/go-ds/stacks/array"
)

func main() {
	exampleForImpl()
	fmt.Println("==================== This is a split line ====================")
	exampleForGenericImpl()
}

func exampleForImpl() {
	// Create a new stack
	stack := array.New()

	// Push elements onto the stack
	stack.Push(10, 20, 30, 40)
	fmt.Println("Stack after pushing elements:", stack)

	// Peek the top element of the stack
	topElement := stack.Peek()
	fmt.Println("Top element of the stack:", topElement)

	// Pop elements from the stack
	poppedElement := stack.Pop()
	fmt.Println("Element popped from the stack:", poppedElement)
	fmt.Println("Stack after popping an element:", stack)

	// Check if the stack is empty
	isEmpty := stack.IsEmpty()
	fmt.Println("Is the stack empty?", isEmpty)

	// Get the size of the stack
	stackSize := stack.Size()
	fmt.Println("Size of the stack:", stackSize)

	// Clear the stack
	stack.Clear()
	fmt.Println("Stack after clearing:", stack)

	// Check if the stack is empty after clearing
	isEmptyAfterClear := stack.IsEmpty()
	fmt.Println("Is the stack empty after clearing?", isEmptyAfterClear)
}

func exampleForGenericImpl() {
	// Create a new generic stack for integers
	intStack := array.NewGenericStack[int]()

	// Push some integer values onto the stack
	intStack.Push(10, 20, 30, 40)
	fmt.Println("Integer stack after pushing elements:", intStack)

	// Peek the top element of the stack
	intTop := intStack.Peek()
	fmt.Println("Top element of the integer stack:", intTop)

	// Pop an element from the stack
	intPopped := intStack.Pop()
	fmt.Println("Element popped from the integer stack:", intPopped)
	fmt.Println("Integer stack after popping an element:", intStack)

	// Create a new generic stack for strings
	stringStack := array.NewGenericStack[string]()

	// Push some string values onto the stack
	stringStack.Push("apple", "banana", "cherry")
	fmt.Println("String stack after pushing elements:", stringStack)

	// Peek the top element of the string stack
	stringTop := stringStack.Peek()
	fmt.Println("Top element of the string stack:", stringTop)

	// Pop an element from the string stack
	stringPopped := stringStack.Pop()
	fmt.Println("Element popped from the string stack:", stringPopped)
	fmt.Println("String stack after popping an element:", stringStack)

	// Check if the stacks are empty
	fmt.Println("Is the integer stack empty?", intStack.IsEmpty())
	fmt.Println("Is the string stack empty?", stringStack.IsEmpty())

	// Get the size of the integer stack
	fmt.Println("Size of the integer stack:", intStack.Size())

	// Clear the integer stack
	intStack.Clear()
	fmt.Println("Integer stack after clearing:", intStack)

	// Check if the integer stack is empty after clearing
	fmt.Println("Is the integer stack empty after clearing?", intStack.IsEmpty())
}
