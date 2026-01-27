package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int]()

	if s.IsEmpty() != true {
		t.Fatal("Expected stack to be empty")
	}

	// Test method Push
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Test method Size
	if s.Size() != 3 {
		t.Fatalf("Expected stack size to be 3, got %d", s.Size())
	}

	// Test method Peek
	value, ok := s.Peek()
	if !ok || value != 30 {
		t.Fatalf("Expected top element to be 30, got %d", value)
	}

	// Test method Pop
	value, ok = s.Pop()
	if !ok || value != 30 {
		t.Fatalf("Expected popped element to be 30, got %d", value)
	}
	if s.Size() != 2 {
		t.Fatalf("Expected stack size to be 2 after pop, got %d", s.Size())
	}

	for s.Size() > 0 {
		s.Pop()
	}
	if s.IsEmpty() != true {
		t.Fatal("Expected stack to be empty after popping all elements")
	}

}
