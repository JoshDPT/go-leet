package main

import (
	"testing"
)

func Test_longestSubstringNoRepeat(t *testing.T) {
	result := longestSubstringNoRepeat("abcabcbb")
	if result != 3 {
		t.Errorf("Expected result to be 3, but got %v", result)
	}
	result1 := longestSubstringNoRepeat("bbbbb")
	if result1 != 1 {
		t.Errorf("Expected result to be 1, but got %v", result1)
	}

	result2 := longestSubstringNoRepeat("pwwkew")
	if result2 != 3 {
		t.Errorf("Expected result to be 3, but got %v", result2)
	}
}
