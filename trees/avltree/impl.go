// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

// Reference: https://en.wikipedia.org/wiki/AVL_tree

package avltree

import (
	"fmt"
	"strings"

	"github.com/ethan-gao-code/go-ds/utils"
)

// New creates and returns a new AVL tree.
func New[K comparable, V any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

// Put inserts a key-value pair into the AVL tree, or updates the value if the key already exists.
func (t *Tree[K, V]) Put(key K, value V) {
	t.root = t.putNode(t.root, key, value)
}

// Remove deletes the key-value pair from the AVL tree.
func (t *Tree[K, V]) Remove(key K) {
	// Start recursive deletion from the root
	t.root, _ = t.removeNode(t.root, key)
}

// Get retrieves the value associated with the given key from the AVL tree.
// Returns the value and a boolean indicating whether the key exists in the tree.
func (t *Tree[K, V]) Get(key K) (V, bool) {
	// Start recursive search from the root
	node := t.getNode(t.root, key)
	if node != nil {
		return node.value, true
	}
	// Return zero value for the type V and false if the key does not exist
	var zeroValue V
	return zeroValue, false
}

// Contains checks if the AVL tree contains the given key.
func (t *Tree[K, V]) Contains(key K) bool {
	// Start searching from the root
	node := t.getNode(t.root, key)
	return node != nil
}

// Len returns the number of nodes in the AVL tree.
func (t *Tree[K, V]) Len() int {
	return t.size
}

// IsEmpty returns true if the AVL tree is empty, otherwise false.
func (t *Tree[K, V]) IsEmpty() bool {
	return t.size == 0
}

// InOrder performs an in-order traversal of the AVL tree and returns the result as a slice of values.
func (t *Tree[K, V]) InOrder() []interface{} {
	var result []interface{}
	t.inOrderTraversal(t.root, &result)
	return result
}

// PreOrder performs a pre-order traversal of the AVL tree and returns the result as a slice of values.
func (t *Tree[K, V]) PreOrder() []interface{} {
	var result []interface{}
	t.preOrderTraversal(t.root, &result)
	return result
}

// PostOrder performs a post-order traversal of the AVL tree and returns the result as a slice of values.
func (t *Tree[K, V]) PostOrder() []interface{} {
	var result []interface{}
	t.postOrderTraversal(t.root, &result)
	return result
}

// String returns a string representation of the AVL tree, showing its structure in levels.
func (t *Tree[K, V]) String() string {
	if t.root == nil {
		return ""
	}

	var result []string
	level := 0
	queue := []*Node[K, V]{t.root}

	// BFS traversal to print nodes by levels
	for len(queue) > 0 {
		var levelNodes []string
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, fmt.Sprintf("%v", node.value))

			// Add child nodes to the queue
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		// Join the nodes at this level and add to result
		result = append(result, fmt.Sprintf("Level %d: %s", level, strings.Join(levelNodes, " ")))
		level++
	}

	return strings.Join(result, "\n")
}

// putNode is a recursive helper function that inserts a key-value pair into the tree and balances the tree.
func (t *Tree[K, V]) putNode(node *Node[K, V], key K, value V) *Node[K, V] {
	if node == nil {
		// If the node is nil, create a new node
		t.size++ // Increment the size here when a new node is created
		return &Node[K, V]{key: key, value: value, height: 1}
	}

	// Use CompareObjects to compare the keys
	comp := utils.CompareObjects(key, node.key)

	if comp < 0 {
		// If the key is less than the current node's key, insert it into the left subtree
		node.left = t.putNode(node.left, key, value)
	} else if comp > 0 {
		// If the key is greater than the current node's key, insert it into the right subtree
		node.right = t.putNode(node.right, key, value)
	} else {
		// If the key is equal, update the value
		node.value = value
	}

	// Update the height of the current node
	node.height = 1 + max(t.getHeight(node.left), t.getHeight(node.right))

	// Balance the node and return the potentially new root of the subtree
	return t.balance(node)
}

// balance balances the node and performs rotations if necessary.
func (t *Tree[K, V]) balance(node *Node[K, V]) *Node[K, V] {
	balanceFactor := t.getBalance(node)

	// Left heavy case
	if balanceFactor > 1 {
		if t.getBalance(node.left) < 0 {
			// Left-right case, perform a left rotation on the left child
			node.left = t.leftRotate(node.left)
		}
		// Left-left case, perform a right rotation
		return t.rightRotate(node)
	}

	// Right heavy case
	if balanceFactor < -1 {
		if t.getBalance(node.right) > 0 {
			// Right-left case, perform a right rotation on the right child
			node.right = t.rightRotate(node.right)
		}
		// Right-right case, perform a left rotation
		return t.leftRotate(node)
	}

	// No imbalance, return the node
	return node
}

