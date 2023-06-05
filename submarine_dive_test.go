package main

import "testing"

func Test_submarineDive(t *testing.T) {
	str := `forward 5 
	down 5 
	forward 8 
	up 3 
	down 8 
	forward 2`

	result := submarineDive(str)

	if result != 150 {
		t.Errorf("Expected result to be 150, but got %v", result)
	}

	str2 := `forward 1 
	down 1 
	forward 1 
	up 1 
	down 1 
	forward 1`

	result2 := submarineDive(str2)

	if result2 != 3 {
		t.Errorf("Expected result to be 3, but got %v", result)
	}
}
