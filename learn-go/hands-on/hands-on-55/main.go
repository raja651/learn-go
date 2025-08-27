package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make   string
	modal  string
	doors  int
	colour string
}

func main() {

	nissan := vehicle{
		engine: engine{
			electric: false,
		},
		make:   "Japan",
		modal:  "Gtx",
		doors:  4,
		colour: "Red",
	}

	kia := vehicle{
		engine: engine{
			electric: false,
		},
		make:   "India",
		modal:  "carens",
		doors:  5,
		colour: "Black",
	}

	fmt.Println(nissan)
	fmt.Println(kia)

	fmt.Println(nissan.electric, nissan.modal)
	fmt.Println(kia.engine.electric, kia.modal)
}
