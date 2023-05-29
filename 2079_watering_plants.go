package main

func wateringPlants(plants []int, capacity int) int {
	jug, steps, position := capacity, 0, -1

	for plants[len(plants)-1] > 0 {
		if jug >= plants[position+1] {
			position += 1
			jug -= plants[position]
			plants[position] = 0
			steps += 1
		} else {
			steps += (position + 1)
			position = -1
			jug = capacity
		}
	}
	return steps
}
