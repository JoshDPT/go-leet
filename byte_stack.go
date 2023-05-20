package main

import (
	"errors"
)

// defining type stack which as items which is an array of values
type ByteStack struct {
	items []byte
}

// s pointing to the ByteStack is similar to this in TS
func (s *ByteStack) Push (value byte) {
	s.items = append(s.items, value)
}

// we can return two variables using parentheses 
func (s *ByteStack) Pop () (byte, error) {
	// if it is empty, return 0 and throw error
	if s.IsEmpty() {
		return 0, errors.New("ByteStack is empty")
	}

	// idiomatic was to pop - get last index, define value, slice the slice at that index
	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]

	return value, nil
}

func (s *ByteStack) IsEmpty () bool {
	return len(s.items) == 0
}

func (s *ByteStack) Peek () byte {
	return s.items[len(s.items)-1]
}

func (s *ByteStack) Size () int {
	return len(s.items)
}
