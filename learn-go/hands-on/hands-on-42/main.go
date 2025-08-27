package main

import "fmt"

func main() {
	xi := [5]int{}

	for i := 0; i < len(xi); i++ {
		xi[i] = i
	}

	for _, v := range xi {
		fmt.Println(v)
	}

	xis := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range xis {
		fmt.Printf("%v - %T - %v \n", i, v, v)
	}

	// create seprate slices
	sli1 := xis[:5]
	sli2 := xis[5:]
	sli3 := xis[2:7]
	sli4 := xis[1:6]

	fmt.Printf("%#v  \t %v \n", sli1, sli1)
	fmt.Printf("%#v  \t %v \n", sli2, sli2)
	fmt.Printf("%#v  \t %v \n", sli3, sli3)
	fmt.Printf("%#v  \t %v \n", sli4, sli4)

}
