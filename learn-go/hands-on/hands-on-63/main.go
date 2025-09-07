package main

import (
	"fmt"
	"math"
)

func main() {

	// anonymous function
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Test anonymous", i)
		}
	}()

	//function expression
	testFunc := func() int {
		return 40
	}

	fmt.Println(testFunc())

	// function returning a func
	a := returnAFunc()
	fmt.Println(a())

	sampleCallbackFunc(22, square)

}

func returnAFunc() func() int {
	return func() int {
		return 41
	}
}

func sampleCallbackFunc(a int, callback func(num int)) {
	// do something before callback
	callback(a)
}

func square(num int) {
	fmt.Println((math.Pow(float64(num), 2)))
}
