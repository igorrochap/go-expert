package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	igor := Client{
		Name:   "Igor",
		Age:    22,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", igor.Name, igor.Age, igor.Active)

	igor.City = "Maceió"

	fmt.Println(igor.Address)

}
