package main

func rotateImage(m [][]int) [][]int {
	l := len(m)

	// invert
	for i := 0; i < l; i++ {
		for j := i+1; j < l; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	
	// reverse
	for i := 0; i < l; i++ {
		for j := 0; j < (l/2); j++ {
			m[i][j], m[i][l-j-1] = m[i][l-j-1], m[i][j]
		}
	}
	return m
}