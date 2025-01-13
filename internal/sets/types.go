package sets

// Sets defines a collection type implemented using a hash map.
type Sets struct {
	items map[interface{}]struct{}
}
