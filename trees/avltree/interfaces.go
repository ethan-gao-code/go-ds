// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

// Reference: https://en.wikipedia.org/wiki/AVL_tree

package avltree

// AVLTree defines the operations that an AVL tree should support
type AVLTree interface {
	Put(key interface{}, value interface{})  // Insert a key-value pair into the tree
	Remove(key interface{})                  // Delete the node with the given key from the tree
	Get(key interface{}) (interface{}, bool) // Retrieve the value associated with the key from the tree
	Contains(key interface{}) bool           // Check if the tree contains the key
	Len() int                                // Return the number of nodes in the tree
	IsEmpty() bool                           // Check if the tree is empty
	InOrder() []interface{}                  // Return the in-order traversal of the tree
	PreOrder() []interface{}                 // Return the pre-order traversal of the tree
	PostOrder() []interface{}                // Return the post-order traversal of the tree
	String() string                          // Return a string representation of the tree
}
