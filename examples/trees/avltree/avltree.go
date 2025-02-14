// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ethan-gao-code/go-ds/trees/avltree"
)

func main() {
	// Create a new AVL tree
	tree := avltree.New[int, string]()

	// Insert key-value pairs
	tree.Put(10, "apple")
	tree.Put(20, "banana")
	tree.Put(5, "cherry")

	// Get values associated with keys
	if value, found := tree.Get(10); found {
		fmt.Println("Key 10:", value) // Expected: "apple"
	}
	if value, found := tree.Get(20); found {
		fmt.Println("Key 20:", value) // Expected: "banana"
	}

	// Perform InOrder traversal
	inOrder := tree.InOrder()
	fmt.Println("InOrder traversal:", inOrder) // Expected: ["cherry", "apple", "banana"]

	// Check if the tree contains a key
	fmt.Println("Contains key 10:", tree.Contains(10)) // Expected: true
	fmt.Println("Contains key 15:", tree.Contains(15)) // Expected: false

	// Remove a key-value pair
	tree.Remove(20)

	// Check tree size after removal
	fmt.Println("Tree size after removal:", tree.Len()) // Expected: 2

	// Print the tree structure
	fmt.Println("Tree structure:\n", tree.String())
}
