package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 6)
	copy(b, a)

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("---------------")

	a[0] = 9
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("---------------")

}
