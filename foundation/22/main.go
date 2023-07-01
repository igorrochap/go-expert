package main

func main() {
	x := 4
	y := 1
	z := 3

	if x > y && z > x {
		println("x > y && z > x")
	}

	// if x > y {
	// 	println(x)
	// } else {
	// 	println(y)
	// }

	switch x {
	case 1:
		println("y")
	case 2:
		println("x")
	case 3:
		println("z")
	default:
		println("default case")
	}
}
