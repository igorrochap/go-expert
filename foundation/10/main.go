package main

import (
	"fmt"
)

func main() {
	sumDoubled := func() int {
		return sum(1, 2, 3, 4, 6, 7, 8, 9, 10) * 2
	}()

	fmt.Println(sumDoubled)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
