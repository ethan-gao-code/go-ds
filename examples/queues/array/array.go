// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ethan-gao-code/go-ds/queues/array"
)

func main() {
	exampleForImpl()
	fmt.Println("==================== This is a split line ====================")
	exampleForGenericImpl()
}

func exampleForImpl() {
	// Create a new queue instance
	queue := array.New()

	// Enqueue multiple elements into the queue
	queue.Enqueue(10, "Hello", 3.14)
	fmt.Println("Queue after Enqueue:", queue)

	// Peek the front element without removing it
	peekedValue := queue.Peek()
	fmt.Println("Peeked value:", peekedValue)

	// Dequeue an element from the queue
	dequeuedValue := queue.Dequeue()
	fmt.Println("Dequeued value:", dequeuedValue)
	fmt.Println("Queue after Dequeue:", queue)

	// Enqueue more elements
	queue.Enqueue("World", 42)
	fmt.Println("Queue after more Enqueue:", queue)

	// Get the size of the queue
	queueSize := queue.Size()
	fmt.Println("Size of the queue:", queueSize)

	// Dequeue until the queue is empty
	for !queue.IsEmpty() {
		fmt.Println("Dequeued value:", queue.Dequeue())
	}

	// Print the final state of the queue
	fmt.Println("Queue after clearing all elements:", queue)
}

func exampleForGenericImpl() {
	// Create a new generic queue for integers
	intQueue := array.NewGenericQueue[int]()

	// Enqueue some elements
	intQueue.Enqueue(10, 20, 30)
	fmt.Println("Integer Queue after Enqueue:", intQueue)

	// Peek the front element without removing it
	peekedValue := intQueue.Peek()
	fmt.Println("Peeked value from integer queue:", peekedValue)

	// Dequeue an element from the queue
	dequeuedValue := intQueue.Dequeue()
	fmt.Println("Dequeued value from integer queue:", dequeuedValue)
	fmt.Println("Integer Queue after Dequeue:", intQueue)

	// Enqueue more elements
	intQueue.Enqueue(40, 50)
	fmt.Println("Integer Queue after more Enqueue:", intQueue)

	// Create a new generic queue for strings
	strQueue := array.NewGenericQueue[string]()

	// Enqueue some string elements
	strQueue.Enqueue("Hello", "World")
	fmt.Println("String Queue after Enqueue:", strQueue)

	// Peek the front element of the string queue
	strPeekedValue := strQueue.Peek()
	fmt.Println("Peeked value from string queue:", strPeekedValue)

	// Dequeue an element from the string queue
	strDequeuedValue := strQueue.Dequeue()
	fmt.Println("Dequeued value from string queue:", strDequeuedValue)
	fmt.Println("String Queue after Dequeue:", strQueue)

	// Get the size of the integer queue
	intQueueSize := intQueue.Size()
	fmt.Println("Size of the integer queue:", intQueueSize)

	// Clear the string queue
	strQueue.Clear()
	fmt.Println("String Queue after Clear:", strQueue)

	// Dequeue until the integer queue is empty
	for !intQueue.IsEmpty() {
		fmt.Println("Dequeued value from integer queue:", intQueue.Dequeue())
	}

	// Final state of the integer queue
	fmt.Println("Integer Queue after clearing all elements:", intQueue.String())
}
