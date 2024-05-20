package main

import (
	"github.com/google/uuid"
	"github.com/igorrochap/goexpert/5-packaging/3/math"
)

func main() {
	m := math.Construct(1, 2)
	println(m.Add())

	println(uuid.New().String())

	// go mod tidy -e
}
