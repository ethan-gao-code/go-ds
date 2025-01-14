// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"testing"
)

func TestNew(t *testing.T) {
	stack := New()

	if stack == nil {
		t.Error("Expected New() to return a non-nil instance")
	}
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after initialization")
	}
}

func TestPush(t *testing.T) {
	stack := New()

	stack.Push(1, "hello")
	if stack.Size() != 2 {
		t.Errorf("Expected stack size to be 2, got %d", stack.Size())
	}
	if stack.Peek() != "hello" {
		t.Errorf("Expected top element to be 'hello', got %v", stack.Peek())
	}

	stack.Push(3.14, "world")
	if stack.Size() != 4 {
		t.Errorf("Expected stack size to be 4, got %d", stack.Size())
	}
	if stack.Peek() != "world" {
		t.Errorf("Expected top element to be 'world', got %v", stack.Peek())
	}
}

func TestPop(t *testing.T) {
	stack := New()

	stack.Push(1, "hello", 3.14)
	if stack.Pop() != 3.14 {
		t.Errorf("Expected popped element to be 3.14, got %v", stack.Pop())
	}
	if stack.Size() != 2 {
		t.Errorf("Expected stack size to be 2, got %d", stack.Size())
	}

	stack.Pop()
	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after popping all elements")
	}
	if stack.Pop() != nil {
		t.Errorf("Expected popped element to be nil when stack is empty, got %v", stack.Pop())
	}
}

func TestPeek(t *testing.T) {
	stack := New()

	if stack.Peek() != nil {
		t.Errorf("Expected Peek to return nil for an empty stack, got %v", stack.Peek())
	}

	stack.Push(1, "hello")
	if stack.Peek() != "hello" {
		t.Errorf("Expected Peek to return 'hello', got %v", stack.Peek())
	}

	stack.Pop()
	if stack.Peek() != 1 {
		t.Errorf("Expected Peek to return 1, got %v", stack.Peek())
	}
}

func TestIsEmpty(t *testing.T) {
	stack := New()

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after initialization")
	}

	stack.Push(1, "hello")
	if stack.IsEmpty() {
		t.Error("Expected stack to be non-empty after pushing elements")
	}

	stack.Pop()
	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after popping all elements")
	}
}

func TestSize(t *testing.T) {
	stack := New()

	if stack.Size() != 0 {
		t.Errorf("Expected stack size to be 0, got %d", stack.Size())
	}

	stack.Push(1, "hello", 3.14)
	if stack.Size() != 3 {
		t.Errorf("Expected stack size to be 3, got %d", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 2 {
		t.Errorf("Expected stack size to be 2, got %d", stack.Size())
	}
}

func TestClear(t *testing.T) {
	stack := New()

	stack.Push(1, "hello", 3.14)
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after Clear()")
	}
	if stack.Size() != 0 {
		t.Errorf("Expected stack size to be 0 after Clear(), got %d", stack.Size())
	}
}

func TestValues(t *testing.T) {
	stack := New()

	stack.Push(1, "hello", 3.14)
	values := stack.Values()
	expected := []interface{}{1, "hello", 3.14}

	if len(values) != len(expected) {
		t.Errorf("Expected values length to be %d, got %d", len(expected), len(values))
	}

	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected value at index %d to be %v, got %v", i, expected[i], v)
		}
	}
}

func TestStacks_String_WithIntegers(t *testing.T) {
	stack := New()
	stack.Push(1, 2, 3)
	expected := "Stacks elements: [1, 2, 3]"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}

func TestStacks_String_WithStrings(t *testing.T) {
	stack := New()
	stack.Push("a", "b", "c")
	expected := "Stacks elements: [a, b, c]"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}

func TestStacks_String_MixedTypes(t *testing.T) {
	stack := New()
	stack.Push(1, "a", 2.5)
	expected := "Stacks elements: [1, a, 2.5]"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}

func TestStacks_String_EmptyStack(t *testing.T) {
	stack := New()
	expected := "Stacks elements: []"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}

func TestStacks_String_WithStructs(t *testing.T) {
	type Point struct {
		X, Y int
	}
	stack := New()
	stack.Push(Point{X: 1, Y: 2}, Point{X: 3, Y: 4})
	expected := "Stacks elements: [{1 2}, {3 4}]"
	if stack.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stack.String())
	}
}
