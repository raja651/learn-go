package main

import "fmt"

func addOne(v int) int {
	return v + 1
}

func addOnePointer(v *int) {
	*v += 1
}

func main() {

	// value semantics
	a := 1
	fmt.Println(a)
	b := addOne(a)
	fmt.Println(b)

	// pointer semantics
	c := 3
	fmt.Println(c)
	fmt.Println(&c)
	addOnePointer(&c)
	fmt.Println(c)

}
