package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name     string
	Age      int
	Location string
}

func main() {
	p1 := Person{
		Name:     `Dark Knight`,
		Age:      45,
		Location: "Gotham",
	}

	p2 := Person{
		Name:     "Superman",
		Age:      100,
		Location: "New york",
	}

	personSlice := []Person{p1, p2}

	stringify, err := json.Marshal(&personSlice)

	if err != nil {
		log.Fatalf("This is a error: %s", err)
	}

	fmt.Println(personSlice)
	fmt.Println(string(stringify))
}
