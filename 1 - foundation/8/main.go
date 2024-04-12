package main

import (
	"errors"
	"fmt"
)

func main() {
	value1, err := sum(2, 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value1)

	value2, err := sum(10, 41)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value2)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("the sum is greater then 50")
	}
	return a + b, nil
}
