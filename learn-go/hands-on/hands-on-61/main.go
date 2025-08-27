package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Printf("My name is %s and age is %v", p.name, p.age)
}

func main() {
	person1 := person{
		name: "Haritha",
		age:  26,
	}

	person1.speak()
}
