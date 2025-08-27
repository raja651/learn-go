package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("error in main: %s", err)
	}

	fmt.Println(content)
	fmt.Println(string(content))
}

func readFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("there was an error in reading the file : %s", err)
	}

	return content, nil
}
