// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"testing"
)

func TestGenericStacks_Push(t *testing.T) {
	stack := NewGenericStack[int]()
	stack.Push(1, 2, 3)
	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}
}

func TestGenericStacks_Pop(t *testing.T) {
	stack := NewGenericStack[int]()
	stack.Push(1, 2, 3)

	pop := stack.Pop()
	if pop != 3 {
		t.Errorf("Expected pop 3, got %d", pop)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", stack.Size())
	}

	// Test pop on empty stack
	stack.Clear()
	pop = stack.Pop()
	if pop != 0 {
		t.Errorf("Expected pop 0 (zero value), got %d", pop)
	}
}

func TestGenericStacks_Peek(t *testing.T) {
	stack := NewGenericStack[int]()
	stack.Push(1, 2, 3)

	peek := stack.Peek()
	if peek != 3 {
		t.Errorf("Expected peek 3, got %d", peek)
	}

	// Test peek on empty stack
	stack.Clear()
	peek = stack.Peek()
	if peek != 0 {
		t.Errorf("Expected peek 0 (zero value), got %d", peek)
	}
}

func TestGenericStacks_IsEmpty(t *testing.T) {
	stack := NewGenericStack[int]()

	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("Expected stack to be non-empty")
	}
}

func TestGenericStacks_Size(t *testing.T) {
	stack := NewGenericStack[int]()
	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}

	stack.Push(1, 2, 3)
	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}
}

func TestGenericStacks_Clear(t *testing.T) {
	stack := NewGenericStack[int]()
	stack.Push(1, 2, 3)
	stack.Clear()

	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after clear")
	}
}

func TestGenericStacks_Values(t *testing.T) {
	stack := NewGenericStack[int]()
	stack.Push(1, 2, 3)

	values := stack.Values()
	expected := []int{1, 2, 3}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected values[%d] = %d, got %d", i, expected[i], v)
		}
	}
}

func TestGenericStacks_WithStrings(t *testing.T) {
	stack := NewGenericStack[string]()
	stack.Push("a", "b", "c")

	peek := stack.Peek()
	if peek != "c" {
		t.Errorf("Expected peek 'c', got '%s'", peek)
	}

	pop := stack.Pop()
	if pop != "c" {
		t.Errorf("Expected pop 'c', got '%s'", pop)
	}

	values := stack.Values()
	expected := []string{"a", "b"}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected values[%d] = '%s', got '%s'", i, expected[i], v)
		}
	}
}

func TestGenericStacks_WithStruct(t *testing.T) {
	type Point struct {
		X, Y int
	}
	stack := NewGenericStack[Point]()
	stack.Push(Point{X: 1, Y: 2}, Point{X: 3, Y: 4})

	peek := stack.Peek()
	if peek != (Point{X: 3, Y: 4}) {
		t.Errorf("Expected peek {3,4}, got %v", peek)
	}

	pop := stack.Pop()
	if pop != (Point{X: 3, Y: 4}) {
		t.Errorf("Expected pop {3,4}, got %v", pop)
	}

	values := stack.Values()
	expected := []Point{{X: 1, Y: 2}}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected values[%d] = %v, got %v", i, expected[i], v)
		}
	}
}

func TestGenericStacks_String_WithIntegers(t *testing.T) {
	stack := NewGenericStack[int]()
	stack.Push(1, 2, 3)
	expected := "Stacks elements: [1, 2, 3]"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}

func TestGenericStacks_String_WithStrings(t *testing.T) {
	stack := NewGenericStack[string]()
	stack.Push("a", "b", "c")
	expected := "Stacks elements: [a, b, c]"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}

func TestGenericStacks_String_EmptyStack(t *testing.T) {
	stack := NewGenericStack[int]()
	expected := "Stacks elements: []"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}

func TestGenericStacks_String_WithStructs(t *testing.T) {
	type Point struct {
		X, Y int
	}
	stack := NewGenericStack[Point]()
	stack.Push(Point{X: 1, Y: 2}, Point{X: 3, Y: 4})
	expected := "Stacks elements: [{1 2}, {3 4}]"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}
