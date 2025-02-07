// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package skiplist

import (
	"testing"
)

func TestList_Add(t *testing.T) {
	// Create a new skip list
	skipList := New()

	// Test inserting the first element
	skipList.Add(1.0, "one")
	if skipList.Size() != 1 {
		t.Errorf("Expected size 1, got %d", skipList.Size())
	}
	if rank := skipList.Rank(1.0, "one"); rank != 1 {
		t.Errorf("Expected rank 1, got %d", rank)
	}

	// Test inserting more elements
	skipList.Add(3.0, "three")
	skipList.Add(2.0, "two")

	// Check the size after inserting 3 elements
	if skipList.Size() != 3 {
		t.Errorf("Expected size 3, got %d", skipList.Size())
	}

	// Check if the elements are in the correct order
	if rank := skipList.Rank(1.0, "one"); rank != 1 {
		t.Errorf("Expected rank 1 for 'one', got %d", rank)
	}
	if rank := skipList.Rank(2.0, "two"); rank != 2 {
		t.Errorf("Expected rank 2 for 'two', got %d", rank)
	}
	if rank := skipList.Rank(3.0, "three"); rank != 3 {
		t.Errorf("Expected rank 3 for 'three', got %d", rank)
	}

	// Test that a non-existing element returns -1
	if rank := skipList.Rank(4.0, "four"); rank != -1 {
		t.Errorf("Expected rank -1 for non-existing element, got %d", rank)
	}
}

func TestList_Remove(t *testing.T) {
	// Create a skip list and add some nodes
	zsl := New()
	zsl.Add(1.0, "a")
	zsl.Add(2.0, "b")
	zsl.Add(3.0, "c")

	// Check the length before removal
	if zsl.Size() != 3 {
		t.Fatalf("Expected skiplist length to be 3, got %d", zsl.Size())
	}

	// Test removal of existing node
	removed := zsl.Remove(2.0, "b")
	if !removed {
		t.Fatal("Expected to remove node with score 2.0 and object 'b'")
	}

	// Check the length after removal
	if zsl.Size() != 2 {
		t.Fatalf("Expected skiplist length to be 2, got %d", zsl.Size())
	}

	// Verify the node is indeed removed
	// Try to find the node that was removed (it should not be found)
	if rank := zsl.Rank(2.0, "b"); rank != -1 {
		t.Fatalf("Expected to not find node with score 2.0 and object 'b', got rank %d", rank)
	}

	// Test removal of non-existing node
	removed = zsl.Remove(4.0, "d")
	if removed {
		t.Fatal("Expected not to remove node with score 4.0 and object 'd', but it was removed")
	}

	// Check the length remains the same
	if zsl.Size() != 2 {
		t.Fatalf("Expected skiplist length to be 2 after trying to remove a non-existing node, got %d", zsl.Size())
	}

	// Test removal of the last node
	removed = zsl.Remove(3.0, "c")
	if !removed {
		t.Fatal("Expected to remove node with score 3.0 and object 'c'")
	}

	// Check the length after removal of the last node
	if zsl.Size() != 1 {
		t.Fatalf("Expected skiplist length to be 1 after removing last node, got %d", zsl.Size())
	}

	// Verify the node is removed and only one node remains
	if rank := zsl.Rank(3.0, "c"); rank != -1 {
		t.Fatalf("Expected to not find node with score 3.0 and object 'c', got rank %d", rank)
	}
}

func TestList_Remove_EmptyList(t *testing.T) {
	// Create an empty skip list
	zsl := New()

	// Try to remove from the empty list
	removed := zsl.Remove(1.0, "a")
	if removed {
		t.Fatal("Expected not to remove node from an empty list")
	}
}

func TestList_Find(t *testing.T) {
	// Create a skip list and add some nodes
	zsl := New()
	zsl.Add(1.0, "a")
	zsl.Add(2.0, "b")
	zsl.Add(3.0, "c")

	// Test finding an existing node
	node := zsl.Find(2.0, "b")
	if node == nil {
		t.Fatal("Expected to find node with score 2.0 and object 'b'")
	}
	if node.score != 2.0 || node.obj != "b" {
		t.Fatalf("Found node with incorrect values, got score: %f, obj: %s", node.score, node.obj)
	}

	// Test finding a non-existing node
	node = zsl.Find(4.0, "d")
	if node != nil {
		t.Fatal("Expected to not find node with score 4.0 and object 'd'")
	}

	// Test finding a node with the same score but different object
	node = zsl.Find(2.0, "a")
	if node != nil {
		t.Fatal("Expected to not find node with score 2.0 and object 'a'")
	}
}

