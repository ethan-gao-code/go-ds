// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"fmt"
	"strings"
)

// NewGenericStack creates a new instance of a generic stack.
func NewGenericStack[T any]() *GenericStacks[T] {
	return &GenericStacks[T]{}
}

// Push adds one or more elements to the stack.
func (s *GenericStacks[T]) Push(values ...T) {
	s.items = append(s.items, values...)
}

// Pop removes and returns the top element from the stack.
// It returns the zero value of the type if the stack is empty.
func (s *GenericStacks[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	topIndex := len(s.items) - 1
	top := s.items[topIndex]
	s.items = s.items[:topIndex]
	return top
}

// Peek returns the top element from the stack without removing it.
// It returns the zero value of the type if the stack is empty.
func (s *GenericStacks[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
func (s *GenericStacks[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack.
func (s *GenericStacks[T]) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack.
func (s *GenericStacks[T]) Clear() {
	s.items = []T{}
}

// Values returns a slice containing all elements in the stack.
func (s *GenericStacks[T]) Values() []T {
	return s.items
}

// String returns a string representation of the stack.
func (s *GenericStacks[T]) String() string {
	// Create a slice to hold string representations of the elements
	elements := make([]string, len(s.items))

	// Iterate over the set and convert each element to a string
	for i, item := range s.items {
		elements[i] = fmt.Sprintf("%v", item)
	}

	// Join all the elements with commas and wrap them in square brackets
	return fmt.Sprintf("Stacks elements: [%s]", strings.Join(elements, ", "))
}
