

// Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
// Output: true

// Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
// Output: false

package main

import "testing"

func Test_checkStraightLine(t *testing.T) {
	result := checkStraightLine([][]int{{1,2},{2,3},{3,4},{4,5},{5,6},{6,7}})
	if result != true {
		t.Errorf("Expected result to be true, but got %t", result)
	}
}