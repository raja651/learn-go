package main

import "fmt"

func main() {
	res := foo()
	res2, res3 := bar()

	fmt.Println(res)
	fmt.Println(res2, res3)

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sumOfSliceVariadic(xi...)
	sum2 := sumOfSlice(xi)
	fmt.Println(sum)
	fmt.Println(sum2)
}

func foo() int {
	return 32
}

func bar() (string, int) {
	return "My age is: ", 32
}

func sumOfSliceVariadic(s ...int) int {
	fmt.Println(s)
	total := 0
	for _, val := range s {
		total += val
	}

	return total
}

func sumOfSlice(s []int) int {
	total := 0
	for _, val := range s {
		total += val
	}

	return total
}
