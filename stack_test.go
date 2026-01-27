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
	if s.Len() != 3 {
		t.Fatalf("Expected stack Len to be 3, got %d", s.Len())
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
	if s.Len() != 2 {
		t.Fatalf("Expected stack Len() to be 2 after pop, got %d", s.Len())
	}

	for s.Len() > 0 {
		s.Pop()
	}
	if s.IsEmpty() != true {
		t.Fatal("Expected stack to be empty after popping all elements")
	}

}
