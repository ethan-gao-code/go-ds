// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"testing"
)

// TestGenericQueue_Enqueue 测试 Enqueue 方法
func TestGenericQueue_Enqueue(t *testing.T) {
	queue := NewGenericQueue[int]()
	queue.Enqueue(1, 2, 3)
	if queue.Size() != 3 {
		t.Errorf("Expected size 3, got %d", queue.Size())
	}
}

// TestGenericQueue_Dequeue 测试 Dequeue 方法
func TestGenericQueue_Dequeue(t *testing.T) {
	queue := NewGenericQueue[int]()
	queue.Enqueue(1, 2, 3)

	dequeue := queue.Dequeue()
	if dequeue != 1 {
		t.Errorf("Expected dequeue 1, got %d", dequeue)
	}
	if queue.Size() != 2 {
		t.Errorf("Expected size 2 after dequeue, got %d", queue.Size())
	}

	// 测试空队列 Dequeue
	queue.Clear()
	dequeue = queue.Dequeue()
	if dequeue != 0 {
		t.Errorf("Expected dequeue 0 (zero value), got %d", dequeue)
	}
}

// TestGenericQueue_Peek 测试 Peek 方法
func TestGenericQueue_Peek(t *testing.T) {
	queue := NewGenericQueue[int]()
	queue.Enqueue(1, 2, 3)

	peek := queue.Peek()
	if peek != 1 {
		t.Errorf("Expected peek 1, got %d", peek)
	}

	// 测试空队列 Peek
	queue.Clear()
	peek = queue.Peek()
	if peek != 0 {
		t.Errorf("Expected peek 0 (zero value), got %d", peek)
	}
}

// TestGenericQueue_IsEmpty 测试 IsEmpty 方法
func TestGenericQueue_IsEmpty(t *testing.T) {
	queue := NewGenericQueue[int]()

	if !queue.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.Errorf("Expected queue to be non-empty")
	}
}

// TestGenericQueue_Size 测试 Size 方法
func TestGenericQueue_Size(t *testing.T) {
	queue := NewGenericQueue[int]()
	if queue.Size() != 0 {
		t.Errorf("Expected size 0, got %d", queue.Size())
	}

	queue.Enqueue(1, 2, 3)
	if queue.Size() != 3 {
		t.Errorf("Expected size 3, got %d", queue.Size())
	}
}

// TestGenericQueue_Clear 测试 Clear 方法
func TestGenericQueue_Clear(t *testing.T) {
	queue := NewGenericQueue[int]()
	queue.Enqueue(1, 2, 3)
	queue.Clear()

	if !queue.IsEmpty() {
		t.Errorf("Expected queue to be empty after clear")
	}
}

// TestGenericQueue_Values 测试 Values 方法
func TestGenericQueue_Values(t *testing.T) {
	queue := NewGenericQueue[int]()
	queue.Enqueue(1, 2, 3)

	values := queue.Values()
	expected := []int{1, 2, 3}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected values[%d] = %d, got %d", i, expected[i], v)
		}
	}
}

// TestGenericQueue_WithStrings 测试字符串类型的队列
func TestGenericQueue_WithStrings(t *testing.T) {
	queue := NewGenericQueue[string]()
	queue.Enqueue("a", "b", "c")

	peek := queue.Peek()
	if peek != "a" {
		t.Errorf("Expected peek 'a', got '%s'", peek)
	}

	dequeue := queue.Dequeue()
	if dequeue != "a" {
		t.Errorf("Expected dequeue 'a', got '%s'", dequeue)
	}

	values := queue.Values()
	expected := []string{"b", "c"}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected values[%d] = '%s', got '%s'", i, expected[i], v)
		}
	}
}

// TestGenericQueue_WithStructs 测试结构体类型的队列
func TestGenericQueue_WithStructs(t *testing.T) {
	type Point struct {
		X, Y int
	}
	queue := NewGenericQueue[Point]()
	queue.Enqueue(Point{X: 1, Y: 2}, Point{X: 3, Y: 4})

	peek := queue.Peek()
	if peek != (Point{X: 1, Y: 2}) {
		t.Errorf("Expected peek {1, 2}, got %v", peek)
	}

	dequeue := queue.Dequeue()
	if dequeue != (Point{X: 1, Y: 2}) {
		t.Errorf("Expected dequeue {1, 2}, got %v", dequeue)
	}

	values := queue.Values()
	expected := []Point{{X: 3, Y: 4}}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected values[%d] = %v, got %v", i, expected[i], v)
		}
	}
}

// TestGenericQueue_String_WithIntegers 测试整数类型队列的 String 方法
func TestGenericQueue_String_WithIntegers(t *testing.T) {
	queue := NewGenericQueue[int]()
	queue.Enqueue(1, 2, 3)
	expected := "Queue elements: [1, 2, 3]"
	if queue.String() != expected {
		t.Errorf("Expected %s, got %s", expected, queue.String())
	}
}

// TestGenericQueue_String_WithStrings 测试字符串类型队列的 String 方法
func TestGenericQueue_String_WithStrings(t *testing.T) {
	queue := NewGenericQueue[string]()
	queue.Enqueue("a", "b", "c")
	expected := "Queue elements: [a, b, c]"
	if queue.String() != expected {
		t.Errorf("Expected %s, got %s", expected, queue.String())
	}
}

// TestGenericQueue_String_EmptyQueue 测试空队列的 String 方法
func TestGenericQueue_String_EmptyQueue(t *testing.T) {
	queue := NewGenericQueue[int]()
	expected := "Queue elements: []"
	if queue.String() != expected {
		t.Errorf("Expected %s, got %s", expected, queue.String())
	}
}

// TestGenericQueue_String_WithStructs 测试结构体类型队列的 String 方法
func TestGenericQueue_String_WithStructs(t *testing.T) {
	type Point struct {
		X, Y int
	}
	queue := NewGenericQueue[Point]()
	queue.Enqueue(Point{X: 1, Y: 2}, Point{X: 3, Y: 4})
	expected := "Queue elements: [{1 2}, {3 4}]"
	if queue.String() != expected {
		t.Errorf("Expected %s, got %s", expected, queue.String())
	}
}
