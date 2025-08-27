package main

import "fmt"

func main() {

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("%v is at the index %v \n", i, i)
	}

	// for loop like while
	var y int = 7
	for y < 10 {
		fmt.Printf("%v is value of y \n", y)
		y++
	}

	// for loop with break
	for {
		fmt.Printf("%v is value of y \n", y)
		if y > 20 {
			break
		}
		y++
	}

	// for loop with continue
	for z := 0; z < 10; z++ {
		if z%2 != 0 {
			continue
		}
		fmt.Println("value of z is", z)
	}

	//for range loops
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("value of %v at index %v \n", v, i)
	}

	//for range loops on map
	m := map[string]int{
		"Raja":    33,
		"haritha": 28,
		"krithik": 2,
	}

	for k, v := range m {
		fmt.Printf("%s has age of %v \n", k, v)
	}
}
