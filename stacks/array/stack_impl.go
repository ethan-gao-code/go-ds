// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"fmt"
	"strings"
)

// Stacks defines a stack data structure without generics.
type Stacks struct {
	items []interface{}
}

// New creates a new instance of a non-generic stack.
func New() *Stacks {
	return &Stacks{}
}

// Push adds one or more elements to the stack.
func (s *Stacks) Push(values ...interface{}) {
	s.items = append(s.items, values...)
}

// Pop removes and returns the top element from the stack.
// It returns nil if the stack is empty.
func (s *Stacks) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	topIndex := len(s.items) - 1
	top := s.items[topIndex]
	s.items = s.items[:topIndex]
	return top
}

// Peek returns the top element from the stack without removing it.
// It returns nil if the stack is empty.
func (s *Stacks) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stacks) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack.
func (s *Stacks) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack.
func (s *Stacks) Clear() {
	s.items = []interface{}{}
}

// Values returns a slice containing all elements in the stack.
func (s *Stacks) Values() []interface{} {
	return s.items
}

// String returns a slice containing all elements in the stack.
func (s *Stacks) String() string {
	// Create a slice to hold string representations of the elements
	elements := make([]string, len(s.items))

	// Iterate over the set and convert each element to a string
	for i, item := range s.items {
		elements[i] = fmt.Sprintf("%v", item)
	}

	// Join all the elements with commas and wrap them in square brackets
	return fmt.Sprintf("Stacks elements: [%s]", strings.Join(elements, ", "))
}
