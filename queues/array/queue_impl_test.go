// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := New()
	q.Enqueue(1, 2, 3)
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := New()
	q.Enqueue(1, 2, 3)
	result := q.Dequeue()
	if result != 1 {
		t.Errorf("Expected 1, got %v", result)
	}
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
}

func TestQueue_Peek(t *testing.T) {
	q := New()
	q.Enqueue(1, 2, 3)
	result := q.Peek()
	if result != 1 {
		t.Errorf("Expected 1, got %v", result)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Errorf("Expected queue to not be empty")
	}
}

func TestQueue_Size(t *testing.T) {
	q := New()
	q.Enqueue(1, 2)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
}

func TestQueue_Clear(t *testing.T) {
	q := New()
	q.Enqueue(1, 2, 3)
	q.Clear()
	if q.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", q.Size())
	}
}

func TestQueue_String(t *testing.T) {
	q := New()
	q.Enqueue(1, 2, 3)
	expected := "Queue elements: [1, 2, 3]"
	if q.String() != expected {
		t.Errorf("Expected %s, got %s", expected, q.String())
	}
}

func TestQueue_Dequeue_Empty(t *testing.T) {
	q := New()
	result := q.Dequeue()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