func TestList_Rank(t *testing.T) {
	// Create a new skip list
	skipList := New()

	// Add some elements with different scores and objects
	skipList.Add(1.0, "one")
	skipList.Add(2.0, "two")
	skipList.Add(3.0, "three")

	// Test rank retrieval
	rank := skipList.Rank(2.0, "two")
	if rank != 2 {
		t.Errorf("Expected rank 2, got %d", rank)
	}

	rank = skipList.Rank(1.0, "one")
	if rank != 1 {
		t.Errorf("Expected rank 1, got %d", rank)
	}

	rank = skipList.Rank(4.0, "four")
	if rank != -1 {
		t.Errorf("Expected rank -1, got %d", rank)
	}
}

func TestList_GetByRank(t *testing.T) {
	// Create a skip list and add some nodes
	zsl := New()
	zsl.Add(1.0, "a")
	zsl.Add(2.0, "b")
	zsl.Add(3.0, "c")

	// Test GetByRank with valid rank
	node := zsl.GetByRank(1)
	if node == nil || node.score != 1.0 || node.obj != "a" {
		t.Fatal("Expected to find node with rank 1, score 1.0 and object 'a'")
	}

	node = zsl.GetByRank(2)
	if node == nil || node.score != 2.0 || node.obj != "b" {
		t.Fatal("Expected to find node with rank 2, score 2.0 and object 'b'")
	}

	node = zsl.GetByRank(3)
	if node == nil || node.score != 3.0 || node.obj != "c" {
		t.Fatal("Expected to find node with rank 3, score 3.0 and object 'c'")
	}

	// Test GetByRank with invalid rank
	node = zsl.GetByRank(4)
	if node != nil {
		t.Fatal("Expected to not find node with rank 4")
	}

	node = zsl.GetByRank(0)
	if node != nil {
		t.Fatal("Expected to not find node with rank 0")
	}
}

func TestList_Size(t *testing.T) {
	// Create a new skip list
	zsl := New()

	// Test size on an empty list
	if size := zsl.Size(); size != 0 {
		t.Fatalf("Expected size 0, got %d", size)
	}

	// Add some nodes to the skip list
	zsl.Add(1.0, "a")
	zsl.Add(2.0, "b")
	zsl.Add(3.0, "c")

	// Test size after adding nodes
	if size := zsl.Size(); size != 3 {
		t.Fatalf("Expected size 3, got %d", size)
	}

	// Remove a node and test size again
	zsl.Remove(2.0, "b")
	if size := zsl.Size(); size != 2 {
		t.Fatalf("Expected size 2, got %d", size)
	}

	// Remove all nodes and test size again
	zsl.Remove(1.0, "a")
	zsl.Remove(3.0, "c")
	if size := zsl.Size(); size != 0 {
		t.Fatalf("Expected size 0, got %d", size)
	}
}

func TestList_IsEmpty(t *testing.T) {
	// Create a new skip list
	zsl := New()

	// Test isEmpty on an empty list
	if !zsl.IsEmpty() {
		t.Fatal("Expected skip list to be empty")
	}

	// Add some nodes to the skip list
	zsl.Add(1.0, "a")
	zsl.Add(2.0, "b")
	zsl.Add(3.0, "c")

	// Test isEmpty after adding nodes
	if zsl.IsEmpty() {
		t.Fatal("Expected skip list to be non-empty")
	}

	// Remove a node and test isEmpty again
	zsl.Remove(2.0, "b")
	if zsl.IsEmpty() {
		t.Fatal("Expected skip list to be non-empty")
	}

	// Remove all nodes and test isEmpty again
	zsl.Remove(1.0, "a")
	zsl.Remove(3.0, "c")
	if !zsl.IsEmpty() {
		t.Fatal("Expected skip list to be empty")
	}
}

func TestList_String(t *testing.T) {
	zsl := New()

	// Add some nodes
	zsl.Add(1.0, "a")
	zsl.Add(2.0, "b")
	zsl.Add(3.0, "c")

	expected := "SkipList elements: [score: 1.00, obj: a, score: 2.00, obj: b, score: 3.00, obj: c]"
	result := zsl.String()

	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}

	// Remove a node and check again
	zsl.Remove(2.0, "b")
	expected = "SkipList elements: [score: 1.00, obj: a, score: 3.00, obj: c]"
	result = zsl.String()

	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
