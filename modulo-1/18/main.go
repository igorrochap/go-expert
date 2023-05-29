package main

type Number int

type Salary interface {
	~int | ~float64
}

func salarySum[T Salary](m map[string]T) T {
	var sum T
	for _, value := range m {
		sum += value
	}
	return sum
}

func isEqual[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	salaryMap := map[string]int{"Igor": 1000, "Maria": 5000, "João": 3000}
	salaryMap2 := map[string]float64{"Igor": 1000.50, "Maria": 3000.15, "João": 2000.20}

	println(salarySum(salaryMap))
	println(salarySum(salaryMap2))

	println(isEqual(10, 10))
}
