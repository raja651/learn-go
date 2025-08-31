package main

import "fmt"

func main() {
	x := 43
	xPtr := &x

	fmt.Println(x, xPtr, *xPtr)

	fmt.Printf("%T ---- %v \n", x, x)
	fmt.Printf("%T ---- %v \n", xPtr, xPtr)
	fmt.Printf("%T ---- %v \n", *xPtr, *xPtr)
	fmt.Printf("%T ---- %v \n", &xPtr, &xPtr)

	*xPtr = 45

	fmt.Println(x)
}
