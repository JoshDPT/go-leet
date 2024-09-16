package main

import "testing"

func Test_decodeAtIndex(t *testing.T) {
	result1 := decodeAtIndex("leet2code3", 10)
	if result1 != "o" {
		t.Errorf("Expected result to be false, but got %s", result1)
	}
	result2 := decodeAtIndex("ha22", 5)
	if result2 != "h" {
		t.Errorf("Expected result to be true, but got %s", result2)
	}
	result3 := decodeAtIndex("a2345678999999999999999", 1)
	if result3 != "a" {
		t.Errorf("Expected result to be true, but got %s", result3)
	}
}
