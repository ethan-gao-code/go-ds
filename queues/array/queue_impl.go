// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"fmt"
	"strings"
)

// New creates a new instance of a non-generic queue
func New() *Queues {
	return &Queues{}
}

// Enqueue adds one or more elements to the queue
func (q *Queues) Enqueue(values ...interface{}) {
	q.items = append(q.items, values...)
}

// Dequeue removes and returns the front element from the queue
// It returns nil if the queue is empty
func (q *Queues) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front
}

// Peek returns the front element from the queue without removing it
// It returns nil if the queue is empty
func (q *Queues) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[0]
}

// IsEmpty checks if the queue is empty
func (q *Queues) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queues) Size() int {
	return len(q.items)
}

// Clear removes all elements from the queue
func (q *Queues) Clear() {
	q.items = []interface{}{}
}

// Values returns a slice containing all elements in the queue
func (q *Queues) Values() []interface{} {
	return q.items
}

// String returns a string representation of the queue
func (q *Queues) String() string {
	// Create a slice to hold string representations of the elements
	elements := make([]string, len(q.items))

	// Iterate over the set and convert each element to a string
	for i, item := range q.items {
		elements[i] = fmt.Sprintf("%v", item)
	}

	// Join all the elements with commas and wrap them in square brackets
	return fmt.Sprintf("Queue elements: [%s]", strings.Join(elements, ", "))
}
