package main

import "fmt"

func main() {
	ch := make(chan<- int)

	// go func() {
	// 	ch <- 42
	// }()
	ch <- 42
	ch <- 43

	fmt.Printf("%T\n", ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
}
