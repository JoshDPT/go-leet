package main

func checkStraightLine(coord [][]int) bool {

	x0, x1 := coord[0][0], coord[1][0]
	y0, y1 := coord[0][1], coord[1][1]

	for i := 1; i < len(coord); i++ {
		x2,y2 := coord[i][0], coord[i][1]
		if (x1 - x0) * (y2-y1) != (x2-x1) * (y1-y0) {
			return false
		}
	}

	return true
}