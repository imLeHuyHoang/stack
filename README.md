# Stack

The **Stack** library provides a generic **LIFO (Last In, First Out)** stack data structure for Go.

## Stack Overview

A Stack is a data structure that follows the **LIFO** principle (Last In, First Out):
the most recently added element is the first one to be removed.

![alt text](image.png)

## Installation

```bash
go get github.com/imLeHuyHoang/stack
```

## Stack Methods

### `NewStack[T any]() *Stack[T]`

Creates a new stack with a generic type.

```go
s := stack.NewStack[int]()        // Stack of integers
s2 := stack.NewStack[string]()    // Stack of strings
```

---

### `Push(v T)`

Adds an element to the top of the stack.

```go
s := stack.NewStack[int]()
s.Push(10)  // Stack: [10]
s.Push(20)  // Stack: [10, 20]
s.Push(30)  // Stack: [10, 20, 30]
```

---

### `Pop() (v T, ok bool)`

Removes and returns the top element of the stack.

* Returns `ok = true` if successful
* Returns `ok = false` if the stack is empty (value `v` will be the zero value of type `T`)

```go
v, ok := s.Pop()  // v = 30, ok = true, Stack: [10, 20]
if ok {
    fmt.Println(v)  // Output: 30
}

// When the stack is empty
empty := stack.NewStack[int]()
v, ok := empty.Pop()  // v = 0 (zero value), ok = false
```

---

### `Peek() (v T, ok bool)`

Returns the top element of the stack **without removing it**.

```go
v, ok := s.Peek()  // v = 20, ok = true, Stack remains: [10, 20]
if ok {
    fmt.Println(v)  // Output: 20
}
```

---

### `Len() int`

Returns the number of elements in the stack.

```go
count := s.Len()  // count = 2
```

---

### `IsEmpty() bool`

Checks whether the stack is empty.

```go
if s.IsEmpty() {
    fmt.Println("Stack is empty")
} else {
    fmt.Println("Stack has", s.Len(), "elements")
}
```

---

## Full Example

```go
package main

import (
	"fmt"
	"github.com/imLeHuyHoang/stack"
)

func main() {
	// Create a stack of integers
	s := stack.NewStack[int]()

	// Check if the stack is empty
	fmt.Println("Stack is empty:", s.IsEmpty())  // true

	// Push elements
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Number of elements:", s.Len())  // 3

	// Peek top element
	top, ok := s.Peek()
	if ok {
		fmt.Println("Top element:", top)  // 30
	}

	// Pop elements (LIFO)
	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Println("Pop:", v)  // 30, 20, 10
	}

	// Try to pop when stack is empty
	v, ok := s.Pop()
	if !ok {
		fmt.Println("Stack is empty, cannot pop")
	}
}
```
