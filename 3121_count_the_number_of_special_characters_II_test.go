package main

import (
	"testing"
)

func Test_numberOfSpecialCharsTwo(t *testing.T) {
	result := numberOfSpecialCharsTwo("aaAbcBC")
	if result != 3 {
		t.Errorf("Expected result to be 3, but got %v", result)
	}
	result2 := numberOfSpecialCharsTwo("abc")
	if result2 != 0 {
		t.Errorf("Expected result to be 0, but got %v", result2)
	}
	result3 := numberOfSpecialCharsTwo("AbBCab")
	if result3 != 0 {
		t.Errorf("Expected result to be 1, but got %v", result3)
	}
}
