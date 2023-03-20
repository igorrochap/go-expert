package main

import "fmt"

func main() {
	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 15

	// fmt.Println(myArray[len(myArray)-1])

	for i, v := range myArray {
		fmt.Printf("The value of index %d is %d\n", i, v)
	}
}
