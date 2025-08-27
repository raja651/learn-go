package main

import "fmt"

func main() {
	total := sum(5, 5)

	fmt.Println(total)
}

func sum(a, b int) int {
	return a + b
}
