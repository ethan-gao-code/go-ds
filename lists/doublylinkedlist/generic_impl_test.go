// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package doublylinkedlist

import (
	"testing"
)

type TestUser struct {
	Name string
	Age  int
}

func TestGenericList_AddFirst_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddFirst(10)
	if list.PeekFirst() != 10 {
		t.Errorf("Expected 10, got %d", list.PeekFirst())
	}
}

func TestGenericList_AddFirst_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddFirst("hello")
	if list.PeekFirst() != "hello" {
		t.Errorf("Expected 'hello', got %s", list.PeekFirst())
	}
}

func TestGenericList_AddFirst_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddFirst(TestUser{"Alice", 25})
	if list.PeekFirst().Name != "Alice" {
		t.Errorf("Expected 'Alice', got %s", list.PeekFirst().Name)
	}
}

func TestGenericList_AddLast_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddLast(20)
	if list.PeekLast() != 20 {
		t.Errorf("Expected 20, got %d", list.PeekLast())
	}
}

func TestGenericList_AddLast_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddLast("world")
	if list.PeekLast() != "world" {
		t.Errorf("Expected 'world', got %s", list.PeekLast())
	}
}

func TestGenericList_AddLast_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddLast(TestUser{"Bob", 30})
	if list.PeekLast().Name != "Bob" {
		t.Errorf("Expected 'Bob', got %s", list.PeekLast().Name)
	}
}

func TestGenericList_RemoveFirst_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddFirst(10)
	removedValue := list.RemoveFirst()
	if removedValue != 10 {
		t.Errorf("Expected 10, got %d", removedValue)
	}
}

func TestGenericList_RemoveFirst_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddFirst("hello")
	removedValue := list.RemoveFirst()
	if removedValue != "hello" {
		t.Errorf("Expected 'hello', got %s", removedValue)
	}
}

func TestGenericList_RemoveFirst_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddFirst(TestUser{"Alice", 25})
	removedValue := list.RemoveFirst()
	if removedValue.Name != "Alice" {
		t.Errorf("Expected 'Alice', got %s", removedValue.Name)
	}
}

func TestGenericList_RemoveLast_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddLast(20)
	removedValue := list.RemoveLast()
	if removedValue != 20 {
		t.Errorf("Expected 20, got %d", removedValue)
	}
}

func TestGenericList_RemoveLast_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddLast("world")
	removedValue := list.RemoveLast()
	if removedValue != "world" {
		t.Errorf("Expected 'world', got %s", removedValue)
	}
}

func TestGenericList_RemoveLast_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddLast(TestUser{"Bob", 30})
	removedValue := list.RemoveLast()
	if removedValue.Name != "Bob" {
		t.Errorf("Expected 'Bob', got %s", removedValue.Name)
	}
}

func TestGenericList_PeekFirst_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddFirst(10)
	peekValue := list.PeekFirst()
	if peekValue != 10 {
		t.Errorf("Expected 10, got %d", peekValue)
	}
}

func TestGenericList_PeekFirst_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddFirst("hello")
	peekValue := list.PeekFirst()
	if peekValue != "hello" {
		t.Errorf("Expected 'hello', got %s", peekValue)
	}
}

func TestGenericList_PeekFirst_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddFirst(TestUser{"Alice", 25})
	peekValue := list.PeekFirst()
	if peekValue.Name != "Alice" {
		t.Errorf("Expected 'Alice', got %s", peekValue.Name)
	}
}

func TestGenericList_PeekLast_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddLast(20)
	peekValue := list.PeekLast()
	if peekValue != 20 {
		t.Errorf("Expected 20, got %d", peekValue)
	}
}

func TestGenericList_PeekLast_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddLast("world")
	peekValue := list.PeekLast()
	if peekValue != "world" {
		t.Errorf("Expected 'world', got %s", peekValue)
	}
}

func TestGenericList_PeekLast_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddLast(TestUser{"Bob", 30})
	peekValue := list.PeekLast()
	if peekValue.Name != "Bob" {
		t.Errorf("Expected 'Bob', got %s", peekValue.Name)
	}
}

