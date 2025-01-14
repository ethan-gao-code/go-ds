// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"fmt"
	"strings"
)

// NewGenericQueue creates a new instance of a generic queue
func NewGenericQueue[T any]() *GenericQueue[T] {
	return &GenericQueue[T]{}
}

// Enqueue adds one or more elements to the queue
func (q *GenericQueue[T]) Enqueue(values ...T) {
	q.items = append(q.items, values...)
}

// Dequeue removes and returns the front element from the queue
// It returns the zero value of the type if the queue is empty
func (q *GenericQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front
}

// Peek returns the front element from the queue without removing it
// It returns the zero value of the type if the queue is empty
func (q *GenericQueue[T]) Peek() T {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return q.items[0]
}

// IsEmpty checks if the queue is empty
func (q *GenericQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *GenericQueue[T]) Size() int {
	return len(q.items)
}

// Clear removes all elements from the queue
func (q *GenericQueue[T]) Clear() {
	q.items = []T{}
}

// Values returns a slice containing all elements in the queue
func (q *GenericQueue[T]) Values() []T {
	return q.items
}

// String returns a string representation of the queue
func (q *GenericQueue[T]) String() string {
	// Create a slice to hold string representations of the elements
	elements := make([]string, len(q.items))

	// Iterate over the set and convert each element to a string
	for i, item := range q.items {
		elements[i] = fmt.Sprintf("%v", item)
	}

	// Join all the elements with commas and wrap them in square brackets
	return fmt.Sprintf("Queue elements: [%s]", strings.Join(elements, ", "))
}
