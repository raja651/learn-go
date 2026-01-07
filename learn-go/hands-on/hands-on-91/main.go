package main

import "fmt"

func main() {
	ch := make(chan int)

	go send(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func send(ch chan<- int) {
	for v := range 10 {
		ch <- v
	}
	close(ch)
}
