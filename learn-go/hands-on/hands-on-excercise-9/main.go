package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var age int = 32

const name string = "Raja"

func init() {
	fmt.Println("Before running the main function we run niladic init")
}

func main() {
	s1 := "Hello Gophers!!!"
	x := 42

	if x < 40 {
		fmt.Println("I'm not above 40")
	}

	if x > 40 {
		fmt.Println("I'm above 40")
	} else {
		fmt.Println("Guess my age")
	}

	//statement statement idiom
	if y := 10; y < 15 {
		fmt.Println("I'm less than 15 my age is", y)
	}

	if z := 2 * rand.IntN(x); z >= x {
		fmt.Printf("z is %v which is greater than or equal to x %v \n", z, x)
	} else {
		fmt.Printf("z is %v which is less than x %v \n", z, x)
	}

	//switch statements
	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x >= 42:
		fmt.Println("x is greater than or equal to 42")
	default:
		fmt.Println("x doesnt match any cases above")
	}

	switch x {
	case 43:
		fmt.Println("x is 43")
	case 42:
		fmt.Println("x 42")
		fallthrough
	default:
		fmt.Println("I made it fall through")

	}

	//select statement with channel and go routines
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int64N(250))
	d2 := time.Duration(rand.Int64N(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)

	}
	fmt.Println(s1+" "+"from"+" "+name, age)
}
