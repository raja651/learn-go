package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return `Hello I'm ` + p.name
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p := person{name: "Raja"}
	saySomething(&p)
}
