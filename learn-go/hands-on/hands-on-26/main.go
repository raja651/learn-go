package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var y int = 1
	for {
		fmt.Printf("The value of y is %v \n", y)
		if y > 1 {
			break
		}
		y++
	}

	// use modulus and continue statement to print all odd numbers
	for x := 0; x <= 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}

	//nest a loop within a loop
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Printf("%v \t %v \n", i, j)
	// 	}

	// 	fmt.Println("*****************")
	// }

	//for loop range
	xi := []int{42, 43, 44, 45, 46, 47, 48}

	for i, v := range xi {
		fmt.Printf("The index is %v and value is %v \n", i, v)
	}

	m := map[string]int{
		"Haritha": 27,
		"Raja":    32,
		"Mugilan": 2,
		"Mohan":   28,
	}

	rajaAge := m["Raja"]
	fmt.Println("Raja age is", rajaAge)

	age := m["Q"]
	fmt.Println(age)

	if qAge, ok := m["Q"]; ok {
		fmt.Println("Value exists for age", qAge)
	}

	for k, v := range m {
		fmt.Printf("The key is %s and the value is %v \n", k, v)
	}

	// statement statement idiom diff from comma ok idiom
	count := 1
	for i := 0; i < 100; i++ {
		if a := rand.IntN(5); a == 3 {
			fmt.Println(i, a, count)
			count++
		}
	}
}
