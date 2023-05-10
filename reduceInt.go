package main

func reduceInt (arr []int, f func(int, int)int, start int ) int {
	acc := start

	for _, v := range arr {
		acc = f(acc, v)
	}

	return acc
}