package sets

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	s := New()

	// Test adding a single element
	s.Add(1)
	if _, exists := s.items[1]; !exists {
		t.Errorf("Add failed: Set should contain element 1")
	}

	// Test adding multiple elements
	s.Add(2, 3)
	if _, exists := s.items[2]; !exists {
		t.Errorf("Add failed: Set should contain element 2")
	}
	if _, exists := s.items[3]; !exists {
		t.Errorf("Add failed: Set should contain element 3")
	}

	// Test that duplicates are not added
	s.Add(1) // Adding 1 again should not change the set
	if len(s.items) != 3 {
		t.Errorf("Add failed: Set should have 3 unique elements, got %d", len(s.items))
	}

	// Test adding string elements
	s.Add("apple", "banana")
	if _, exists := s.items["apple"]; !exists {
		t.Errorf("Add failed: Set should contain element 'apple'")
	}
	if _, exists := s.items["banana"]; !exists {
		t.Errorf("Add failed: Set should contain element 'banana'")
	}

	// Test that duplicates are not added
	s.Add(1) // Adding 1 again should not change the set
	if len(s.items) != 5 {
		t.Errorf("Add failed: Set should have 5 unique elements, got %d", len(s.items))
	}
}

func TestAddAll(t *testing.T) {
	s := New()

	// Test adding a slice of number elements using AddAll
	s.AddAll([]interface{}{1, 2, 3})
	if _, exists := s.items[1]; !exists {
		t.Errorf("AddAll failed: Set should contain element 1")
	}
	if _, exists := s.items[2]; !exists {
		t.Errorf("AddAll failed: Set should contain element 2")
	}
	if _, exists := s.items[3]; !exists {
		t.Errorf("AddAll failed: Set should contain element 3")
	}

	// Test adding string elements using AddAll
	s.AddAll([]interface{}{"apple", "banana", "cherry"})
	if _, exists := s.items["apple"]; !exists {
		t.Errorf("AddAll failed: Set should contain element 'apple'")
	}
	if _, exists := s.items["banana"]; !exists {
		t.Errorf("AddAll failed: Set should contain element 'banana'")
	}
	if _, exists := s.items["cherry"]; !exists {
		t.Errorf("AddAll failed: Set should contain element 'cherry'")
	}

	// Test adding elements that already exist
	s.AddAll([]interface{}{3, 4})
	if len(s.items) != 7 {
		t.Errorf("AddAll failed: Set should have 7 unique elements, got %d", len(s.items))
	}
}

func TestRemove(t *testing.T) {
	s := New()
	s.Add(1, 2, 3, 4)
	s.Add("apple", "banana", "cherry")

	// Test removing a single number element
	s.Remove(1)
	if _, exists := s.items[1]; exists {
		t.Errorf("Remove failed: Set should not contain element 1")
	}

	// Test removing a single string element
	s.Remove("apple")
	if _, exists := s.items["apple"]; exists {
		t.Errorf("Remove failed: Set should not contain element 'apple'")
	}

	// Test removing multiple elements (numbers)
	s.Remove(2, 3)
	if _, exists := s.items[2]; exists {
		t.Errorf("Remove failed: Set should not contain element 2")
	}
	if _, exists := s.items[3]; exists {
		t.Errorf("Remove failed: Set should not contain element 3")
	}

	// Test removing multiple elements (strings)
	s.Remove("banana", "cherry")
	if _, exists := s.items["banana"]; exists {
		t.Errorf("Remove failed: Set should not contain element 'banana'")
	}
	if _, exists := s.items["cherry"]; exists {
		t.Errorf("Remove failed: Set should not contain element 'cherry'")
	}
}

func TestRemoveAll(t *testing.T) {
	s := New()
	s.Add(1, 2, 3, 4)
	s.Add("apple", "banana", "cherry")

	// Test removing multiple number elements using RemoveAll
	s.RemoveAll([]interface{}{1, 2})
	if _, exists := s.items[1]; exists {
		t.Errorf("RemoveAll failed: Set should not contain element 1")
	}
	if _, exists := s.items[2]; exists {
		t.Errorf("RemoveAll failed: Set should not contain element 2")
	}
	if _, exists := s.items[3]; !exists {
		t.Errorf("RemoveAll failed: Set should contain element 3")
	}

	// Test removing multiple string elements using RemoveAll
	s.RemoveAll([]interface{}{"banana", "cherry"})
	if _, exists := s.items["banana"]; exists {
		t.Errorf("RemoveAll failed: Set should not contain element 'banana'")
	}
	if _, exists := s.items["cherry"]; exists {
		t.Errorf("RemoveAll failed: Set should not contain element 'cherry'")
	}

	// Test removing elements that don't exist
	s.RemoveAll([]interface{}{"grape", "kiwi"}) // These elements are not in the set
	if len(s.items) != 3 {
		t.Errorf("RemoveAll failed: Set should still contain 3 elements, got %d", len(s.items))
	}
}