func TestGenericList_GetByIndex_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddFirst(10)
	list.AddLast(20)
	getValue := list.GetByIndex(1)
	if getValue != 20 {
		t.Errorf("Expected 20, got %d", getValue)
	}
}

func TestGenericList_GetByIndex_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddFirst("hello")
	list.AddLast("world")
	getValue := list.GetByIndex(1)
	if getValue != "world" {
		t.Errorf("Expected 'world', got %s", getValue)
	}
}

func TestGenericList_GetByIndex_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddFirst(TestUser{"Alice", 25})
	list.AddLast(TestUser{"Bob", 30})
	getValue := list.GetByIndex(1)
	if getValue.Name != "Bob" {
		t.Errorf("Expected 'Bob', got %s", getValue.Name)
	}
}

func TestGenericList_IndexOf_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddFirst(10)
	list.AddLast(20)
	index := list.IndexOf(20)
	if index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}
}

func TestGenericList_IndexOf_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddFirst("hello")
	list.AddLast("world")
	index := list.IndexOf("world")
	if index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}
}

func TestGenericList_IndexOf_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddFirst(TestUser{"Alice", 25})
	list.AddLast(TestUser{"Bob", 30})
	index := list.IndexOf(TestUser{"Bob", 30})
	if index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}
}

func TestGenericList_Size_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddFirst(10)
	list.AddLast(20)
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
}

func TestGenericList_Size_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddFirst("hello")
	list.AddLast("world")
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
}

func TestGenericList_Size_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddFirst(TestUser{"Alice", 25})
	list.AddLast(TestUser{"Bob", 30})
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
}

func TestGenericList_Values_Int(t *testing.T) {
	list := NewGenericList[int]()
	list.AddFirst(10)
	list.AddLast(20)
	values := list.Values()
	if values[0] != 10 || values[1] != 20 {
		t.Errorf("Expected [10, 20], got %v", values)
	}
}

func TestGenericList_Values_String(t *testing.T) {
	list := NewGenericList[string]()
	list.AddFirst("hello")
	list.AddLast("world")
	values := list.Values()
	if values[0] != "hello" || values[1] != "world" {
		t.Errorf("Expected ['hello', 'world'], got %v", values)
	}
}

func TestGenericList_Values_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	list.AddFirst(TestUser{"Alice", 25})
	list.AddLast(TestUser{"Bob", 30})
	values := list.Values() // Now values will be of type []Person
	if values[0].Name != "Alice" || values[1].Name != "Bob" {
		t.Errorf("Expected [{Alice 25}, {Bob 30}], got %v", values)
	}
}

func TestGenericList_IsEmpty_Int(t *testing.T) {
	list := NewGenericList[int]()
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}
	list.AddFirst(10)
	if list.IsEmpty() {
		t.Errorf("Expected list to be non-empty")
	}
}

func TestGenericList_IsEmpty_String(t *testing.T) {
	list := NewGenericList[string]()
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}
	list.AddFirst("hello")
	if list.IsEmpty() {
		t.Errorf("Expected list to be non-empty")
	}
}

func TestGenericList_IsEmpty_Struct(t *testing.T) {
	list := NewGenericList[TestUser]()
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}
	list.AddFirst(TestUser{"Alice", 25})
	if list.IsEmpty() {
		t.Errorf("Expected list to be non-empty")
	}
}

func TestGenericList_String(t *testing.T) {
	// Test with integers
	list := NewGenericList[int]()
	list.AddFirst(1)
	list.AddLast(2)
	list.AddLast(3)
	result := list.String()
	expected := "List elements: [1, 2, 3]"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	// Test with strings
	list2 := NewGenericList[string]()
	list2.AddFirst("A")
	list2.AddLast("B")
	list2.AddLast("C")
	result = list2.String()
	expected = "List elements: [A, B, C]"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	// Test with struct
	list3 := NewGenericList[TestUser]()
	list3.AddFirst(TestUser{"Alice", 25})
	list3.AddLast(TestUser{"Bob", 30})
	result = list3.String()
	expected = "List elements: [{Alice 25}, {Bob 30}]"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
