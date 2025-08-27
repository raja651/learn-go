package main

import "fmt"

var zero int

const (
	_ = iota
	q
	w
	r
)

func main() {
	short := "short"

	a, b, c, d := 1, "test", 42.0, false

	var flo float32 = 5.3234

	fmt.Printf(" %v \n", zero)
	fmt.Println(short)
	fmt.Printf("%v \t %v \t %v \t %v \t \n", a, b, c, d)
	fmt.Println(flo, q, w, r)

	test2, test3, test4 := "Raja", 32, 11.02

	fmt.Printf("%v %s %T \t %v %b %T \t %v %f %T \t ", test2, test2, test2, test3, test3, test3, test4, test4, test4)
}
