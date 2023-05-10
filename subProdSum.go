package main

import "fmt"

func subtractProductAndSum(n int) int {
	s := digitSum(n)
	p := digitProduct(n)
	dif := s - p
	return dif
}

func main() {
	n := subtractProductAndSum(1234)
	fmt.Println(n)
}
