package main

import "fmt"

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	in := incrementor()
	for i := 0; i < 10; i++ {
		fmt.Println(in())
	}

	fmt.Println(factorial(4))
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	num = num * factorial(num-1)

	return num
}
