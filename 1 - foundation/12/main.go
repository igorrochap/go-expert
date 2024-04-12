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

func (client Client) Deactivate() {
	client.Active = false
	fmt.Printf("The client %s was deactivated\n", client.Name)
}

func main() {
	igor := Client{
		Name:   "Igor",
		Age:    22,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", igor.Name, igor.Age, igor.Active)

	igor.City = "Maceió"

	igor.Deactivate()
}
