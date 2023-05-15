package main

import "testing"

func validPali_test(t *testing.T) {
	result1 := validPali("1234")
	if result1 != false {
		t.Errorf("Expected result to be false, but got %t", result1)
	}
	result2 := validPali("abcdcba")
	if result2 != true {
		t.Errorf("Expected result to be true, but got %t", result2)
	}
	result3 := validPali("Abc, !dc,   ba   #%")
	if result3 != true {
		t.Errorf("Expected result to be true, but got %t", result3)
	}
}