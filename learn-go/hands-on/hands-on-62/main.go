package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	squareParams := square{
		length: 23.5,
		width:  23.5,
	}

	circleParams := circle{
		radius: 4.5,
	}

	info(squareParams)
	info(circleParams)
}
