package main

import "fmt"

func main() {
	a := 10
	var pointer *int = &a
	fmt.Printf("address of a: %v\n", pointer)

	*pointer = 20
	println(a)

	b := &a
	println(b)
	println(*b)

	*b = 19
	println(a)
}
