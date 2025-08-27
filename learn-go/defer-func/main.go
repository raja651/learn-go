package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("poem.txt")

	if err != nil {
		return
	}

	defer file.Close()

	data := make([]byte, 100)
	content, err := file.Read(data)

	if err != nil {
		return
	}

	defer fmt.Println("The content is deferred to print last", content, data)

	fmt.Println("This will print before the deferred")

	defer fmt.Println("The content is deferred and printed as string", content, string(data))

	fmt.Println("This will also print before all the deferred print statments")
}
