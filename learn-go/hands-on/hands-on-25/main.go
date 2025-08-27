package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("The iteration value is %v \t", i)

		switch x {
		case 0:
			fmt.Println("The value of x is 0")
		case 1:
			fmt.Println("The value of x is 1")
		case 2:
			fmt.Println("The value of x is 2")
		case 3:
			fmt.Println("The value of x is 3")
		case 4:
			fmt.Println("The value of x is 4")
		default:
			fmt.Println("The value of x is greater than 4")
		}
	}

	//booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}
