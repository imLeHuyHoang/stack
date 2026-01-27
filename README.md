# Stack

A simple generic stack (LIFO) implementation for Go.

## Installation

```bash
go get github.com/imLeHuyHoang/stack
```
## Usage

```go
package main

import (
	"fmt"
	"github.com/imLeHuyHoang/stack"
)

func main() {
	s := stack.NewStack[int]()
	s.Push(10)
	s.Push(20)

	v, _ := s.Pop()
	fmt.Println(v) // Output: 20
    fmt.Println(s.Peek()) // Output: 10
}
```