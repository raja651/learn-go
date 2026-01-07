package main

import "fmt"

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("%T\n", cs)
}
