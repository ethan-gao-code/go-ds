// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package avltree

import (
	"testing"
)

func TestTree_Put(t *testing.T) {
	tree := New[int, string]()

	// Test 1: Insert a single element
	tree.Put(10, "a")
	if size := tree.Len(); size != 1 {
		t.Errorf("expected size 1, got %d", size)
	}
	if value, found := tree.Get(10); !found || value != "a" {
		t.Errorf("expected value 'a', got %v", value)
	}

	// Test 2: Insert another element
	tree.Put(20, "b")
	if size := tree.Len(); size != 2 {
		t.Errorf("expected size 2, got %d", size)
	}
	if value, found := tree.Get(20); !found || value != "b" {
		t.Errorf("expected value 'b', got %v", value)
	}

	// Test 3: Insert a smaller element (Left rotation should occur)
	tree.Put(5, "c")
	if size := tree.Len(); size != 3 {
		t.Errorf("expected size 3, got %d", size)
	}
	if value, found := tree.Get(5); !found || value != "c" {
		t.Errorf("expected value 'c', got %v", value)
	}

	// Test 4: Update an existing key
	tree.Put(10, "updated")
	if value, found := tree.Get(10); !found || value != "updated" {
		t.Errorf("expected updated value 'updated', got %v", value)
	}

	// Test 5: Insert a large number of elements and ensure tree remains balanced
	for i := 30; i < 100; i += 10 {
		tree.Put(i, "value")
	}
	if size := tree.Len(); size != 10 { // 10 + 1 (for 20) + 3 (additional elements) = 14
		t.Errorf("expected size 14, got %d", size)
	}

	// Test 6: Verify that a non-existing key returns false
	if _, found := tree.Get(999); found {
		t.Error("expected key 999 not to be found")
	}
}

func TestTree_Remove(t *testing.T) {
	tree := New[int, string]()

	// Insert some elements
	tree.Put(10, "a")
	tree.Put(20, "b")
	tree.Put(5, "c")
	tree.Put(15, "d")
	tree.Put(25, "e")

	// Test 1: Remove a leaf node (no children)
	tree.Remove(5) // Removing the leaf node 5
	if size := tree.Len(); size != 4 {
		t.Errorf("expected size 4, got %d", size)
	}
	if _, found := tree.Get(5); found {
		t.Errorf("expected key 5 to be removed")
	}

	// Test 2: Remove a node with one child
	tree.Remove(20) // Removing node with one child (15)
	if size := tree.Len(); size != 3 {
		t.Errorf("expected size 3, got %d", size)
	}
	if _, found := tree.Get(20); found {
		t.Errorf("expected key 20 to be removed")
	}
	if value, found := tree.Get(15); !found || value != "d" {
		t.Errorf("expected value 'd' for key 15, got %v", value)
	}

	// Test 3: Remove a node with two children
	tree.Remove(10) // Removing node with two children (15 and 25)
	if size := tree.Len(); size != 2 {
		t.Errorf("expected size 2, got %d", size)
	}
	if _, found := tree.Get(10); found {
		t.Errorf("expected key 10 to be removed")
	}
	if value, found := tree.Get(15); !found || value != "d" {
		t.Errorf("expected value 'd' for key 15, got %v", value)
	}
	if value, found := tree.Get(25); !found || value != "e" {
		t.Errorf("expected value 'e' for key 25, got %v", value)
	}

	// Test 4: Remove a node that does not exist
	tree.Remove(999) // Removing a non-existing node
	if size := tree.Len(); size != 2 {
		t.Errorf("expected size 2 after removing non-existing key, got %d", size)
	}
}

func TestTree_Get(t *testing.T) {
	tree := New[int, string]()

	// Insert some elements
	tree.Put(10, "a")
	tree.Put(20, "b")
	tree.Put(5, "c")

	// Test 1: Get an existing node
	value, found := tree.Get(10)
	if !found || value != "a" {
		t.Errorf("expected value 'a' for key 10, got %v", value)
	}

	// Test 2: Get a non-existing node
	if _, found = tree.Get(999); found {
		t.Error("expected key 999 not to be found")
	}

	// Test 3: Get an updated value
	tree.Put(10, "updated")
	value, found = tree.Get(10)
	if !found || value != "updated" {
		t.Errorf("expected updated value 'updated' for key 10, got %v", value)
	}

	// Test 4: Get a value for a different existing key
	value, found = tree.Get(20)
	if !found || value != "b" {
		t.Errorf("expected value 'b' for key 20, got %v", value)
	}

	// Test 5: Get a value for a non-existing key (zero-value check)
	if value, found = tree.Get(100); found || value != "" {
		t.Errorf("expected zero value for key 100, got %v", value)
	}
}

