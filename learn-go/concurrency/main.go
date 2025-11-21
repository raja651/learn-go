package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS \t", runtime.GOOS)
	fmt.Println("Arch \t", runtime.GOARCH)
	fmt.Println("Cpu's \t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	fmt.Println("--------------------------------")
	wg.Add(1)
	go foo()

	fmt.Println("--------------------------------")
	bar()

	fmt.Println("--------------------------------")
	fmt.Println("Cpu's \t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Printf(" foo %v \t", i)
	}
	wg.Done()
}

func bar() {
	i := 0
	for i < 10 {
		fmt.Printf("%v \t", i)
		i++
	}
}
