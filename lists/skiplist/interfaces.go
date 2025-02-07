// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package skiplist

// SkipList defines the interface for a skip list
type SkipList interface {
	Add(score float64, value interface{}) *Node     // Adds an element with a given score and value
	Remove(score float64, value interface{}) bool   // Removes an element by score and value
	Find(score float64, value interface{}) *Node    // Finds an element by score and value
	Contains(score float64, value interface{}) bool // Checks whether an element with the given score and object exists in the skip list
	Rank(score float64, value interface{}) int      // Gets the rank of an element (1-based index). Return -1 if not found
	GetByRank(rank int) *Node                       // Retrieves an element by its rank (1-based index)
	Size() int                                      // Returns the number of elements in the skip list
	IsEmpty() bool                                  // Checks if the skip list is empty
	String() string                                 // Returns a string representation of the skip list
}
