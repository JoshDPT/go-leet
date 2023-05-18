package main

import "testing"

func Test_validPar(t *testing.T) {
	result1 := validParenthesis("()")
	if result1 != true {
		t.Errorf("Expected result to be true, but got %t", result1)
	}
	result2 := validParenthesis("({}[")
	if result2 != false {
		t.Errorf("Expected result to be false, but got %t", result2)
	}
	result3 := validParenthesis("()[]{}")
	if result3 != true {
		t.Errorf("Expected result to be false, but got %t", result3)
	}
}
