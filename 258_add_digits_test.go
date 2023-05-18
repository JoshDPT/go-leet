package main

import (
	"testing"
)

func Test_addDigits (t *testing.T) {
	result := addDigits(38)
	if result != 2 {
		t.Errorf("Expected result to be true, but got %v", result)
	}
}