package main
import (
	"fmt"
	"strings"
	"strconv"
)

func highAndLow(in string) string {
	array := strings.Split(in, " ")
	maxElement, _ := strconv.ParseInt(array[0], 10, 64)
	minElement, _ := strconv.ParseInt(array[0], 10, 64)
	for _, element := range array {
		num, _ := strconv.ParseInt(element, 10, 64)
		if (maxElement < num ) {
			maxElement = num
		} else if (minElement > num) {
			minElement = num
		}
	}
	return strconv.FormatInt(maxElement, 10) + " " + strconv.FormatInt(minElement, 10)
}

func main() {
	fmt.Println(highAndLow("1 -2 3 100"))
}