package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	igor := Client{
		Name:   "Igor",
		Age:    22,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", igor.Name, igor.Age, igor.Active)

	igor.Active = false

	fmt.Printf("Active: %t\n", igor.Active)

}
