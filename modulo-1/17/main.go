package main

import "fmt"

func main() {
	var myVar interface{} = "Igor"
	println(myVar.(string))

	result, ok := myVar.(int)
	fmt.Printf("The value of the result is %v and the result of \"ok\" is %v\n", result, ok)

	result2 := myVar.(int)
	fmt.Printf("The value of the result is %v", result2)
}
