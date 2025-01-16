// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package bloomfilters

// BloomFilter defines the behavior of a Bloom Filter.
type BloomFilter interface {
	Add(item string)           // Add an item to the Bloom Filter.
	Contains(item string) bool // Contains checks if an item might exist in the Bloom Filter.
	Size() uint64              // Size returns the size of the Bloom Filter (number of bits).
	HashCount() uint64         // HashCount returns the number of hash functions used in the Bloom Filter.
	Reset()                    // Reset clears all bits in the Bloom Filter.
	Values() []uint64          // Values returns all the indices that are set in the bitmap (for debugging or analysis).
	String() string            // String provides a string representation of the Bloom Filter (e.g., a summary).
}
