// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package doublylinkedlist

import "testing"

type TestPerson struct {
	Name string
	Age  int
}

func TestList_AddFirst(t *testing.T) {
	// Test with integers
	list := New()
	list.AddFirst(1)
	if list.Size() != 1 || list.Values()[0] != 1 {
		t.Errorf("AddFirst failed with integer")
	}

	// Test with strings
	list.AddFirst("hello")
	if list.Size() != 2 || list.Values()[0] != "hello" {
		t.Errorf("AddFirst failed with string")
	}

	// Test with mixed values
	list.AddFirst("world")
	list.AddFirst(42)
	if list.Values()[0] != 42 {
		t.Errorf("AddFirst failed with mixed types")
	}

	// Test with structs
	person := TestPerson{Name: "John", Age: 30}
	list.AddFirst(person)
	if list.Values()[0].(TestPerson).Name != "John" {
		t.Errorf("AddFirst failed with struct")
	}
}

func TestList_AddLast(t *testing.T) {
	// Test with integers
	list := New()
	list.AddLast(1)
	if list.Size() != 1 || list.Values()[0] != 1 {
		t.Errorf("AddLast failed with integer")
	}

	// Test with strings
	list.AddLast("hello")
	if list.Size() != 2 || list.Values()[1] != "hello" {
		t.Errorf("AddLast failed with string")
	}

	// Test with mixed values
	list.AddLast(42)
	list.AddLast("world")
	if list.Values()[2] != 42 {
		t.Errorf("AddLast failed with mixed types")
	}

	// Test with structs
	person := TestPerson{Name: "Jane", Age: 25}
	list.AddLast(person)
	if list.Values()[4].(TestPerson).Name != "Jane" {
		t.Errorf("AddLast failed with struct")
	}
}

func TestList_Prepend(t *testing.T) {
	// Test with integers
	list := New()
	list.Prepend(10)
	if list.Size() != 1 || list.Values()[0] != 10 {
		t.Errorf("Prepend failed with integer")
	}

	// Test with strings
	list.Prepend("test")
	if list.Size() != 2 || list.Values()[0] != "test" {
		t.Errorf("Prepend failed with string")
	}

	// Test with mixed values
	list.Prepend(100)
	list.Prepend("mixed")
	if list.Values()[0] != "mixed" {
		t.Errorf("Prepend failed with mixed types")
	}

	// Test with structs
	person := TestPerson{Name: "Alice", Age: 28}
	list.Prepend(person)
	if list.Values()[0].(TestPerson).Name != "Alice" {
		t.Errorf("Prepend failed with struct")
	}
}

func TestList_Append(t *testing.T) {
	// Test with integers
	list := New()
	list.Append(100)
	if list.Size() != 1 || list.Values()[0] != 100 {
		t.Errorf("Append failed with integer")
	}

	// Test with strings
	list.Append("apple")
	if list.Size() != 2 || list.Values()[1] != "apple" {
		t.Errorf("Append failed with string")
	}

	// Test with mixed values
	list.Append(50)
	list.Append("banana")
	if list.Values()[2] != 50 {
		t.Errorf("Append failed with mixed types")
	}

	// Test with structs
	person := TestPerson{Name: "Bob", Age: 40}
	list.Append(person)
	if list.Values()[4].(TestPerson).Name != "Bob" {
		t.Errorf("Append failed with struct")
	}
}

func TestList_RemoveFirst(t *testing.T) {
	// Test with integers
	list := New()
	list.Append(1)
	list.Append(2)
	removed := list.RemoveFirst()
	if removed != 1 || list.Size() != 1 || list.Values()[0] != 2 {
		t.Errorf("RemoveFirst failed with integer. Expected 1, got %v. List size: %d, first value: %v", removed, list.Size(), list.Values()[0])
	}

	// Test with strings
	list.RemoveFirst() // Remove 2
	list.Append("apple")
	removedStr := list.RemoveFirst()
	if removedStr != "apple" || list.Size() != 0 {
		t.Errorf("RemoveFirst failed with string. Expected 'apple', got %v. List size: %d", removedStr, list.Size())
	}

	// Test with structs
	list.AddFirst(TestPerson{Name: "John", Age: 25})
	removedStruct := list.RemoveFirst()
	if removedStruct.(TestPerson).Name != "John" || list.Size() != 0 {
		t.Errorf("RemoveFirst failed with struct. Expected {John, 25}, got %+v. List size: %d", removedStruct, list.Size())
	}
}

func TestList_RemoveLast(t *testing.T) {
	// Test with integers
	list := New()
	list.Append(1)
	list.Append(2)
	removed := list.RemoveLast()
	if removed != 2 || list.Size() != 1 || list.Values()[0] != 1 {
		t.Errorf("RemoveLast failed with integer. Expected 2, got %v. List size: %d, first value: %v", removed, list.Size(), list.Values()[0])
	}

	// Test with strings
	list.RemoveLast() // Remove 1
	list.Append("apple")
	removedStr := list.RemoveLast()
	if removedStr != "apple" || list.Size() != 0 {
		t.Errorf("RemoveLast failed with string. Expected 'apple', got %v. List size: %d", removedStr, list.Size())
	}

	// Test with structs
	list.AddFirst(TestPerson{Name: "John", Age: 25})
	removedStruct := list.RemoveLast()
	if removedStruct.(TestPerson).Name != "John" || list.Size() != 0 {
		t.Errorf("RemoveLast failed with struct. Expected {John, 25}, got %+v. List size: %d", removedStruct, list.Size())
	}
}

