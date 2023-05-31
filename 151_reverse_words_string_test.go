package main

import "testing"

func Test_reverseWords(t *testing.T) {
	result1 := reverseWords("the sky is blue")
	if result1 != "blue is sky the" {
		t.Errorf("Expected result to be (blue is sky the), but got %s", result1)
	}

	result2 := reverseWords("  hello world  ")
	if result2 != "world hello" {
		t.Errorf("Expected result to be (world hello), but got %s", result2)
	}

	result3 := reverseWords("a good   example")
	if result3 != "example good a" {
		t.Errorf("Expected result to be example good a, but got %s", result3)
	}
}
