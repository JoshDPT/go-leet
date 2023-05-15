package main

import "testing"

func validPar_test(t *testing.T) {
	result1 := validPali("()")
	if result1 != true {
		t.Errorf("Expected result to be false, but got %t", result1)
	}
	result2 := validPali("(()")
	if result2 != false {
		t.Errorf("Expected result to be false, but got %t", result2)
	}
}