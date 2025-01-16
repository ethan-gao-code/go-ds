// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package doublylinkedlist

import "testing"

type Person struct {
	Name string
	Age  int
}

func TestAddFirst(t *testing.T) {
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
	person := Person{Name: "John", Age: 30}
	list.AddFirst(person)
	if list.Values()[0].(Person).Name != "John" {
		t.Errorf("AddFirst failed with struct")
	}
}

func TestAddLast(t *testing.T) {
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
	person := Person{Name: "Jane", Age: 25}
	list.AddLast(person)
	if list.Values()[4].(Person).Name != "Jane" {
		t.Errorf("AddLast failed with struct")
	}
}

func TestPrepend(t *testing.T) {
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
	person := Person{Name: "Alice", Age: 28}
	list.Prepend(person)
	if list.Values()[0].(Person).Name != "Alice" {
		t.Errorf("Prepend failed with struct")
	}
}

func TestAppend(t *testing.T) {
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
	person := Person{Name: "Bob", Age: 40}
	list.Append(person)
	if list.Values()[4].(Person).Name != "Bob" {
		t.Errorf("Append failed with struct")
	}
}

func TestRemoveFirst(t *testing.T) {
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
	list.AddFirst(Person{Name: "John", Age: 25})
	removedStruct := list.RemoveFirst()
	if removedStruct.(Person).Name != "John" || list.Size() != 0 {
		t.Errorf("RemoveFirst failed with struct. Expected {John, 25}, got %+v. List size: %d", removedStruct, list.Size())
	}
}

func TestRemoveLast(t *testing.T) {
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
	list.AddFirst(Person{Name: "John", Age: 25})
	removedStruct := list.RemoveLast()
	if removedStruct.(Person).Name != "John" || list.Size() != 0 {
		t.Errorf("RemoveLast failed with struct. Expected {John, 25}, got %+v. List size: %d", removedStruct, list.Size())
	}
}

func TestIterate(t *testing.T) {
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
	list.Append(Person{Name: "Mike", Age: 35})
	values = list.Iterate()
	if len(values) != 6 || values[5].(Person).Name != "Mike" {
		t.Errorf("Iterate failed with structs")
	}
}

func TestValues(t *testing.T) {
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
	list.Append(Person{Name: "Tom", Age: 40})
	values = list.Values()
	if len(values) != 4 || values[3].(Person).Name != "Tom" {
		t.Errorf("Values failed with structs")
	}
}
