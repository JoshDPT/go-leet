package main

import (
	"testing"
)

func Test_queue(t *testing.T) {

	q := Queue{}
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)

	result1 := q.Size()
	if result1 != 3 {
		t.Errorf("Expected result1 to be 3, but got %v", result1)
	}

	result2, _ := q.Dequeue()
	if result2 != 6 {
		t.Errorf("Expected result2 to be 6, but got %v", result2)
	}

	result3, _ := q.Dequeue()
	if result3 != 7 {
		t.Errorf("Expected result3 to be 7, but got %v", result3)
	}

	result4 := q.IsEmpty()
	if result4 != false {
		t.Errorf("Expected result4 to be false, but got %v", result4)
	}

	result5, _ := q.Dequeue()
	if result5 != 8 {
		t.Errorf("Expected result5 to be 8, but got %v", result5)
	}

	result6 := q.IsEmpty()
	if result6 != true {
		t.Errorf("Expected result6 to be true, but got %v", result6)
	}

}
