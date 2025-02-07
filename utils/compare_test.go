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

		// Comparing bool
		{true, true, 0},   // true == true
		{true, false, 1},  // true > false
		{false, true, -1}, // false < true

		// Comparing int
		{5, 5, 0},  // 5 == 5
		{5, 2, 1},  // 5 > 2
		{2, 5, -1}, // 2 < 5

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

func TestEqualObjects(t *testing.T) {
	tests := []struct {
		obj1     interface{}
		obj2     interface{}
		expected bool
	}{
		// Comparing strings
		{"apple", "banana", false}, // "apple" < "banana"
		{"banana", "apple", false}, // "banana" > "apple"
		{"apple", "apple", true},   // "apple" == "apple"

		// Comparing floats
		{10.5, 10.7, false}, // 10.5 < 10.7
		{10.7, 10.5, false}, // 10.7 > 10.5
		{10.5, 10.5, true},  // 10.5 == 10.5
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := EqualObjects(test.obj1, test.obj2)
			if result != test.expected {
				t.Errorf("EqualObjects(%v, %v) = %t; expected %t", test.obj1, test.obj2, result, test.expected)
			}
		})
	}
}
