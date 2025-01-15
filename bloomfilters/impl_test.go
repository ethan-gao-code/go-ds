// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package bloomfilters

import (
	"testing"
)

// TestBloomFilter tests the basic operations of the Bloom Filter
func TestBloomFilter(t *testing.T) {
	// Initialize a new Bloom Filter with a false positive rate of 1% and expected 1000 elements
	bf := New(0.01, 1000)

	// Test Add and Contains for an existing item
	item1 := "testItem1"
	bf.Add(item1)
	if !bf.Contains(item1) {
		t.Errorf("Expected BloomFilter to contain %s after adding it", item1)
	}

	// Test Contains for a non-existent item (should return false)
	item2 := "testItem2"
	if bf.Contains(item2) {
		t.Errorf("Expected BloomFilter to NOT contain %s as it was not added", item2)
	}

	// Test Add and Contains for another item
	item3 := "testItem3"
	bf.Add(item3)
	if !bf.Contains(item3) {
		t.Errorf("Expected BloomFilter to contain %s after adding it", item3)
	}

	// Test Reset method to clear all elements
	bf.Reset()
	if bf.Contains(item1) {
		t.Errorf("Expected BloomFilter to NOT contain %s after resetting", item1)
	}
	if bf.Contains(item3) {
		t.Errorf("Expected BloomFilter to NOT contain %s after resetting", item3)
	}

	// Test Size and HashCount methods
	if size := bf.Size(); size != uint64(0) {
		t.Errorf("Expected BloomFilter size to be %d, got %d", 0, size)
	}

	if hashCount := bf.HashCount(); hashCount != uint64(7) {
		t.Errorf("Expected BloomFilter hash count to be %d, got %d", 9, hashCount)
	}
}

// TestBloomFilterWithDifferentParams tests the Bloom Filter with different initialization parameters
func TestBloomFilterWithDifferentParams(t *testing.T) {
	// Initialize Bloom Filter with a lower false positive rate
	bf := New(0.001, 1000)
	if bf.HashCount() == 0 {
		t.Errorf("Expected BloomFilter hash count to be greater than 0 with a false positive rate of 0.001")
	}

	// Initialize Bloom Filter with a larger expected element count
	bf2 := New(0.01, 5000)
	if bf2.HashCount() == 0 {
		t.Errorf("Expected BloomFilter hash count to be greater than 0 with 5000 expected elements")
	}
}

// TestBloomFilterValues tests the Values method which returns the bit positions that are set
func TestBloomFilterValues(t *testing.T) {
	bf := New(0.01, 1000)

	// Add items to the Bloom Filter
	item1 := "item1"
	item2 := "item2"
	bf.Add(item1)
	bf.Add(item2)

	// Ensure that Values returns a non-empty slice of set bit positions
	indices := bf.Values()
	if len(indices) == 0 {
		t.Errorf("Expected BloomFilter.Values() to return bit positions, but got an empty slice")
	}

	// Verify that the bit positions are related to the items added
	// This is a bit tricky to check directly, so we are mainly verifying that the method works
	if len(indices) < 2 {
		t.Errorf("Expected BloomFilter.Values() to return at least 2 bit positions, got %d", len(indices))
	}
}

// TestBloomFilterEdgeCases tests edge cases for Bloom Filter
func TestBloomFilterEdgeCases(t *testing.T) {
	// Test invalid false positive rate (out of bounds)
	bf := New(0.0, 1000) // Invalid false positive rate
	if bf == nil {
		t.Errorf("Expected BloomFilter to be created with a valid false positive rate")
	}

	// Test Reset on an empty Bloom Filter
	bf.Reset()
	if size := bf.Size(); size != 0 {
		t.Errorf("Expected BloomFilter size to be 0 after Reset, got %d", size)
	}

	// Test invalid expected elements count (0)
	bf2 := New(0.01, 0) // Invalid expected element count
	if bf2 == nil {
		t.Errorf("Expected BloomFilter to be created with a valid expected item count")
	}
}
