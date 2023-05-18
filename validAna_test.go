package main

import "testing"

func Test_validAna(t *testing.T) {
	result1 := isAnagram("rest", "stre")
	if result1 != true {
		t.Errorf("Expected result to be true, but got %t", result1)
	}
	result2 := isAnagram("car", "carrr")
	if result2 != false {
		t.Errorf("Expected result to be false, but got %t", result2)
	}
	result3 := isAnagram("crap", "carp")
	if result3 != true {
		t.Errorf("Expected result to be false, but got %t", result3)
	}
}
