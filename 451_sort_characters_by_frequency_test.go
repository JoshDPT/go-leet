package main

import (
	"testing"
)

func Test_frequencySort(t *testing.T) {
	result1 := frequencySort("banana")
	if result1 != "aaannb" {
		t.Errorf("Expected result to be (aaannb), but got %s", result1)
	}

	result2 := frequencySort("cccaa")
	if result2 != "cccaa" {
		t.Errorf("Expected result to be (cccaa), but got %s", result2)
	}

	result3 := frequencySort("Aaabbb")
	if result3 != "bbbaaA" {
		t.Errorf("Expected result to be (bbbaaA), but got %s", result3)
	}
}
