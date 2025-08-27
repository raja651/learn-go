package main

import (
	"fmt"
	"math"
	"mymodule/swap"
	"time"
)

var name string = "Raja"

const age int = 32

func main() {
	now := time.Now()
	fmt.Println(now)

	// var a = 2548;clear
	// fmt.Printf("%x", a);

	// var g int;
	// fmt.Println(g);

	// v,b,c,d := 1, 2, 3, 4;

	// fmt.Println(v,b,c,d);

	// print these values as both binary and hexadecimal

	a, b, c, d, e, f := 0, 16, 26, 3, 46, 53

	a = 9

	fmt.Printf("%b %b %b %b %b %b \n", a, b, c, d, e, f)

	fmt.Printf("%X %X %X %X %x %x \n", a, b, c, d, e, f)

	name = "Haritha"
	fmt.Println(name, age)

	fmt.Println(math.Sqrt(9), math.Pi)
	fmt.Println(add(9, 9))

	fmt.Printf("%v %T \n", a, a)

	type ByteSize float64

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		EB
		PB
		ZB
		YB
	)

	fmt.Printf("%T %d %b \n", KB, KB, KB)
	fmt.Printf("%T %d %b \n", MB, MB, MB)
	fmt.Printf("%T %d %b \n", GB, GB, GB)
	fmt.Printf("%T %d %b \n", TB, TB, TB)
	fmt.Printf("%T %d %b \n", PB, PB, PB)
	fmt.Printf("%T %d %b \n", EB, EB, EB)
	fmt.Printf("%T %f %b \n", float64(ZB), float64(ZB), float64(ZB))
	fmt.Printf("%T %f %b \n", float64(YB), float64(YB), float64(YB))

	const (
		_  = iota
		c0 = 1 << iota
		c1
		c2
		c3
		c4
		c5
		c6
	)

	fmt.Printf("%T %v %b \n", c0, c0, c0)
	fmt.Printf("%T %v %b \n", c1, c1, c1)
	fmt.Printf("%T %v %b \n", c2, c2, c2)
	fmt.Printf("%T %v %b \n", c3, c3, c3)
	fmt.Printf("%T %v %b \n", c4, c4, c4)
	fmt.Printf("%T %v %b \n", c5, c5, c5)
	fmt.Printf("%T %v %b \n", c6, c6, c6)

	swap.Test()
}

func add(x, y int) int {
	return x + y
}
