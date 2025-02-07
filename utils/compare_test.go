// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package utils

import "testing"

func TestCompareObjects(t *testing.T) {
	tests := []struct {
		obj1     interface{}
		obj2     interface{}
		expected int
	}{
		// Comparing strings
		{"apple", "banana", -1}, // "apple" < "banana"
		{"banana", "apple", 1},  // "banana" > "apple"
		{"apple", "apple", 0},   // "apple" == "apple"

		// Comparing floats
		{10.5, 10.7, -1}, // 10.5 < 10.7
		{10.7, 10.5, 1},  // 10.7 > 10.5
		{10.5, 10.5, 0},  // 10.5 == 10.5

		// Comparing string and float (non-equal types)
		{"10.5", 10.5, 0}, // "10.5" == 10.5 when converted to string

		// Comparing nil
		{nil, nil, 0},       // nil == nil
		{"apple", nil, 1},   // non-nil > nil
		{nil, "banana", -1}, // nil < non-nil

		// Default comparison for mixed types (using string format)
		{10, "10", 0}, // 10 == "10" when converted to string
		{true, 1, 0},  // true == 1 when converted to string
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := CompareObjects(test.obj1, test.obj2)
			if result != test.expected {
				t.Errorf("CompareObjects(%v, %v) = %d; expected %d", test.obj1, test.obj2, result, test.expected)
			}
		})
	}
}
