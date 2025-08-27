package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func doArithmetic(a, b int, arith func(int, int) int) int {
	return arith(a, b)
}

func main() {
	sum := doArithmetic(5, 7, add)
	minus := doArithmetic(10, 7, subtract)

	fmt.Println(sum, minus)
}
