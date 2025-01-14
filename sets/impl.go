// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package sets

import (
	"fmt"
	"strings"
)

// Ensure Set implements the SetInterface
var _ Set = (*Sets)(nil)

// New creates a new instance of Set
func New(values ...interface{}) *Sets {
	result := &Sets{items: make(map[interface{}]struct{})}
	if len(values) > 0 {
		result.Add(values...)
	}
	return result
}

// Add adds one or more elements to the set
// It accepts variadic parameters, allowing one or multiple elements to be added
func (s *Sets) Add(values ...interface{}) {
	for _, value := range values {
		// Adds each element to the set.
		s.items[value] = struct{}{}
	}
}

// AddAll adds a batch of elements to the set
// It accepts a slice of elements to be added
func (s *Sets) AddAll(values []interface{}) {
	s.Add(values...)
}

// Remove removes one or more elements from the set
// It accepts variadic parameters, allowing one or multiple elements to be removed.
func (s *Sets) Remove(values ...interface{}) {
	for _, value := range values {
		// Removes each element passed as argument.
		delete(s.items, value)
	}
}

// RemoveAll removes a batch of elements from the set
// It accepts a slice of elements to be removed
func (s *Sets) RemoveAll(values []interface{}) {
	s.Remove(values...)
}

// Contains checks if all the given elements are present in the set
func (s *Sets) Contains(values ...interface{}) bool {
	for _, value := range values {
		if _, exists := s.items[value]; !exists {
			return false
		}
	}
	return true
}

// ContainsAll checks if the set contains all the elements in the provided slice
func (s *Sets) ContainsAll(items []interface{}) bool {
	return s.Contains(items...)
}

// Size returns the number of elements in the set
func (s *Sets) Size() int {
	return len(s.items)
}

// IsEmpty checks if the set is empty
func (s *Sets) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all elements from the set.
func (s *Sets) Clear() {
	s.items = make(map[interface{}]struct{})
}

// Values returns a slice containing all elements in the set
func (s *Sets) Values() []interface{} {
	values := make([]interface{}, 0, s.Size())
	for key := range s.items {
		values = append(values, key)
	}
	return values
}

// Intersection returns a new set that contains the elements
// that are present in both sets
func (s *Sets) Intersection(other *Sets) *Sets {
	result := New() // Create a new set for the result
	for item := range s.items {
		if other.Contains(item) { // Check if the item is in the other set
			result.Add(item) // Add to the result set if present in both
		}
	}
	return result
}

// Union returns a new set that contains all the elements
// from both sets (union of the sets)
func (s *Sets) Union(other *Sets) *Sets {
	result := New(s.Values()...)
	// Add all elements from the other set
	for item := range other.items {
		result.Add(item)
	}

	return result
}

// Difference returns a new set that contains elements that are in the current set
// but not in the other set (the difference of the sets)
func (s *Sets) Difference(other *Sets) *Sets {
	result := New() // Create a new set for the result

	// Iterate over the elements of the current set
	for item := range s.items {
		// Add the item to the result if it is not in the other set
		if !other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// Subset checks if the current set is a subset of the other set
func (s *Sets) Subset(other *Sets) bool {
	// Iterate over the elements of the current set
	for item := range s.items {
		// If the current element is not in the other set, return false
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// String returns a string representation of the set
func (s *Sets) String() string {
	// Create a slice to hold string representations of the elements
	var elements []string

	// Iterate over the set and convert each element to a string
	for item := range s.items {
		elements = append(elements, fmt.Sprintf("%v", item))
	}

	// Join all the elements with commas and wrap them in square brackets
	return fmt.Sprintf("Sets elements: [%s]", strings.Join(elements, ", "))
}
