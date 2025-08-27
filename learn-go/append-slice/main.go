package main

import "fmt"

func main() {

	xi := []int{42, 42, 44}

	fmt.Println(xi[0])
	fmt.Println(xi)

	//append something to slide
	// variadic parameter meaning any number of parameters
	xi = append(xi, 45, 46, 46)

	fmt.Println(xi)

	xs := [3]string{"Raja", "hello"}
	xs[2] = "world"
	fmt.Println(xs)
}
