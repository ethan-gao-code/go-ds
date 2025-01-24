// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ethan-gao-code/go-ds/bloomfilters"
)

// ArrayListExample to demonstrate basic usage of ArrayList
func main() {
	// Create a new BloomFilter with a desired false positive rate of 0.01 and expected 1000 items
	bf := bloomfilters.New(0.01, 1000)

	// Add some elements to the Bloom Filter
	bf.Add("apple")
	bf.Add("banana")
	bf.Add("cherry")

	// Check if some elements are in the Bloom Filter
	fmt.Println("Is 'apple' in the Bloom Filter?", bf.Contains("apple"))   // Expected: true
	fmt.Println("Is 'banana' in the Bloom Filter?", bf.Contains("banana")) // Expected: true
	fmt.Println("Is 'cherry' in the Bloom Filter?", bf.Contains("cherry")) // Expected: true
	fmt.Println("Is 'orange' in the Bloom Filter?", bf.Contains("orange")) // Expected: false (or maybe true due to false positive rate)

	// Check the size of the Bloom Filter (number of bits set to 1)
	fmt.Println("Number of bits set to 1:", bf.Size())

	// Get the number of hash functions used by the Bloom Filter
	fmt.Println("Number of hash functions:", bf.HashCount())

	// Reset the Bloom Filter and check its status
	bf.Reset()
	fmt.Println("After reset:")
	fmt.Println("Is 'apple' in the Bloom Filter?", bf.Contains("apple")) // Expected: false
	fmt.Println("Number of bits set to 1:", bf.Size())

	// Print the string representation of the Bloom Filter
	fmt.Println("Bloom Filter:", bf.String())

	// Print the set bits in the bitmap (for debugging purposes)
	fmt.Println("Set bits indices:", bf.Values())
}
