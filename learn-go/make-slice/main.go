package main

import "fmt"

func main() {
	xi := make([]int, 0, 10)

	fmt.Println(xi)
	fmt.Println("size", len(xi))
	fmt.Println("capactiy", cap(xi))

	fmt.Println("-------------------------------------------")

	// lets add some value to slice
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("size", len(xi))
	fmt.Println("capactiy", cap(xi))

	fmt.Println("-------------------------------------------")

	//add some more values to size and see if the capactiy increases
	xi = append(xi, 10, 11, 12, 13, 14)
	fmt.Println(xi)
	fmt.Println("size", len(xi))
	fmt.Println("capactiy", cap(xi))

	fmt.Println("-------------------------------------------")

	//add some more
	xi = append(xi, 15, 16, 17, 18, 19, 20, 21, 22)
	fmt.Println(xi)
	fmt.Println("size", len(xi))
	fmt.Println("capactiy", cap(xi))
	fmt.Println("-------------------------------------------")

	// delete some values to see if capacity goes down
	xi = append(xi[:5], xi[19:]...)
	fmt.Println(xi)
	fmt.Println("size", len(xi))
	fmt.Println("capactiy", cap(xi))
}
