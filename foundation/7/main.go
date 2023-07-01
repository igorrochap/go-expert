package main

import "fmt"

func main() {
	salaries := map[string]int{"Igor": 1000, "John": 3000, "Mary": 5000}
	fmt.Println(salaries["Igor"])

	delete(salaries, "John")
	fmt.Println(salaries)

	salaries["Mark"] = 3500

	fmt.Println(salaries)

	functions := make(map[string]string)
	functions["Mark"] = "Developer"

	fmt.Println(functions)

	for name, salary := range salaries {
		fmt.Printf("\nThe salary of %s is %d", name, salary)
	}
}
