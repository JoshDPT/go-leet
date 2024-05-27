package main

import (
	"testing"
)

func Test_numberOfSpecialChars(t *testing.T) {
	result := numberOfSpecialChars("aaAbcBC")
	if result != 3 {
		t.Errorf("Expected result to be 3, but got %v", result)
	}
	result2 := numberOfSpecialChars("abc")
	if result2 != 0 {
		t.Errorf("Expected result to be 0, but got %v", result2)
	}
	result3 := numberOfSpecialChars("abBCab")
	if result3 != 1 {
		t.Errorf("Expected result to be 1, but got %v", result3)
	}
	result4 := numberOfSpecialChars("cCC")
	if result4 != 1 {
		t.Errorf("Expected result to be 1, but got %v", result4)
	}
}
