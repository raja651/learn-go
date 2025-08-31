package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println("I'm walking", d.name)
	d.name = "rosie"
	fmt.Println("I'm walking", d.name)
}

func (d dog) run() {
	fmt.Println("I'm running", d.name)
	d.name = "Jhonny"
	fmt.Println("I'm running", d.name)
}

type youngin interface {
	run()
	walk()
}

func younginrun(y youngin) {
	y.run()
}

func main() {
	dog1 := dog{
		name: "Simba",
	}

	dog1.walk()
	younginrun(dog1)

	fmt.Println("----------------------------------------")

	dog2 := dog{
		name: "Blacky",
	}

	dog2.walk()
	younginrun(dog2)
}
