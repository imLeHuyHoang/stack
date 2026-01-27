package stack

// Build stack based on slice
type Stack[T any] struct {
	elements []T
}

// Create a new stack
func NewStack[T any]() *Stack[T] {
	elements := make([]T, 0)
	return &Stack[T]{elements: elements}
}

// Method 1: Push
// Use to add an element to the top of the stack
func (stack *Stack[T]) Push(value T) {
	stack.elements = append(stack.elements, value)
}

// Method 2: Pop
// Use to remove and return the top element of the stack

func (stack *Stack[T]) Pop() (T, bool) {
	var value T
	if len(stack.elements) == 0 {
		return value, false
	}

	lastIndex := len(stack.elements) - 1
	value = stack.elements[lastIndex]
	stack.elements = stack.elements[:lastIndex]
	return value, true
}

// Method 3: Peek
// Use to return the top element of the stack without removing it

func (stack *Stack[T]) Peek() (T, bool) {
	var value T
	if len(stack.elements) == 0 {
		return value, false
	}
	lastIndex := len(stack.elements) - 1
	value = stack.elements[lastIndex]
	return value, true
}

// Method 4: IsEmpty
// Use to check if the stack is empty
func (stack *Stack[T]) IsEmpty() bool {
	length := len(stack.elements)
	return length == 0
}

// Method 5: Size
// Use to return the number of elements in the stack

func (stack *Stack[T]) Size() int {
	return len(stack.elements)
}
