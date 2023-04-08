package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Person interface {
	Deactivate()
}

func Deactivation(person Person) {
	person.Deactivate()
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

	Deactivation(igor)
}
