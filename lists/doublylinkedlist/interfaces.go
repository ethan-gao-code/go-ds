// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package doublylinkedlist

// DoublyLinkedList defines the interface for a doubly linked list.
type DoublyLinkedList interface {
	IsEmpty() bool          // Returns true if the list is empty.
	Size() int              // Returns the number of elements in the list.
	Iterate() []interface{} // Returns a slice of all elements in the list.
	Values() []interface{}  // Alias for Iterate, returns a slice of all elements in the list.

	Add(value interface{})      // Adds an element to the list (implementation will vary).
	AddFirst(value interface{}) // Adds an element at the beginning of the list.
	Append(value interface{})   // Adds an element at the end (alias for AddLast).
	AddLast(value interface{})  // Adds an element at the end of the list.
	Prepend(value interface{})  // Adds an element at the beginning (alias for AddFirst).

	RemoveFirst() interface{} // Removes and returns the first element of the list.
	RemoveLast() interface{}  // Removes and returns the last element of the list.

	PeekFirst() interface{} // PeekFirst returns the first element of the list without removing it.
	PeekLast() interface{}  // PeekLast returns the last element of the list without removing it.

	// GetByIndex returns the element at the specified index.
	// Returns nil if index is out of bounds.
	GetByIndex(index int) interface{}

	// IndexOf returns the index of the first occurrence of the given value.
	// Returns -1 if the value is not found.
	IndexOf(value interface{}) int

	String() string // Returns a string representation of the list.
}
