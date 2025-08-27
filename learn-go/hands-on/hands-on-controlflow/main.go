package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where the intialization of my program occurs")
}

func main() {
	x := rand.Intn(250)

	fmt.Printf("x has value of %d \n", x)

	if x <= 100 {
		fmt.Println("The value is between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("The value is between 101 and 200")
	} else if x >= 210 && x <= 250 {
		fmt.Println("The value is between 201 and 250")
	} else {
		fmt.Println("This was more than 250")
	}

	y := rand.Intn(3)

	fmt.Printf("The value of y is %v \n", y)

	// using switch statement
	switch {
	case x <= 100:
		fmt.Println("The value is between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("The value is between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("The value is between 201 and 250")
	default:
		fmt.Println("This was more than 250")
	}

}
