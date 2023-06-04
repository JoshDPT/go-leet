package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type Tuple struct {
	x int
	y int
}

func parseLine (line string) (Tuple, error) {
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
			}, nil
		}
		case "down": {
			return Tuple{
				x: 0,
				y: amount,
			}, nil
		}
		case "up": {
			return Tuple{
				x: 0,
				y: -amount,
			}, nil
		}
		default: {
			return Tuple{0,0}, errors.New("This should not happen")
		}
	}
}

func submarineDive (s string) int {
		lines := strings.Split(s, "\n")
		
		pos := Tuple{0,0}

		for _, line := range lines {
			coord,err := parseLine(line)

			if err != nil {
				log.Fatal("This should not happen")
			}

			pos.x += coord.x
			pos.y += coord.y
		}

		return pos.x * pos.y
}