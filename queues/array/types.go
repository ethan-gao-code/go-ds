// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package array

// Queues defines a queue data structure
type Queues struct {
	items []interface{}
}

// GenericQueue defines a queue data structure with a generic type
type GenericQueue[T any] struct {
	items []T
}
