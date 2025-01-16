// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package doublylinkedlist

import "fmt"

// New creates a new doubly linked list.
func New() *List {
	return &List{}
}

// IsEmpty checks if the list is empty.
func (l *List) IsEmpty() bool {
	return l.size == 0
}

// Size returns the number of elements in the list.
func (l *List) Size() int {
	return l.size
}

// Iterate returns all elements in the list as a slice.
func (l *List) Iterate() []interface{} {
	var values []interface{}
	current := l.head
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	return values
}

// Values is an alias for Iterate, returning all elements in the list as a slice.
func (l *List) Values() []interface{} {
	return l.Iterate()
}

// Add adds an element to the list (appends it at the end).
func (l *List) Add(value interface{}) {
	l.AddFirst(value)
}

// AddFirst adds an element at the beginning of the list.
func (l *List) AddFirst(value interface{}) {
	node := &Node{value: value}
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
func (l *List) Append(value interface{}) {
	l.AddLast(value)
}

// AddLast adds an element at the end of the list.
func (l *List) AddLast(value interface{}) {
	node := &Node{value: value}
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
func (l *List) Prepend(value interface{}) {
	l.AddFirst(value)
}

// RemoveFirst removes and returns the first element of the list.
func (l *List) RemoveFirst() interface{} {
	if l.IsEmpty() {
		return nil
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
func (l *List) RemoveLast() interface{} {
	if l.IsEmpty() {
		return nil
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

// String returns a string representation of the list.
func (l *List) String() string {
	var result string
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%v ", current.value)
		current = current.next
	}
	return result
}
