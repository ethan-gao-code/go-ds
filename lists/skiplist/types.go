// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package skiplist

// List represents the skip list structure itself
type List struct {
	header *Node  // Pointer to the header node (level 0)
	tail   *Node  // Pointer to the tail node (level 0)
	length uint64 // Number of nodes in the skip list
	level  int    // Maximum level in the skip list
}

// Node represents a node in the skip list
type Node struct {
	// forward stores the pointers to the next nodes at each level
	level []*nodeLevel
	// backward stores the pointer to the previous node
	backward *Node
	// score represents the score of the node (typically used for ordered data)
	score float64
	// obj represents the object stored in the node (could be any type)
	obj interface{}
}

// nodeLevel represents a nodeLevel in the skip list with forward pointer and span
type nodeLevel struct {
	forward *Node // Pointer to the next node at this nodeLevel
	span    int   // Span represents the number of nodes between this and the next node at this nodeLevel
}

func (n *Node) GetScore() float64 {
	return n.score
}

func (n *Node) GetObj() interface{} {
	return n.obj
}
