# stack

A simple, efficient, and type-safe generic stack (LIFO - Last In, First Out) implementation for Go.

## Features

- **Type-safe**: Built with Go generics (requires Go 1.18+)
- **Simple API**: Easy-to-use interface with intuitive method names
- **Efficient**: Backed by Go's dynamic slices for optimal performance
- **Zero dependencies**: No external dependencies required
- **Well-tested**: Comprehensive test coverage
- **Production-ready**: Suitable for real-world applications

## Installation

```bash
go get github.com/imLeHuyHoang/stack
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/imLeHuyHoang/stack"
)

func main() {
    // Create a new stack for integers
    s := stack.New[int]()
    
    // Push elements onto the stack
    s.Push(10)
    s.Push(20)
    s.Push(30)
    
    // Peek at the top element without removing it
    if top, ok := s.Peek(); ok {
        fmt.Println("Top element:", top) // Output: Top element: 30
    }
    
    // Pop elements from the stack
    if val, ok := s.Pop(); ok {
        fmt.Println("Popped:", val) // Output: Popped: 30
    }
    
    // Check the size
    fmt.Println("Size:", s.Size()) // Output: Size: 2
    
    // Check if empty
    fmt.Println("Is empty:", s.IsEmpty()) // Output: Is empty: false
}
```

## Usage Examples

### Working with Different Types

The stack is fully generic and works with any type:

```go
// String stack
stringStack := stack.New[string]()
stringStack.Push("hello")
stringStack.Push("world")

// Struct stack
type Person struct {
    Name string
    Age  int
}

personStack := stack.New[Person]()
personStack.Push(Person{Name: "Alice", Age: 30})
personStack.Push(Person{Name: "Bob", Age: 25})

// Pointer stack
ptrStack := stack.New[*Person]()
ptrStack.Push(&Person{Name: "Charlie", Age: 35})
```

### Iterating Through Stack Elements

```go
s := stack.New[int]()
s.Push(1)
s.Push(2)
s.Push(3)

// Pop all elements
for !s.IsEmpty() {
    if val, ok := s.Pop(); ok {
        fmt.Println(val)
    }
}
// Output: 3, 2, 1 (LIFO order)
```

### Safe Empty Stack Operations

```go
s := stack.New[int]()

// Pop from empty stack returns zero value and false
val, ok := s.Pop()
if !ok {
    fmt.Println("Stack is empty")
}

// Peek on empty stack also returns zero value and false
top, ok := s.Peek()
if !ok {
    fmt.Println("Nothing to peek")
}
```

### Clearing the Stack

```go
s := stack.New[int]()
s.Push(1)
s.Push(2)
s.Push(3)

s.Clear()
fmt.Println("Size after clear:", s.Size()) // Output: Size after clear: 0
```

## API Reference

### Creating a Stack

```go
func New[T any]() *Stack[T]
```
Creates and returns a new empty stack that can hold elements of type `T`.

### Methods

#### Push
```go
func (s *Stack[T]) Push(value T)
```
Adds an element to the top of the stack.

**Example:**
```go
s.Push(42)
```

#### Pop
```go
func (s *Stack[T]) Pop() (T, bool)
```
Removes and returns the element at the top of the stack. Returns the zero value of type `T` and `false` if the stack is empty.

**Example:**
```go
if val, ok := s.Pop(); ok {
    fmt.Println("Popped:", val)
} else {
    fmt.Println("Stack is empty")
}
```

#### Peek
```go
func (s *Stack[T]) Peek() (T, bool)
```
Returns the element at the top of the stack without removing it. Returns the zero value of type `T` and `false` if the stack is empty.

**Example:**
```go
if top, ok := s.Peek(); ok {
    fmt.Println("Top element:", top)
}
```

#### Size
```go
func (s *Stack[T]) Size() int
```
Returns the number of elements currently in the stack.

**Example:**
```go
fmt.Println("Stack has", s.Size(), "elements")
```

#### IsEmpty
```go
func (s *Stack[T]) IsEmpty() bool
```
Returns `true` if the stack contains no elements, `false` otherwise.

**Example:**
```go
if s.IsEmpty() {
    fmt.Println("Stack is empty")
}
```

#### Clear
```go
func (s *Stack[T]) Clear()
```
Removes all elements from the stack, making it empty.

**Example:**
```go
s.Clear()
```

## Performance Characteristics

| Operation | Time Complexity | Space Complexity |
|-----------|----------------|------------------|
| Push      | O(1) amortized* | O(1)            |
| Pop       | O(1)           | O(1)            |
| Peek      | O(1)           | O(1)            |
| Size      | O(1)           | O(1)            |
| IsEmpty   | O(1)           | O(1)            |
| Clear     | O(1)           | O(1)            |

\* **Note on Push complexity**: Push operations are O(1) amortized. While most push operations are constant time, occasional slice reallocations may require O(n) time when the underlying array grows. However, due to Go's doubling strategy for slice growth, the amortized cost remains O(1).

## Thread Safety

⚠️ **Note**: This stack implementation is **not thread-safe**. If you need to use it in concurrent scenarios, you should protect it with a mutex:

```go
type SafeStack[T any] struct {
    stack *stack.Stack[T]
    mu    sync.Mutex
}

func (s *SafeStack[T]) Push(value T) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.stack.Push(value)
}

func (s *SafeStack[T]) Pop() (T, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.stack.Pop()
}
```

## Common Use Cases

- **Expression evaluation**: Parsing and evaluating mathematical expressions
- **Undo/Redo functionality**: Implementing undo/redo in applications
- **Backtracking algorithms**: Maze solving, graph traversal (DFS)
- **Function call management**: Understanding recursion and call stacks
- **Browser history**: Forward/backward navigation
- **Text editor operations**: Matching brackets, parentheses validation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Testing

Run the tests with:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -cover
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Le Huy Hoang** - [@imLeHuyHoang](https://github.com/imLeHuyHoang)

## Acknowledgments

- Inspired by classic data structure implementations
- Built with Go's powerful generics feature
- Thanks to the Go community for best practices and feedback

## Support

If you find this package helpful, please consider giving it a ⭐️ on GitHub!