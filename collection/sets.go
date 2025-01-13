package collection

// Set interface defines the behavior of a set.
type Set interface {
	Add(values ...interface{})             // Add elements to the set.
	AddAll(values []interface{})           // Add multiple elements to the set.
	Remove(values ...interface{})          // Remove elements from the set.
	RemoveAll(values []interface{})        // Remove multiple elements from the set.
	Contains(values ...interface{}) bool   // Check if the set contains all the given elements.
	ContainsAll(values []interface{}) bool // Check if the set contains all the elements in the slice.
	Size() int                             // Return the number of elements in the set.
	IsEmpty() bool                         // Check if the set is empty.
	Clear()                                // Clear all elements from the set.
	Values() []interface{}                 // Return all elements in the set as a slice.
}
