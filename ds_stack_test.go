package main

import (
	"testing"
)

func Test_stack (t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(1)
	s.Push(1)
	result1 := s.Size()
	if result1 != 3 {
		t.Errorf("Expected result1 to be 3, but got %v", result1)
	}
	result2 := s.Peek()
	if result2 != 1 {
		t.Errorf("Expected result2 to be 1, but got %v", result2)
	}
	result3,_ := s.Pop()
	if result3 != 1 {
		t.Errorf("Expected result3 to be 1, but got %v", result3)
	}
	result4 := s.Size()
	if result4 != 2 {
		t.Errorf("Expected result4 to be 2, but got %v", result1)
	}
}