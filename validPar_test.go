package main

import "testing"

func Test_validPar(t *testing.T) {
	result1 := validPali("()")
	if result1 != true {
		t.Errorf("Expected result to be true, but got %t", result1)
	}
	result2 := validPali("({}[")
	if result2 != false {
		t.Errorf("Expected result to be false, but got %t", result2)
	}
	result3 := validPali("()[]{}")
	if result3 != true {
		t.Errorf("Expected result to be false, but got %t", result3)
	}
}
