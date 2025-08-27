package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "Martini", "chocolate"}
	jp := []string{"jenny", "penny", "vodka", "strawberry"}

	fmt.Println(jb)
	fmt.Println(jp)

	xxs := [][]string{jb, jp}
	fmt.Println(xxs)
}
