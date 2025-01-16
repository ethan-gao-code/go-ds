// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

// Queue interface defines the behavior of a queue.
type Queue interface {
	Enqueue(value ...interface{}) // Add an element to the queue
	Dequeue() interface{}         // Remove and return the front element from the queue
	Peek() interface{}            // Return the front element without removing it
	IsEmpty() bool                // Check if the queue is empty
	Size() int                    // Return the number of elements in the queue
	Clear()                       // Remove all elements from the queue
	Values() []interface{}        // Return all elements in the queue as a slice
	String() string               // Return a string representation of the queue
}
