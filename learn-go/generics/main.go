package main

import "fmt"

type MyNumbers interface {
	int | float64
}

type age int

type Number interface {
	~int | ~float64
}

func addT[T MyNumbers](a, b T) T {
	return a + b
}

func mutiply[T Number](a, b T) T {
	return a * b
}

func main() {
	fmt.Println(addT(4, 4))
	fmt.Println(addT(5.6, 6.45))

	var a age = 54
	var b age = 51

	fmt.Println(mutiply(a, b))
}
