// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package utils

import (
	"fmt"
)

// CompareObjects compares two objects. This is a simple version that supports string comparison.
func CompareObjects(obj1, obj2 interface{}) int {
	switch v1 := obj1.(type) {
	case string:
		if v2, ok := obj2.(string); ok {
			return compareStrings(v1, v2)
		}
	case float64:
		if v2, ok := obj2.(float64); ok {
			return compareFloats(v1, v2)
		}
	case bool:
		if v2, ok := obj2.(bool); ok {
			if v1 == v2 {
				return 0
			} else if v1 {
				return 1
			} else {
				return -1
			}
		}
		if v2, ok := obj2.(int); ok {
			if v1 && v2 == 1 {
				return 0
			}
			if v1 {
				return 1
			}
			return -1
		}
	case int:
		if v2, ok := obj2.(int); ok {
			return compareInts(v1, v2)
		}
		if v2, ok := obj2.(bool); ok {
			if v1 == 1 && v2 {
				return 0
			}
			if v1 == 1 {
				return 1
			}
			return -1
		}
	}
	// Default case
	return compareStrings(fmt.Sprintf("%v", obj1), fmt.Sprintf("%v", obj2))
}

func compareStrings(str1, str2 string) int {
	if str1 < str2 {
		return -1
	} else if str1 > str2 {
		return 1
	}
	return 0
}

func compareFloats(f1, f2 float64) int {
	if f1 < f2 {
		return -1
	} else if f1 > f2 {
		return 1
	}
	return 0
}

func compareInts(i1, i2 int) int {
	if i1 < i2 {
		return -1
	} else if i1 > i2 {
		return 1
	}
	return 0
}

func EqualObjects(obj1, obj2 interface{}) bool {
	return CompareObjects(obj1, obj2) == 0
}
