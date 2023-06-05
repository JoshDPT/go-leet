// Input: ransomNote = "a", magazine = "b"
// Output: false

// Input: ransomNote = "aa", magazine = "ab"
// Output: false

// Input: ransomNote = "aa", magazine = "aab"
// Output: true

package main

import (
	"testing"
)

func Test_ransomNote(t *testing.T) {

	result := ransomNote("a", "b")
	if result != false {
		t.Errorf("Expected result to be false, but got %v", result)
	}

	result1 := ransomNote("aa", "bb")
	if result1 != false {
		t.Errorf("Expected result1 to be false, but got %v", result1)
	}

	result2 := ransomNote("aa", "aab")
	if result2 != true {
		t.Errorf("Expected result2 to be false, but got %v", result2)
	}
}
