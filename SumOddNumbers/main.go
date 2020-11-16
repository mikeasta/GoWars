package main

import "fmt"

// Given the triangle of consecutive odd numbers:

//              1
//           3     5
//        7     9    11
//    13    15    17    19
// 21    23    25    27    29
// ...

// Calculate the row sums of this triangle from the row index (starting at index 1) e.g.:
// rowSumOddNumbers(1); // 1
// rowSumOddNumbers(2); // 3 + 5 = 8

func rowSumOddNumbers(n int) int {
	result  := 0
	counter := 0
	for j:= 1; j < n; j++{
		counter += j
	}

	for i:=1; i <= n; i++ {
		result += (counter + i - 1) * 2 + 1
	}
	return result
}

func main() {
	fmt.Println(rowSumOddNumbers(2))
}