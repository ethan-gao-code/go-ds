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
