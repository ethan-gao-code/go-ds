// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package avltree

// Tree represents the AVL tree itself.
type Tree[K comparable, V any] struct {
	root *Node[K, V] // Root node of the AVL tree
	size int         // Number of nodes in the tree
}

// Node represents a node in the AVL tree.
type Node[K comparable, V any] struct {
	key    K           // Key must be comparable
	value  V           // Value can be any type
	height int         // Height of the node in the tree
	left   *Node[K, V] // Left child node
	right  *Node[K, V] // Right child node
}