func TestContains(t *testing.T) {
	s := New()
	s.Add(1, 2, 3, "apple", "banana")

	// Test checking if a number is contained
	if !s.Contains(1) {
		t.Errorf("Contains failed: Set should contain element 1")
	}

	// Test checking if a string is contained
	if !s.Contains("apple") {
		t.Errorf("Contains failed: Set should contain element 'apple'")
	}

	// Test checking if multiple elements are contained (all values in variadic args)
	if !s.Contains(1, "apple", 2) {
		t.Errorf("Contains failed: Set should contain elements 1, 'apple', and 2")
	}

	// Test checking if an element is not contained
	if s.Contains(4) {
		t.Errorf("Contains failed: Set should not contain element 4")
	}

	if s.Contains("grape") {
		t.Errorf("Contains failed: Set should not contain element 'grape'")
	}

	// Test checking if multiple elements are not contained (should return false)
	if s.Contains(1, "apple", "grape") {
		t.Errorf("Contains failed: Set should not contain 'grape'")
	}
}

func TestContainsAll(t *testing.T) {
	s := New()
	s.Add(1, 2, 3, "apple", "banana")

	// Test checking if all elements are contained
	if !s.ContainsAll([]interface{}{1, "apple"}) {
		t.Errorf("ContainsAll failed: Set should contain elements 1 and 'apple'")
	}

	// Test checking if not all elements are contained
	if s.ContainsAll([]interface{}{1, 4}) {
		t.Errorf("ContainsAll failed: Set should not contain element 4")
	}

	// Test checking if set contains all elements that don't exist
	if s.ContainsAll([]interface{}{"grape", "kiwi"}) {
		t.Errorf("ContainsAll failed: Set should not contain elements 'grape' and 'kiwi'")
	}
}

func TestSize(t *testing.T) {
	s := New()

	// Test size of an empty set
	if s.Size() != 0 {
		t.Errorf("Size failed: Expected size 0, got %d", s.Size())
	}

	// Add elements and test size
	s.Add(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("Size failed: Expected size 3, got %d", s.Size())
	}

	// Add duplicate elements and test size (should not increase)
	s.Add(3, 2, 1)
	if s.Size() != 3 {
		t.Errorf("Size failed: Expected size 3 after adding duplicates, got %d", s.Size())
	}

	// Add more elements and test size
	s.Add("apple", "banana")
	if s.Size() != 5 {
		t.Errorf("Size failed: Expected size 5, got %d", s.Size())
	}

	// Remove elements and test size
	s.Remove(1)
	if s.Size() != 4 {
		t.Errorf("Size failed: Expected size 4 after removing an element, got %d", s.Size())
	}

	// Remove non-existent element and test size (should not change)
	s.Remove("grape")
	if s.Size() != 4 {
		t.Errorf("Size failed: Expected size 4 after removing non-existent element, got %d", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := New()

	// Test empty set
	if !s.IsEmpty() {
		t.Errorf("IsEmpty failed: Expected true for an empty set")
	}

	// Add elements and test
	s.Add(1, 2, 3)
	if s.IsEmpty() {
		t.Errorf("IsEmpty failed: Expected false for a non-empty set")
	}

	// Remove all elements and test
	s.Remove(1, 2, 3)
	if !s.IsEmpty() {
		t.Errorf("IsEmpty failed: Expected true after removing all elements")
	}
}

func TestClear(t *testing.T) {
	s := New()

	// Add some elements
	s.Add(1, 2, 3, "apple", "banana")

	// Clear the set
	s.Clear()

	// Test if the set is empty
	if !s.IsEmpty() {
		t.Errorf("Clear failed: Set should be empty after clear")
	}

	// Test the size of the set
	if s.Size() != 0 {
		t.Errorf("Clear failed: Expected size 0, got %d", s.Size())
	}

	// Test adding elements after clear
	s.Add(4, 5)
	if s.Size() != 2 {
		t.Errorf("Clear failed: Expected size 2 after adding elements, got %d", s.Size())
	}
	if !s.Contains(4, 5) {
		t.Errorf("Clear failed: Set should contain elements 4 and 5 after adding")
	}
}

func TestValues(t *testing.T) {
	s := New()

	// Helper function to check if a slice contains a value
	contains := func(slice []interface{}, value interface{}) bool {
		for _, v := range slice {
			if reflect.DeepEqual(v, value) {
				return true
			}
		}
		return false
	}

	// Test empty set
	if len(s.Values()) != 0 {
		t.Errorf("Values failed: Expected empty slice, got %v", s.Values())
	}

	// Add some elements
	s.Add(1, 2, 3, "apple", "banana")

	// Get values and test
	values := s.Values()

	// Ensure the correct number of elements are returned
	if len(values) != 5 {
		t.Errorf("Values failed: Expected 5 elements, got %d", len(values))
	}

	// Ensure all elements are present in the returned slice
	expected := []interface{}{1, 2, 3, "apple", "banana"}
	for _, v := range expected {
		if !contains(values, v) {
			t.Errorf("Values failed: Expected element %v to be present in the result", v)
		}
	}
}
