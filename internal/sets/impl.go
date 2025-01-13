package sets

import "github.com/ethan-gao-code/go-ds/collection"

// Ensure Set implements the SetInterface
var _ collection.Set = (*Sets)(nil)

// New creates a new instance of Set.
func New() *Sets {
	return &Sets{items: make(map[interface{}]struct{})}
}

// Add adds one or more elements to the set.
// It accepts variadic parameters, allowing one or multiple elements to be added.
func (s *Sets) Add(values ...interface{}) {
	for _, value := range values {
		// Adds each element to the set.
		s.items[value] = struct{}{}
	}
}

// AddAll adds a batch of elements to the set.
// It accepts a slice of elements to be added.
func (s *Sets) AddAll(values []interface{}) {
	s.Add(values...)
}

// Remove removes one or more elements from the set.
// It accepts variadic parameters, allowing one or multiple elements to be removed.
func (s *Sets) Remove(values ...interface{}) {
	for _, value := range values {
		// Removes each element passed as argument.
		delete(s.items, value)
	}
}

// RemoveAll removes a batch of elements from the set.
// It accepts a slice of elements to be removed.
func (s *Sets) RemoveAll(values []interface{}) {
	s.Remove(values...)
}

// Contains checks if all the given elements are present in the set.
func (s *Sets) Contains(values ...interface{}) bool {
	for _, value := range values {
		if _, exists := s.items[value]; !exists {
			return false
		}
	}
	return true
}

// ContainsAll checks if the set contains all the elements in the provided slice.
func (s *Sets) ContainsAll(items []interface{}) bool {
	return s.Contains(items...)
}

// Size returns the number of elements in the set.
func (s *Sets) Size() int {
	return len(s.items)
}

// IsEmpty checks if the set is empty.
func (s *Sets) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all elements from the set.
func (s *Sets) Clear() {
	s.items = make(map[interface{}]struct{})
}

// Values returns a slice containing all elements in the set.
func (s *Sets) Values() []interface{} {
	values := make([]interface{}, 0, s.Size())
	for key := range s.items {
		values = append(values, key)
	}
	return values
}
