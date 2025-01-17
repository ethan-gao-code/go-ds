// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package doublylinkedlist

import (
	"fmt"
	"strings"
)

// NewGenericList creates a new generic doubly linked list.
func NewGenericList[T comparable]() *GenericList[T] {
	return &GenericList[T]{}
}

// IsEmpty checks if the list is empty.
func (l *GenericList[T]) IsEmpty() bool {
	return l.size == 0
}

// Size returns the number of elements in the list.
func (l *GenericList[T]) Size() int {
	return l.size
}

// Iterate returns all elements in the list as a slice.
func (l *GenericList[T]) Iterate() []T {
	var values []T
	current := l.head
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	return values
}

// Values is an alias for Iterate, returning all elements in the list as a slice.
func (l *GenericList[T]) Values() []T {
	return l.Iterate()
}

// Add adds an element to the list (appends it at the end).
func (l *GenericList[T]) Add(value T) {
	l.AddLast(value)
}

// AddFirst adds an element at the beginning of the list.
func (l *GenericList[T]) AddFirst(value T) {
	node := &GenericNode[T]{value: value}
	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size++
}

// Append adds an element at the end (alias for AddLast).
func (l *GenericList[T]) Append(value T) {
	l.AddLast(value)
}

// AddLast adds an element at the end of the list.
func (l *GenericList[T]) AddLast(value T) {
	node := &GenericNode[T]{value: value}
	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
	l.size++
}

// Prepend adds an element at the beginning (alias for AddFirst).
func (l *GenericList[T]) Prepend(value T) {
	l.AddFirst(value)
}

// RemoveFirst removes and returns the first element of the list.
func (l *GenericList[T]) RemoveFirst() T {
	if l.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	removedNode := l.head
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}

	l.size--
	return removedNode.value
}

// RemoveLast removes and returns the last element of the list.
func (l *GenericList[T]) RemoveLast() T {
	if l.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	removedNode := l.tail
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	l.size--
	return removedNode.value
}

// PeekFirst returns the first element of the doubly linked list without removing it.
func (l *GenericList[T]) PeekFirst() T {
	if l.head == nil {
		var zeroValue T
		return zeroValue
	}
	return l.head.value
}

// PeekLast returns the last element of the doubly linked list without removing it.
func (l *GenericList[T]) PeekLast() T {
	if l.tail == nil {
		var zeroValue T
		return zeroValue
	}
	return l.tail.value
}

// GetByIndex returns the element at the specified index.
func (l *GenericList[T]) GetByIndex(index int) T {
	if index < 0 || index >= l.size {
		var zeroValue T
		return zeroValue
	}

	if index < l.size/2 {
		current := l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current.value
	} else {
		current := l.tail
		for i := l.size - 1; i > index; i-- {
			current = current.prev
		}
		return current.value
	}
}

// IndexOf returns the index of the first occurrence of the given value.
// Returns -1 if the value is not found.
func (l *GenericList[T]) IndexOf(value T) int {
	current := l.head
	for i := 0; current != nil; i++ {
		if current.value == value {
			return i
		}
		current = current.next
	}
	return -1
}

// String returns a string representation of the list.
func (l *GenericList[T]) String() string {
	// Create a slice to hold string representations of the elements
	elements := make([]string, l.size)

	// Iterate over the list and convert each element to a string
	current := l.head
	for i := 0; current != nil; i++ {
		elements[i] = fmt.Sprintf("%v", current.value)
		current = current.next
	}

	// Join all the elements with commas and wrap them in square brackets
	return fmt.Sprintf("List elements: [%s]", strings.Join(elements, ", "))
}
