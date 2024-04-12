package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(5, 5, 5))
	fmt.Println(sum(1, 2, 3, 4, 6, 7, 8, 9, 10))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
