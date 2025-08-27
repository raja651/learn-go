package main

import "fmt"

func main() {
	m := map[string]int{
		"Henry": 32,
		"Jason": 34,
		"Msd":   23,
		"lsd":   35,
	}

	fmt.Println(m)

	delete(m, "Msd")

	fmt.Println(m)

	v, ok := m["Msd"]
	if ok {
		fmt.Println("The key exists", v)
	} else {
		fmt.Println("The key doesn't exist", v)
	}

	// idomatic go code to check comma ok idioms
	// go process to if we encounter errors lets deal with it now rather than later
	if v, ok := m["lsd"]; !ok {
		fmt.Println("the value doesn't exist", v)
	} else {
		fmt.Println("the value does exist", v)
	}

}
