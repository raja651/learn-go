package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// go func() {
	// 	ch <- 42
	// }()
	ch <- 42
	fmt.Println(<-ch)

	ch <- 43
	ch <- 44

	fmt.Println(<-ch)
}
