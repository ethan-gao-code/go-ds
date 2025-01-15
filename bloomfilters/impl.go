// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package bloomfilters

import (
	"fmt"
	"hash/fnv"
	"math"
)

// New creates a new Bloom Filter
// falsePositiveRate: desired false positive rate (default 0.01)
// expectedItemsCount: the expected number of items to be added to the Bloom Filter (default 1000)
func New(falsePositiveRate float64, expectedItemsCount int) *BloomFilters {
	// Use default values if inputs are invalid
	if falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		falsePositiveRate = DefaultFalsePositiveRate
	}
	if expectedItemsCount <= 0 {
		expectedItemsCount = DefaultExpectedItemsCount
	}

	// Calculate optimal size of the bitmap (m) and hash function count (k)
	size := optimalBitmapSize(expectedItemsCount, falsePositiveRate)
	hashCount := optimalHashCount(expectedItemsCount, size)

	return &BloomFilters{
		bitmap:     make([]byte, (size+7)/8), // size in bytes
		bitmapSize: size,
		hashCount:  hashCount,
	}
}

// Add adds an item to the Bloom Filter by setting the corresponding bits in the bitmap.
func (bf *BloomFilters) Add(item string) {
	for i := uint64(0); i < bf.hashCount; i++ {
		hash := bf.hash(item, i) // Use hash32
		index := hash % bf.bitmapSize
		bf.setBit(index)
	}
}

// Contains checks if an item might exist in the Bloom Filter.
// It returns true if the item might exist (may have a false positive), false if the item does not exist.
func (bf *BloomFilters) Contains(item string) bool {
	for i := uint64(0); i < bf.hashCount; i++ {
		hash := bf.hash(item, i) // Use hash32
		index := hash % bf.bitmapSize
		if !bf.getBit(index) {
			return false
		}
	}
	return true
}

// Size returns the number of bits set to 1 in the bitmap.
func (bf *BloomFilters) Size() uint64 {
	var count uint64
	// Iterate through each byte in the bitmap
	for _, byteVal := range bf.bitmap {
		// Count the number of 1 bit in this byte
		for i := 0; i < 8; i++ {
			if (byteVal & (1 << i)) != 0 {
				count++
			}
		}
	}
	return count
}

// HashCount returns the number of hash functions used by the Bloom Filter.
func (bf *BloomFilters) HashCount() uint64 {
	return bf.hashCount
}

// Reset clears all bits in the Bloom Filter bitmap.
func (bf *BloomFilters) Reset() {
	// Reset all bits to 0
	for i := range bf.bitmap {
		bf.bitmap[i] = 0
	}
}

// Values returns the indices of all bits set in the bitmap (for debugging or analysis purposes).
func (bf *BloomFilters) Values() []uint64 {
	var indices []uint64
	for i, b := range bf.bitmap {
		if b != 0 {
			// Convert byte index to individual bit positions (not the best approach but simple)
			for j := uint64(0); j < 8; j++ {
				if (b & (1 << j)) != 0 {
					indices = append(indices, uint64(i)*8+j)
				}
			}
		}
	}
	return indices
}

// String provides a string representation of the Bloom Filter.
func (bf *BloomFilters) String() string {
	return fmt.Sprintf("BloomFilter {Size: %d, HashCount: %d}", bf.bitmapSize, bf.hashCount)
}

// setBit sets a bit at the specified index in the bitmap.
func (bf *BloomFilters) setBit(index uint64) {
	byteIndex := index / 8
	bitIndex := index % 8
	bf.bitmap[byteIndex] |= 1 << bitIndex
}

// getBit checks if a bit at the specified index is set in the bitmap.
func (bf *BloomFilters) getBit(index uint64) bool {
	byteIndex := index / 8
	bitIndex := index % 8
	return (bf.bitmap[byteIndex] & (1 << bitIndex)) != 0
}

// hash generates a hash value using FNV-1a (32-bit) with a seed.
func (bf *BloomFilters) hash(item string, seed uint64) uint64 {
	// hashFuncs Using a 32-bit hash function to reduce memory usage and improve performance.
	// Explanation of the choice:
	// - Memory Efficiency: A 32-bit hash occupies 4 bytes, while a 64-bit hash would occupy 8 bytes.
	//   Using 32-bit hashes reduces memory consumption, especially when multiple hash functions are used.
	// - Suitable for Bloom Filter Needs: 32-bit hashes provide enough entropy to ensure the proper distribution
	//   of hash values across the bit array, preventing collisions effectively in most cases.
	// - Calculation Speed: 32-bit hashes are computationally faster compared to 64-bit hashes,
	//   as they require fewer bit operations.
	// - Acceptable Trade-off: The use of 32-bit hashes strikes a balance between accuracy and performance,
	//   without introducing significant risk of hash collisions.
	h := fnv.New32a()
	_, _ = h.Write([]byte(item))
	_, _ = h.Write([]byte{byte(seed)})
	return uint64(h.Sum32())
}

// optimalBitmapSize calculates the optimal size of the bitmap (m) given the expected number of items (n) and the false positive rate (p).
func optimalBitmapSize(capacity int, falsePositiveRate float64) uint64 {
	// Formula for optimal bitmap size (m)
	return uint64(math.Ceil(-float64(capacity) * math.Log(falsePositiveRate) / (math.Ln2 * math.Ln2)))
}

// optimalHashCount calculates the optimal number of hash functions (k) given the expected number of items (n) and bitmap size (m).
func optimalHashCount(capacity int, bitmapSize uint64) uint64 {
	// Formula for optimal hash count (k)
	return uint64(math.Ceil(math.Ln2 * float64(bitmapSize) / float64(capacity)))
}
