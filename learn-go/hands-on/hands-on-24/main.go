package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("The value of i is %v \n", i)
	// }

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("The value of x is %v \t", x)
		fmt.Printf("The value of y is %v \t", y)
		fmt.Printf(" The value of iteration i is %v \t", i)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("x and y are both less than 4 \n")
		case x > 6 && y > 6:
			fmt.Printf("x and y are both greateer than 6 \n")
		case x >= 4 && x <= 6:
			fmt.Printf("x is greater than or equal to 4 and less than or equal to 6\n")
		case y != 5:
			fmt.Printf("y is not equal to 5\n")
		default:
			fmt.Printf("None of the previous cases were met \n")
		}
	}
}
