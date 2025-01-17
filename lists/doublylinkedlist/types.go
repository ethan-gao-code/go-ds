// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package doublylinkedlist

// Node defines the structure for a doubly linked list node.
type Node struct {
	value interface{} // Value of the node
	next  *Node       // Pointer to the next node
	prev  *Node       // Pointer to the previous node
}

// List is the implementation of the DoublyLinkedList interface.
type List struct {
	head *Node // Head of the list
	tail *Node // Tail of the list
	size int   // Size of the list
}

// GenericNode defines the structure for a generic doubly linked list node.
type GenericNode[T comparable] struct {
	value T               // Value of the node
	next  *GenericNode[T] // Pointer to the next node
	prev  *GenericNode[T] // Pointer to the previous node
}

// GenericList is the generic implementation of the DoublyLinkedList interface.
type GenericList[T comparable] struct {
	head *GenericNode[T] // Head of the list
	tail *GenericNode[T] // Tail of the list
	size int             // Size of the list
}
