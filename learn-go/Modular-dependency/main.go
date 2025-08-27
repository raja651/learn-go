package main

import (
	"fmt"

	"github.com/raja651/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1, s2)
}
