package main

import "fmt"

func main() {
	person := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: `Raja`,
		friends: map[string]int{
			"school":  2,
			"office":  4,
			"college": 5,
		},
		favDrinks: []string{`Coke`, `Pepsi`, `Milkshake`},
	}

	fmt.Println(person)

	foo()
	bar("Krithick")
	favIceCream := getOneFlavour([]string{`Vanilla`, `Chocloate`, `Mango`})
	fmt.Println(favIceCream)
	str, age := calculateDogyears("Simba", 2)
	fmt.Println(str, age)

}

func foo() {
	fmt.Println("Hi this is the first func outside of main")
}

func bar(name string) {
	fmt.Println("Hi my name is", name)
}

func getOneFlavour(listOfIceCream []string) string {
	return fmt.Sprint("My favourite ice cream is ", listOfIceCream[0])
}

func calculateDogyears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint("I am this old in dog years", name, age), age
}