// getHeight returns the height of the node. If the node is nil, its height is 0.
func (t *Tree[K, V]) getHeight(node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return node.height
}

// getBalance calculates the balance factor of the node. It is the difference between the left and right subtree heights.
func (t *Tree[K, V]) getBalance(node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.left) - t.getHeight(node.right)
}

// leftRotate performs a left rotation on the node.
func (t *Tree[K, V]) leftRotate(x *Node[K, V]) *Node[K, V] {
	y := x.right
	x.right = y.left
	y.left = x
	x.height = 1 + max(t.getHeight(x.left), t.getHeight(x.right))
	y.height = 1 + max(t.getHeight(y.left), t.getHeight(y.right))
	return y
}

// rightRotate performs a right rotation on the node.
func (t *Tree[K, V]) rightRotate(y *Node[K, V]) *Node[K, V] {
	x := y.left
	y.left = x.right
	x.right = y
	y.height = 1 + max(t.getHeight(y.left), t.getHeight(y.right))
	x.height = 1 + max(t.getHeight(x.left), t.getHeight(x.right))
	return x
}

// removeNode is a recursive helper function that deletes a node from the tree and balances the tree.
func (t *Tree[K, V]) removeNode(node *Node[K, V], key K) (*Node[K, V], bool) {
	if node == nil {
		// Node not found, return nil
		return nil, false
	}

	// Use CompareObjects to compare the keys
	comp := utils.CompareObjects(key, node.key)

	var removed bool
	if comp < 0 {
		// If the key is less than the current node's key, recursively delete from the left subtree
		node.left, removed = t.removeNode(node.left, key)
	} else if comp > 0 {
		// If the key is greater than the current node's key, recursively delete from the right subtree
		node.right, removed = t.removeNode(node.right, key)
	} else {
		// The key matches, perform the deletion
		removed = true
		if node.left == nil {
			// Case 1: Node has no left child, replace with right child
			t.size-- // Decrease size when a node is removed
			return node.right, removed
		} else if node.right == nil {
			// Case 2: Node has no right child, replace with left child
			t.size-- // Decrease size when a node is removed
			return node.left, removed
		}

		// Case 3: Node has two children, find the in-order successor (or predecessor)
		// Here, we find the minimum node in the right subtree (in-order successor)
		minNode := t.getMin(node.right)
		node.key, node.value = minNode.key, minNode.value
		// Now delete the in-order successor node
		node.right, _ = t.removeNode(node.right, minNode.key)
	}

	// Update the height of the current node
	node.height = 1 + max(t.getHeight(node.left), t.getHeight(node.right))

	// Balance the node and return the potentially new root of the subtree
	return t.balance(node), removed
}

// getMin finds the minimum node in a subtree.
func (t *Tree[K, V]) getMin(node *Node[K, V]) *Node[K, V] {
	for node.left != nil {
		node = node.left
	}
	return node
}

// getNode is a recursive helper function that searches for a node by key.
func (t *Tree[K, V]) getNode(node *Node[K, V], key K) *Node[K, V] {
	if node == nil {
		// Key not found, return nil
		return nil
	}

	// Use CompareObjects to compare the keys
	comp := utils.CompareObjects(key, node.key)

	if comp < 0 {
		// If the key is less than the current node's key, search the left subtree
		return t.getNode(node.left, key)
	} else if comp > 0 {
		// If the key is greater than the current node's key, search the right subtree
		return t.getNode(node.right, key)
	}

	return node
}

// inOrderTraversal is a recursive helper function that performs in-order traversal.
func (t *Tree[K, V]) inOrderTraversal(node *Node[K, V], result *[]interface{}) {
	if node == nil {
		return
	}
	t.inOrderTraversal(node.left, result)  // Visit left subtree
	*result = append(*result, node.value)  // Visit current node
	t.inOrderTraversal(node.right, result) // Visit right subtree
}

// preOrderTraversal is a recursive helper function that performs pre-order traversal.
func (t *Tree[K, V]) preOrderTraversal(node *Node[K, V], result *[]interface{}) {
	if node == nil {
		return
	}
	*result = append(*result, node.value)   // Visit current node
	t.preOrderTraversal(node.left, result)  // Visit left subtree
	t.preOrderTraversal(node.right, result) // Visit right subtree
}

// postOrderTraversal is a recursive helper function that performs post-order traversal.
func (t *Tree[K, V]) postOrderTraversal(node *Node[K, V], result *[]interface{}) {
	if node == nil {
		return
	}
	t.postOrderTraversal(node.left, result)  // Visit left subtree
	t.postOrderTraversal(node.right, result) // Visit right subtree
	*result = append(*result, node.value)    // Visit current node
}
