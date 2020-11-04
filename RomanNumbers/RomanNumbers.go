package main

import (
	"fmt"
	"strings"
)

var result int
var voc = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func roman(roman string) (result int) {
	result = 0
	array := strings.Split(roman, "")
	for i := 0; i < len(array); i++ {
		if (voc[array[i]] == 0) {
			return 0
		}
		if (((i + 1) < len(array)) && (voc[array[i]] < voc[array[i + 1]])){
			result += voc[array[i + 1]] - voc[array[i]]
			i++
		} else {
			result += voc[array[i]]
		}
	}
  	return 
}

func main() {
	fmt.Println(roman("IX"))
}