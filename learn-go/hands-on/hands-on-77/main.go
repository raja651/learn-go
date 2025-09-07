package main

import "fmt"

type person struct {
	name string
}

func (p person) changeName(newName string) {
	p.name = newName
}

func main() {
	p1 := person{
		"Raja",
	}

	p1.changeName("haritha")
	fmt.Println(p1)

	p2 := person{
		"Krithu",
	}
	p3 := changeName(p2, "Krithik")
	fmt.Println(p3)

	p4 := &person{
		"Mugilan",
	}
	changeNamePtr(p4, "Mugi")
	fmt.Println(*p4)

}

func changeName(p person, newName string) person {
	p.name = newName
	return p
}

func changeNamePtr(p *person, newName string) {
	p.name = newName
}
