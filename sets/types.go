// Copyright (c) 2025 EthanGao
// This file is part of a project licensed under the MIT License.
// License that can be found in the LICENSE file.

package sets

// Sets defines a collection type implemented using a hash map.
type Sets struct {
	items map[interface{}]struct{}
}
