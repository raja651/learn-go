package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210

	fmt.Printf("%v of type %T in \t decimal %d \t binary %b \t\t hexadecimal %#X \n", a, a, a, a, a)
	fmt.Printf("%v of type %T in \t decimal %d \t binary %b \t\t hexadecimal %#X \n", b, b, b, b, b)
	fmt.Printf("%v of type %T in \t decimal %d \t binary %b \t hexadecimal %#X \n", c, c, c, c, c)
}
