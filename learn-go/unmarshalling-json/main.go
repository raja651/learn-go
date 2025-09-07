package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Location string `json:"Location"`
}

func main() {
	personJson := `[{"Name":"Dark Knight","Age":45,"Location":"Gotham"},{"Name":"Superman","Age":100,"Location":"New york"}]`
	personByte := []byte(personJson)

	fmt.Println("Bytes of json", personByte)

	var personSlice []person

	err := json.Unmarshal(personByte, &personSlice)

	if err != nil {
		log.Fatalf("Error while unmarshalling json: %s", err)
	}

	fmt.Println("Marshalled json", personSlice)
}
