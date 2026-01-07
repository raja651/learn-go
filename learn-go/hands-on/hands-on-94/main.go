package main

import "fmt"

func main() {
	gs := 10
	ch := make(chan int)

	send(ch, gs)

	for range 100 {
		fmt.Println(<-ch)
	}
}

func send(ch chan<- int, gs int) {
	for range gs {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
	}
}
