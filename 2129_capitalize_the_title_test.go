package main

import (
	"testing"
)

func Test_capitalizeTitle(t *testing.T) {
	result1 := capitalizeTitle("capiTalIze tHe titLe")
	if result1 != "Capitalize The Title" {
		t.Errorf("Expected result1 to be true, but got %v", result1)
	}
	result2 := capitalizeTitle("First leTTeR of EACH Word")
	if result2 != "First Letter of Each Word" {
		t.Errorf("Expected result2 to be false, but got %v", result2)
	}
	result3 := capitalizeTitle("i lOve leetcode")
	if result3 != "i Love Leetcode" {
		t.Errorf("Expected result2 to be false, but got %v", result2)
	}
}
