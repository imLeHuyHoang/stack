# Stack

Thư viện Stack cung cấp cấu trúc dữ liệu ngăn xếp (LIFO - Last In First Out) với kiểu generic trong Go.

## Tổng quan về Stack

Stack là một cấu trúc dữ liệu hoạt động theo nguyên tắc **LIFO** (Last In First Out - vào sau ra trước). 
![alt text](image.png)

## Cài đặt

```bash
go get github.com/imLeHuyHoang/stack
```

## Method sử dụng với Stack

### `NewStack[T any]() *Stack[T]`
Tạo một stack mới với kiểu dữ liệu tùy chỉnh.

```go
s := stack.NewStack[int]()        // Stack chứa số nguyên
s2 := stack.NewStack[string]()    // Stack chứa chuỗi
```

### `Push(v T)`
Thêm một phần tử vào stack.

```go
s := stack.NewStack[int]()
s.Push(10)  // Stack: [10]
s.Push(20)  // Stack: [10, 20]
s.Push(30)  // Stack: [10, 20, 30]
```

### `Pop() (v T, ok bool)`
Chọn phần tử trên cùng của stack và xóa. Trả về giá trị `v` và `ok=true` nếu thành công, `ok=false` nếu stack rỗng.

```go
v, ok := s.Pop()  // v = 30, ok = true, Stack: [10, 20]
if ok {
    fmt.Println(v)  // Output: 30
}

// Nếu stack rỗng
empty := stack.NewStack[int]()
v, ok := empty.Pop()  // v = 0 (giá trị zero), ok = false
```

### `Peek() (v T, ok bool)`
Xem phần tử ở đỉnh stack và không thực hiện bất cứ thao tác nào với phần tử. Trả về giá trị và `ok=true` nếu thành công, `ok=false` nếu stack rỗng.

```go
v, ok := s.Peek()  // v = 20, ok = true, Stack vẫn: [10, 20]
if ok {
    fmt.Println(v)  // Output: 20
}
```

### `Len() int`
Trả về số lượng phần tử trong stack.

```go
count := s.Len()  // count = 2
```

### `IsEmpty() bool`
Kiểm tra stack có rỗng hay không.

```go
if s.IsEmpty() {
    fmt.Println("Stack rỗng")
} else {
    fmt.Println("Stack có", s.Len(), "phần tử")
}
```

## Ví dụ đầy đủ

```go
package main

import (
	"fmt"
	"github.com/imLeHuyHoang/stack"
)

func main() {
	// Tạo stack chứa số nguyên
	s := stack.NewStack[int]()
	
	// Kiểm tra stack rỗng
	fmt.Println("Stack rỗng:", s.IsEmpty())  // true
	
	// Thêm phần tử
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Số phần tử:", s.Len())  // 3
	
	// Xem phần tử đỉnh
	top, ok := s.Peek()
	if ok {
		fmt.Println("Phần tử đỉnh:", top)  // 30
	}
	
	// Lấy các phần tử ra
	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Println("Pop:", v)  // 30, 20, 10
	}
	
	// Thử pop khi stack rỗng
	v, ok := s.Pop()
	if !ok {
		fmt.Println("Stack đã rỗng, không thể pop")
	}
}
```