// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ethan-gao-code/go-ds/lists/doublylinkedlist"
)

func main() {
	exampleForImpl()
	fmt.Println("==================== This is a split line ====================")
	exampleForGenericImpl()
}

func exampleForImpl() {
	// Create a new doubly linked list for interface{} type
	list := doublylinkedlist.New()

	// Add elements to the list (append at the end)
	list.Add(10)
	list.Add("Hello")
	list.Add(3.14)
	list.AddFirst("First Element")

	// Print the list
	fmt.Println("List after adding elements:", list)

	// Get the size of the list
	fmt.Println("Size of the list:", list.Size())

	// Remove the first element
	removedFirst := list.RemoveFirst()
	fmt.Println("Removed first element:", removedFirst)

	// Remove the last element
	removedLast := list.RemoveLast()
	fmt.Println("Removed last element:", removedLast)

	// Get the element at index 1
	elementAtIndex1 := list.GetByIndex(1)
	fmt.Println("Element at index 1:", elementAtIndex1)

	// Check if a value exists in the list
	index := list.IndexOf("Hello")
	fmt.Println("Index of value 'Hello':", index)

	// Peek the first and last elements without removing them
	peekFirst := list.PeekFirst()
	peekLast := list.PeekLast()
	fmt.Println("Peek first element:", peekFirst)
	fmt.Println("Peek last element:", peekLast)

	// Print the list after all the operations
	fmt.Println("List after operations:", list.String())
}

func exampleForGenericImpl() {
	// Create a new generic doubly linked list for integers
	list := doublylinkedlist.NewGenericList[int]()

	// Add elements to the list
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.AddFirst(5)

	// Print the list
	fmt.Println("List after adding elements:", list)

	// Get the size of the list
	fmt.Println("Size of the list:", list.Size())

	// Remove the first element
	removedFirst := list.RemoveFirst()
	fmt.Println("Removed first element:", removedFirst)

	// Remove the last element
	removedLast := list.RemoveLast()
	fmt.Println("Removed last element:", removedLast)

	// Get the element at index 1
	elementAtIndex1 := list.GetByIndex(1)
	fmt.Println("Element at index 1:", elementAtIndex1)

	// Check if a value exists in the list
	index := list.IndexOf(20)
	fmt.Println("Index of value 20:", index)

	// Peek the first and last elements without removing them
	peekFirst := list.PeekFirst()
	peekLast := list.PeekLast()
	fmt.Println("Peek first element:", peekFirst)
	fmt.Println("Peek last element:", peekLast)

	// Print the list after all the operations
	fmt.Println("List after operations:", list.String())
}
