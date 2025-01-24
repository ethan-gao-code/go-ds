// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ethan-gao-code/go-ds/sets"
)

func main() {
	// Create a new set
	setA := sets.New(1, 2, 3, 4)
	fmt.Println("Set A:", setA)

	// Add elements to the set
	setA.Add(5, 6)
	fmt.Println("Set A after adding elements:", setA)

	// Remove an element from the set
	setA.Remove(2)
	fmt.Println("Set A after removing element 2:", setA)

	// Check if an element is in the set
	containsThree := setA.Contains(3)
	fmt.Println("Does Set A contain 3?", containsThree)

	// Get the size of the set
	setSize := setA.Size()
	fmt.Println("Size of Set A:", setSize)

	// Create another set
	setB := sets.New(3, 4, 5, 6, 7)
	fmt.Println("Set B:", setB)

	// Union of Set A and Set B
	unionSet := setA.Union(setB)
	fmt.Println("Union of Set A and Set B:", unionSet)

	// Intersection of Set A and Set B
	intersectionSet := setA.Intersection(setB)
	fmt.Println("Intersection of Set A and Set B:", intersectionSet)

	// Difference of Set A and Set B
	differenceSet := setA.Difference(setB)
	fmt.Println("Difference of Set A and Set B:", differenceSet)

	// Check if Set A is a subset of Set B
	isSubset := setA.Subset(setB)
	fmt.Println("Is Set A a subset of Set B?", isSubset)

	// Clear Set A
	setA.Clear()
	fmt.Println("Set A after clearing:", setA)

	// Check if Set A is empty
	isEmpty := setA.IsEmpty()
	fmt.Println("Is Set A empty?", isEmpty)
}
