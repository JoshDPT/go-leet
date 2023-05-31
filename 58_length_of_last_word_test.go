package main

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	result1 := lengthOfLastWord("Hello World")
	if result1 != 5 {
		t.Errorf("Expected result to be 5, but got %v", result1)
	}

	result2 := lengthOfLastWord("   fly me   to   the moon  ")
	if result2 != 4 {
		t.Errorf("Expected result to be 4, but got %v", result2)
	}

	result3 := lengthOfLastWord("luffy is still joyboy")
	if result3 != 6 {
		t.Errorf("Expected result to be 6, but got %v", result3)
	}
}
