// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package bloomfilters

// BloomFilters defines the structure of the Bloom Filter
type BloomFilters struct {
	bitmap     []byte // Underlying bitmap for storing bits
	bitmapSize uint64 // Size of the bitmap in bits
	hashCount  uint64 // Number of hash functions
}
