package math

type Math struct {
	a int
	b int
}

func Construct(a, b int) Math {
	return Math{a, b}
}

func (math Math) Add() int {
	return math.a + math.b
}
