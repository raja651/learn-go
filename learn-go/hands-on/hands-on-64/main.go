package main

import (
	"fmt"
	"time"
)

func timedFunction(f func()) {
	start := time.Now()
	f()
	elapsedTime := time.Since(start)

	fmt.Println("Elasped time", elapsedTime)
}

func myFunction() {
	time.Sleep(2 * time.Second)
	fmt.Println("My function completed")
}

func doWork() {
	for i := 0; i < 1_000; i++ {
		fmt.Println(i)
	}
}

func main() {
	timedFunction(myFunction)
	timedFunction(doWork)
}
