package main

import "fmt"

func main() {
	xi := []int{42, 42, 43, 44, 45, 56}

	//slicing [inclusive:exclusive]
	fmt.Println(xi[0:4])

	//[:exclusive]
	fmt.Println(xi[:6])

	//[inclusive:]
	fmt.Println(xi[3:])

	//all [:]
	fmt.Println(xi[:])

	//delete from a slice
	xi = append(xi[:2], xi[3:]...)
	fmt.Println(xi)
}
