package main

import "fmt"

func main() {
	ch := make(chan int)

	go foo(ch)

	bar(ch)

	fmt.Println("About to exit")
}

func foo(ch chan<- int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

func bar(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
