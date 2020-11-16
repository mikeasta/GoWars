package main

import "fmt"

func nbYear(p0 int, percent float64, aug int, p int) int {
	var i int
	var pop float64 = float64(p0)
	var augf float64 = float64(aug)
	for i = 0; pop < float64(p); i++ {
		pop = pop + pop * percent * 0.01 + augf
	}
	return i
}

func main() {
	fmt.Println(nbYear(1500, 5, 100, 5000))
}