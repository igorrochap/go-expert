package main

import "fmt"

type id int

var (
	b bool    = true
	c int     = 10
	d string  = "Igor"
	e float64 = 1.2
	f id      = 123
)

func main() {
	fmt.Printf("O tipo de E é: %T\n", e)
	fmt.Printf("O valor de E é: %v\n", e)

	fmt.Printf("O tipo de F é: %T\n", f)
	fmt.Printf("O valor de F é: %v\n", f)
}
