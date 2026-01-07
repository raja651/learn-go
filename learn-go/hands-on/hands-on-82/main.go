package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go printBye()
	go printHello()

	fmt.Println(runtime.NumCPU(), runtime.NumGoroutine())

	wg.Wait()

}

func printHello() {
	fmt.Println("Hello!!!! I'm back")
	wg.Done()
}

func printBye() {
	fmt.Println("Hello i'm going back")
	wg.Done()
}
