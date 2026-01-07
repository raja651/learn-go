package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Printf("Name:")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Second Name:")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Last Name:")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(answer1, answer2, answer3)
}
