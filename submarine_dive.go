package main

import (
	"strings"
	"strconv"
	"log"
)

type Tuple struct {
	x int
	y int
}

func parseLine (line string) Tuple {
	arr := strings.Fields(line)
	dir := arr[0]
	amount, err := strconv.Atoi(arr[1])

	if err != nil {
		log.Fatal("This should not happen")
	}

	switch dir {
		case "forward": {
			return Tuple{
				x: amount,
				y: 0,
			}
		}
		case "down": {
			return Tuple{
				x: 0,
				y: amount,
			}
		}
		case "up": {
			return Tuple{
				x: 0,
				y: -amount,
			}
		}
	}
	return Tuple{0,0}
}

func submarineDive (s string) int {
		lines := strings.Split(s, "\n")
		
		pos := Tuple{0,0}

		for _, line := range lines {
			coord := parseLine(line)
			pos.x += coord.x
			pos.y += coord.y
		}

		return pos.x * pos.y
}