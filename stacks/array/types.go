// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

// Stacks defines a stack data structure without generics.
type Stacks struct {
	items []interface{}
}

// GenericStacks defines a stack data structure with a generic type.
type GenericStacks[T any] struct {
	items []T
}
