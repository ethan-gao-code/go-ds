// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

// Reference: https://en.wikipedia.org/wiki/Skip_list

package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/ethan-gao-code/go-ds/utils"
)

// MaxLevel defines the maximum number of levels in the skip list
const MaxLevel = 16

// Probability defines the probability factor for random level generation
const Probability = 0.5

// New creates a new skip list instance
func New() *List {
	// Create a new skip list instance
	sl := &List{
		level:  1,                         // Start with level 1
		length: 0,                         // Initially the list is empty
		header: newNode(MaxLevel, 0, nil), // Create the header node with max level
	}

	// Initialize the header node's levels
	for i := 0; i < MaxLevel; i++ {
		sl.header.level[i] = &nodeLevel{
			forward: nil,
			span:    0,
		}
	}
	// Set the backward pointer of the header node to nil
	sl.header.backward = nil
	// Set the tail of the list to nil (no elements in the list yet)
	sl.tail = nil

	return sl
}

// newNode creates and returns a new node for the skip list
func newNode(level int, score float64, obj interface{}) *Node {
	// Allocate memory for the new node
	node := &Node{
		score: score,
		obj:   obj,
		level: make([]*nodeLevel, level),
	}

	// Initialize the forward pointers at each level to nil
	for i := 0; i < level; i++ {
		node.level[i] = &nodeLevel{
			forward: nil,
			span:    0,
		}
	}
	// Set the backward pointer to nil
	node.backward = nil

	return node
}

// randomLevel generates a random level for the new node
func randomLevel() int {
	level := 1
	// Randomly decide the level of the new node based on probability
	for rand.Float64() < Probability && level < MaxLevel {
		level++
	}
	return level
}

// Add adds a new element to the skip list
func (zsl *List) Add(score float64, obj interface{}) *Node {
	// Check if score is valid (not NaN)
	if math.IsNaN(score) {
		return nil
	}

	// Create an array to store the update nodes at each level
	update := make([]*Node, zsl.level)
	// Create an array to store the rank for each level
	rank := make([]int, zsl.level)

	// Start from the highest level and move down
	current := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		// Set the rank for this level
		if i == zsl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		// Move forward in the list until we find the right position
		// The node's score must be strictly greater than the new node's score
		for current.level[i].forward != nil && (current.level[i].forward.score < score ||
			(current.level[i].forward.score == score && utils.CompareObjects(current.level[i].forward.obj, obj) < 0)) {
			// Update the rank for the current level
			rank[i] += current.level[i].span
			// Move to the next node
			current = current.level[i].forward
		}
		// Store the current node at the update array for this level
		update[i] = current
	}

	// Randomly determine the level of the new node
	level := randomLevel()

	// If the new level is higher than the current maximum level, update the skip list
	if level > zsl.level {
		// Update the level and initialize the new levels
		for i := zsl.level; i < level; i++ {
			update = append(update, zsl.header)
			rank = append(rank, 0)
			update[i].level = append(update[i].level, &nodeLevel{}) // Expanding level array for the new level
			update[i].level[i].span = int(zsl.length)
		}
		// Update the current maximum level of the skip list
		zsl.level = level
	}

	// Create the new node
	newListNode := newNode(level, score, obj)

	// Insert the new node at each level
	for i := 0; i < level; i++ {
		// Set the forward pointer of the new node at this level
		newListNode.level[i].forward = update[i].level[i].forward
		// Set the forward pointer of the previous node to point to the new node
		update[i].level[i].forward = newListNode

		// Update the span for the new node and the previous node
		newListNode.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	// Increment span for untouched levels
	for i := level; i < zsl.level; i++ {
		update[i].level[i].span++
	}

	// Set the backward pointer of the new node
	newListNode.backward = update[0]
	if newListNode.level[0].forward != nil {
		newListNode.level[0].forward.backward = newListNode
	} else {
		// If it's the last node, set it as the tail
		zsl.tail = newListNode
	}

	// Update the length of the skip list
	zsl.length++

	return newListNode
}

// Remove removes an element by score and value.
func (zsl *List) Remove(score float64, obj interface{}) bool {
	update := make([]*Node, zsl.level)
	current := zsl.header

	// Traverse the list from top level to find the node to remove
	for i := zsl.level - 1; i >= 0; i-- {
		// Find the largest node whose score is smaller than the node to be removed
		// or if score is equal, compare lexicographically
		for current.level[i].forward != nil &&
			(current.level[i].forward.score < score ||
				(current.level[i].forward.score == score && utils.CompareObjects(current.level[i].forward.obj, obj) < 0)) {
			current = current.level[i].forward
		}
		// Record the node to update at each level
		update[i] = current
	}

	// Move to the node to be removed
	current = current.level[0].forward
	// If the node to remove is found and matches both score and object
	if current != nil && score == current.score && utils.EqualObjects(current.obj, obj) {
		// Remove the node and update all affected levels
		zsl.removeNode(current, update)
		zsl.freeNode(current)
		return true
	}

	// Node not found
	return false
}

