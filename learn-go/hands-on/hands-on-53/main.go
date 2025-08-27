package main

import "fmt"

func main() {

	type person struct {
		firstName       string
		lastName        string
		iceCreamFlavour []string
	}

	person1 := person{
		firstName:       "Dark",
		lastName:        "Knight",
		iceCreamFlavour: []string{`Vanilla`, `strawberry`, `mango`, `rose`},
	}

	person2 := person{
		firstName:       "Iron",
		lastName:        "Man",
		iceCreamFlavour: []string{`chocolate`, `pineapple`, `latee`, `pancar`},
	}

	fmt.Println(person1)
	fmt.Println(person2)

	persons := make(map[string]person)

	persons["Knight"] = person1
	persons["Man"] = person2

	//range over the map
	for key, value := range persons {
		fmt.Println(key, value)

		for k, val := range value.iceCreamFlavour {
			fmt.Println(k, val)
		}
	}
}
