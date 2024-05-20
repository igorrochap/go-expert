package main

import (
	"fmt"

	"github.com/igorrochap/goexpert/5-packaging/1/math"
)

func main() {
	m := math.Construct(1, 2)
	fmt.Println(m.Add())
}