func TestTree_Contains(t *testing.T) {
	tree := New[int, string]()

	// Test 1: Check Contains on an empty tree
	if contains := tree.Contains(10); contains {
		t.Error("expected tree to not contain 10 on empty tree")
	}

	// Test 2: Insert an element and check if Contains works
	tree.Put(10, "a")
	if contains := tree.Contains(10); !contains {
		t.Error("expected tree to contain 10 after insertion")
	}

	// Test 3: Check Contains for a key that does not exist
	if contains := tree.Contains(20); contains {
		t.Error("expected tree to not contain 20, but it does")
	}

	// Test 4: Insert more elements and check Contains
	tree.Put(20, "b")
	tree.Put(30, "c")
	if contains := tree.Contains(20); !contains {
		t.Error("expected tree to contain 20 after insertion")
	}
	if contains := tree.Contains(30); !contains {
		t.Error("expected tree to contain 30 after insertion")
	}

	// Test 5: Remove an element and check Contains
	tree.Remove(10)
	if contains := tree.Contains(10); contains {
		t.Error("expected tree to not contain 10 after removal")
	}

	// Test 6: Check Contains for a non-existing key after removals
	if contains := tree.Contains(999); contains {
		t.Error("expected tree to not contain 999")
	}
}

func TestTree_Len(t *testing.T) {
	tree := New[int, string]()

	// Test 1: Len of an empty tree
	if size := tree.Len(); size != 0 {
		t.Errorf("expected size 0 for empty tree, got %d", size)
	}

	// Test 2: Insert some elements and check size
	tree.Put(10, "a")
	tree.Put(20, "b")
	tree.Put(5, "c")
	if size := tree.Len(); size != 3 {
		t.Errorf("expected size 3, got %d", size)
	}

	// Test 3: Remove an element and check size
	tree.Remove(20)
	if size := tree.Len(); size != 2 {
		t.Errorf("expected size 2 after removal, got %d", size)
	}

	// Test 4: Remove another element and check size
	tree.Remove(5)
	if size := tree.Len(); size != 1 {
		t.Errorf("expected size 1 after removal, got %d", size)
	}

	// Test 5: Remove the last element and check size
	tree.Remove(10)
	if size := tree.Len(); size != 0 {
		t.Errorf("expected size 0 after all removals, got %d", size)
	}
}

func TestTree_IsEmpty(t *testing.T) {
	tree := New[int, string]()

	// Test 1: IsEmpty on an empty tree
	if empty := tree.IsEmpty(); !empty {
		t.Error("expected tree to be empty")
	}

	// Test 2: Insert an element and check if tree is empty
	tree.Put(10, "a")
	if empty := tree.IsEmpty(); empty {
		t.Error("expected tree to not be empty after insertion")
	}

	// Test 3: Remove the element and check if tree is empty
	tree.Remove(10)
	if empty := tree.IsEmpty(); !empty {
		t.Error("expected tree to be empty after removal")
	}

	// Test 4: Insert multiple elements and check if tree is empty
	tree.Put(20, "b")
	tree.Put(30, "c")
	if empty := tree.IsEmpty(); empty {
		t.Error("expected tree to not be empty after multiple insertions")
	}

	// Test 5: Remove all elements and check if tree is empty
	tree.Remove(20)
	tree.Remove(30)
	if empty := tree.IsEmpty(); !empty {
		t.Error("expected tree to be empty after all removals")
	}
}

func TestTree_Traversals(t *testing.T) {
	tree := New[int, string]()

	// Helper function to compare slices
	equal := func(a, b []interface{}) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	// Insert some elements
	tree.Put(10, "a")
	tree.Put(20, "b")
	tree.Put(5, "c")
	tree.Put(15, "d")
	tree.Put(30, "e")

	// Test InOrder traversal
	inOrder := tree.InOrder()
	expectedInOrder := []interface{}{"c", "a", "d", "b", "e"}
	if !equal(inOrder, expectedInOrder) {
		t.Errorf("expected InOrder %v, got %v", expectedInOrder, inOrder)
	}

	// Test PreOrder traversal
	preOrder := tree.PreOrder()
	expectedPreOrder := []interface{}{"a", "c", "b", "d", "e"}
	if !equal(preOrder, expectedPreOrder) {
		t.Errorf("expected PreOrder %v, got %v", expectedPreOrder, preOrder)
	}

	// Test PostOrder traversal
	postOrder := tree.PostOrder()
	expectedPostOrder := []interface{}{"c", "d", "e", "b", "a"}
	if !equal(postOrder, expectedPostOrder) {
		t.Errorf("expected PostOrder %v, got %v", expectedPostOrder, postOrder)
	}
}

func TestTree_String(t *testing.T) {
	tree := New[int, string]()

	// Insert elements
	tree.Put(10, "a")
	tree.Put(5, "b")
	tree.Put(20, "c")
	tree.Put(15, "d")
	tree.Put(25, "e")

	// Test String output
	expected := "Level 0: a\nLevel 1: b c\nLevel 2: d e"
	if got := tree.String(); got != expected {
		t.Errorf("expected String output:\n%s\nbut got:\n%s", expected, got)
	}
}
