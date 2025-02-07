// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ethan-gao-code/go-ds/lists/skiplist"
)

func main() {
	exampleForImpl()
}

func exampleForImpl() {
	// Seed the random number generator to ensure randomness
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a new skip list
	sl := skiplist.New()

	// Add some nodes to the skip list
	sl.Add(10.5, "Item 1")
	sl.Add(20.2, "Item 2")
	sl.Add(15.8, "Item 3")
	sl.Add(5.3, "Item 4")
	sl.Add(30.7, "Item 5")

	// Print the skip list
	fmt.Println("Skip List after adding elements:")
	fmt.Println(sl)

	// Find a node by score and object
	score := 15.8
	obj := "Item 3"
	node := sl.Find(score, obj)
	if node != nil {
		fmt.Printf("\nFound node with score %.2f and object %s\n", node.GetScore(), node.GetObj())
	} else {
		fmt.Println("\nNode not found")
	}

	// Get the rank of a specific node
	rank := sl.Rank(score, obj)
	if rank != -1 {
		fmt.Printf("\nRank of node with score %.2f and object %s: %d\n", score, obj, rank)
	} else {
		fmt.Println("\nRank not found")
	}

	// Remove a node from the skip list
	removed := sl.Remove(20.2, "Item 2")
	if removed {
		fmt.Println("\nRemoved node with score 20.2 and object 'Item 2'")
	} else {
		fmt.Println("\nNode to remove not found")
	}

	// Print the skip list after removal
	fmt.Println("\nSkip List after removing an element:")
	fmt.Println(sl)

	// Get a node by rank
	rankToGet := 2
	nodeByRank := sl.GetByRank(rankToGet)
	if nodeByRank != nil {
		fmt.Printf("\nNode at rank %d: score %.2f, object %s\n", rankToGet, nodeByRank.GetScore(), nodeByRank.GetObj())
	} else {
		fmt.Println("\nNode not found at the given rank")
	}

	// Check the size of the skip list
	fmt.Printf("\nSize of the skip list: %d\n", sl.Size())

	// Check if the skip list is empty
	if sl.IsEmpty() {
		fmt.Println("\nThe skip list is empty.")
	} else {
		fmt.Println("\nThe skip list is not empty.")
	}
}
