package main

import "fmt"

type animal struct {
	name string
}

func (a animal) bark() {
	fmt.Println("I am barking", a.name)
}

type bigAnimal interface {
	bark()
}

func barkForMe(b bigAnimal) {
	b.bark()
}

func main() {

	ints := []int{1, 2, 3, 4, 56, 7, 87, 9}
	total := sum(ints...)

	fmt.Println(total)

	dog := animal{
		name: `Simba`,
	}

	dog.bark()

	barkForMe(dog)
}

func sum(numbers ...int) int {
	fmt.Println(numbers)
	fmt.Printf("%T \n", numbers)

	total := 0
	for _, val := range numbers {
		total += val
	}

	return total
}