func (zsl *List) removeNode(current *Node, update []*Node) {
	// Traverse from the highest level to the lowest level
	for i := 0; i < zsl.level; i++ {
		// If the forward pointer of update[i] points to current, update it to skip current
		if update[i].level[i].forward == current {
			// Update the span and forward pointer
			update[i].level[i].span += current.level[i].span - 1
			update[i].level[i].forward = current.level[i].forward
		} else {
			// If not pointing to current, simply decrement the span
			update[i].level[i].span -= 1
		}
	}

	// If current has a forward node, update its backward pointer
	if current.level[0].forward != nil {
		current.level[0].forward.backward = current.backward
	} else {
		// If current is the tail node, update the tail pointer
		zsl.tail = current.backward
	}

	// Remove the highest level if it's now empty
	for zsl.level > 1 && zsl.header.level[zsl.level-1].forward == nil {
		zsl.level--
	}

	// Decrease the length of the skip list
	zsl.length--
}

func (zsl *List) freeNode(current *Node) {
	// Here you can handle the freeing of resources for the node
	// Depending on your implementation, you might want to free the node's memory or handle it differently
	// In Go, garbage collection will automatically clean up the memory, so I just leave a blank function here
}

// Find finds an element by score and value
// Return nil if not found
func (zsl *List) Find(score float64, obj interface{}) *Node {
	current := zsl.header

	// Iterate from the top level downwards
	for i := zsl.level - 1; i >= 0; i-- {
		// Traverse the list at this level
		for current.level[i].forward != nil &&
			(current.level[i].forward.score < score ||
				(current.level[i].forward.score == score && utils.CompareObjects(current.level[i].forward.obj, obj) < 0)) {
			current = current.level[i].forward
		}
	}

	// Check if we found the node with matching score and object
	current = current.level[0].forward
	if current != nil && current.score == score && utils.EqualObjects(current.obj, obj) {
		return current
	}

	// Return nil if not found
	return nil
}

// Rank calculates the rank of the given score and object in the skip list
// It returns the rank (index) of the element if found, otherwise -1
func (zsl *List) Rank(score float64, obj interface{}) int {
	var rank int
	x := zsl.header

	// Start from the top-most level and move down
	for i := zsl.level - 1; i >= 0; i-- {
		// Traverse forward as long as the score is less than the target score
		// or the score is equal but the object is lexicographically smaller
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && utils.CompareObjects(x.level[i].forward.obj, obj) <= 0)) {

			// Accumulate the span of the nodes traversed
			rank += x.level[i].span
			x = x.level[i].forward
		}

		// If the current node is equal to the target node, return the accumulated rank
		if x.obj != nil && utils.CompareObjects(x.obj, obj) == 0 {
			return rank
		}
	}

	// Return -1 if not found
	return -1
}

// GetByRank retrieves an element by its rank (1-based index).
func (zsl *List) GetByRank(rank int) *Node {
	if rank < 1 || rank > int(zsl.length) {
		return nil // If the rank is out of bounds, return nil
	}

	current := zsl.header
	cumulativeRank := 0

	// Traverse from top to bottom layers
	for i := zsl.level - 1; i >= 0; i-- {
		// Traverse through the level[i] linked list
		for current.level[i].forward != nil && cumulativeRank+current.level[i].span < rank {
			cumulativeRank += current.level[i].span
			current = current.level[i].forward
		}
	}

	// After traversing, `current` will point to the node at the rank position
	current = current.level[0].forward
	if cumulativeRank+1 == rank && current != nil {
		return current // Return the node at the rank position
	}

	return nil // If no node is found at the given rank, return nil
}

// Size returns the number of elements in the skip list.
func (zsl *List) Size() int {
	return int(zsl.length)
}

// IsEmpty returns true if the skip list is empty, otherwise false.
func (zsl *List) IsEmpty() bool {
	return zsl.length == 0
}

// String returns a string representation of the skip list.
func (zsl *List) String() string {
	// Create a slice to hold string representations of the elements
	var elements []string
	// Iterate through the list starting from the lowest level (level 0)
	currentNode := zsl.header.level[0].forward
	for currentNode != nil {
		// Append the string representation of the node to the elements slice
		elements = append(elements, fmt.Sprintf("score: %.2f, obj: %s", currentNode.score, currentNode.obj))
		// Move to the next node at level 0 (forward)
		currentNode = currentNode.level[0].forward
	}

	// Join all the elements with commas and wrap them in square brackets
	return fmt.Sprintf("SkipList elements: [%s]", strings.Join(elements, ", "))
}
