package array

// Stack interface defines the behavior of a stack.
type Stack interface {
	Push(values ...interface{}) // Add elements to the stack.
	Pop() interface{}           // Remove and return the top element from the stack.
	Peek() interface{}          // Return the top element without removing it.
	IsEmpty() bool              // Check if the stack is empty.
	Size() int                  // Return the number of elements in the stack.
	Clear()                     // Remove all elements from the stack.
	Values() []interface{}      // Return all elements in the stack as a slice.
	String() string
}
