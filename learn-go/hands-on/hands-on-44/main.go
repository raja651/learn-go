package main

import "fmt"

func main() {
	xis := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xis = append(xis, 52)
	fmt.Println(xis)

	xis = append(xis, 53, 54, 55)
	fmt.Println(xis)

	y := []int{56, 57, 58, 59, 60}

	xis = append(xis, y...)
	fmt.Println(xis)
	fmt.Println(y)

	// delete values from slice
	xis = append(xis[:3], xis[6:10]...)
	fmt.Println(xis)
}
