package main

import "fmt"

type Client struct {
	Name     string
	Position int
}

func NewClient(name string) *Client {
	return &Client{Name: name, Position: 0}
}

func (client *Client) Walked() {
	client.Position += 1
	fmt.Printf("The client %v walked\n", client.Name)
}

func main() {
	igor := NewClient("Igor")
	fmt.Printf("%v is in position %d\n", igor.Name, igor.Position)

	igor.Walked()
	fmt.Printf("%v is in position %d\n", igor.Name, igor.Position)

	igor.Walked()
	fmt.Printf("%v is in position %d\n", igor.Name, igor.Position)
}