func TestList_Iterate(t *testing.T) {
	// Test with integers
	list := New()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	values := list.Iterate()
	if len(values) != 3 || values[0] != 1 || values[1] != 2 || values[2] != 3 {
		t.Errorf("Iterate failed with integers")
	}

	// Test with mixed values
	list.Append("hello")
	list.Append(3.14)
	values = list.Iterate()
	if len(values) != 5 || values[3] != "hello" || values[4] != 3.14 {
		t.Errorf("Iterate failed with mixed types")
	}

	// Test with structs
	list.Append(TestPerson{Name: "Mike", Age: 35})
	values = list.Iterate()
	if len(values) != 6 || values[5].(TestPerson).Name != "Mike" {
		t.Errorf("Iterate failed with structs")
	}
}

func TestList_PeekFirst(t *testing.T) {
	// Test with an empty list
	list := &List{}
	if result := list.PeekFirst(); result != nil {
		t.Errorf("Expected nil for empty list, got %v", result)
	}

	// Test with a list containing elements
	list.Append(10)
	list.Append(20)
	if result := list.PeekFirst(); result != 10 {
		t.Errorf("Expected 10 for first element, got %v", result)
	}
}

func TestList_PeekLast(t *testing.T) {
	// Test with an empty list
	list := &List{}
	if result := list.PeekLast(); result != nil {
		t.Errorf("Expected nil for empty list, got %v", result)
	}

	// Test with a list containing elements
	list.Append(10)
	list.Append(20)
	if result := list.PeekLast(); result != 20 {
		t.Errorf("Expected 20 for last element, got %v", result)
	}
}

func TestList_GetByIndex(t *testing.T) {
	// Test with an empty list
	list := &List{}
	if result := list.GetByIndex(0); result != nil {
		t.Errorf("Expected nil for empty list, got %v", result)
	}

	// Test with a list containing elements
	list.Append(10)
	list.Append(20)
	list.Append(30)

	// Test valid indices
	if result := list.GetByIndex(0); result != 10 {
		t.Errorf("Expected 10 at index 0, got %v", result)
	}
	if result := list.GetByIndex(1); result != 20 {
		t.Errorf("Expected 20 at index 1, got %v", result)
	}
	if result := list.GetByIndex(2); result != 30 {
		t.Errorf("Expected 30 at index 2, got %v", result)
	}

	// Test is out of bounds indices
	if result := list.GetByIndex(-1); result != nil {
		t.Errorf("Expected nil for invalid index, got %v", result)
	}
	if result := list.GetByIndex(3); result != nil {
		t.Errorf("Expected nil for invalid index, got %v", result)
	}
}

func TestList_IndexOf(t *testing.T) {
	// Test with an empty list
	list := &List{}
	if result := list.IndexOf(10); result != -1 {
		t.Errorf("Expected -1 for non-existent element in empty list, got %v", result)
	}

	// Test with a list containing elements
	list.Append(10)
	list.Append(20)
	list.Append(30)

	// Test valid values
	if result := list.IndexOf(10); result != 0 {
		t.Errorf("Expected 0 for element 10, got %v", result)
	}
	if result := list.IndexOf(20); result != 1 {
		t.Errorf("Expected 1 for element 20, got %v", result)
	}
	if result := list.IndexOf(30); result != 2 {
		t.Errorf("Expected 2 for element 30, got %v", result)
	}

	// Test non-existent value
	if result := list.IndexOf(40); result != -1 {
		t.Errorf("Expected -1 for non-existent element, got %v", result)
	}
}

func TestList_Values(t *testing.T) {
	// Test with integers
	list := New()
	list.Append(10)
	list.Append(20)
	values := list.Values()
	if len(values) != 2 || values[0] != 10 || values[1] != 20 {
		t.Errorf("Values failed with integers")
	}

	// Test with strings
	list.Append("hello")
	values = list.Values()
	if len(values) != 3 || values[2] != "hello" {
		t.Errorf("Values failed with strings")
	}

	// Test with structs
	list.Append(TestPerson{Name: "Tom", Age: 40})
	values = list.Values()
	if len(values) != 4 || values[3].(TestPerson).Name != "Tom" {
		t.Errorf("Values failed with structs")
	}
}

func TestList_String(t *testing.T) {
	// Test with integers
	list := New()
	list.AddFirst(1)
	list.AddLast(2)
	list.AddLast(3)
	result := list.String()
	expected := "List elements: [1, 2, 3]"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	// Test with strings
	list2 := New()
	list2.AddFirst("A")
	list2.AddLast("B")
	list2.AddLast("C")
	result = list2.String()
	expected = "List elements: [A, B, C]"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	// Test with struct
	list3 := New()
	list3.AddFirst(TestPerson{"Alice", 25})
	list3.AddLast(TestPerson{"Bob", 30})
	result = list3.String()
	expected = "List elements: [{Alice 25}, {Bob 30}]"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
