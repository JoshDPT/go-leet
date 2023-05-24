package main

import (
	"errors"
)

// defining type stack which as items which is an array of values
type Stack struct {
	items []interface{}
}

// s pointing to the Stack is similar to "this" in TS
func (s *Stack) Push (value interface{}) {
	s.items = append(s.items, value)
}

// we can return two variables using parentheses 
func (s *Stack) Pop () (interface{}, error) {
	
	// if it is empty, return 0 and throw error
	if s.IsEmpty() {
		ErrStackEmpty := errors.New("Stack is empty")
		return 0, ErrStackEmpty
	}

	// idiomatic was to pop - get last index, define value, slice the slice at that index
	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]

	return value, nil
}

func (s *Stack) IsEmpty () bool {
	return len(s.items) == 0
}

func (s *Stack) Peek () interface{} {
	return s.items[len(s.items)-1]
}

func (s *Stack) Size () int {
	return len(s.items)
}
