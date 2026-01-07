package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	c := send(ch)
	receive(ch, c)
}

func send(ch chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for v := range 10 {
			ch <- v
		}
		c <- 1
		close(ch)
	}()
	return c
}

func receive(ch, c <-chan int) {
	fmt.Println(ch)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-c:
			return
		}
	}
}
