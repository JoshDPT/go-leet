package main

import (
	"strings"
	"strconv"
)

func submarineDive (s string) int {
	x, y := 0, 0
	arr := strings.Fields(s)
	for i,c := range arr {
		if i == len(arr)-1 {
			continue
		} 
		n,_ := strconv.Atoi(arr[i+1])
		switch c {
		case "forward":
			x += n
		case "down":
			y += n
		case "up":
			y -= n
		}
	}
	return x * y;
}