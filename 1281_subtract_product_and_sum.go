package main

func subtractProductAndSum(n int) int {
	s := digitSum(n)
	p := digitProduct(n)
	dif := s - p
	return dif
}