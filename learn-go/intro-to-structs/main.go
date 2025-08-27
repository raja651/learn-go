package main

import "fmt"

func main() {

	type person struct {
		name       string
		lastName   string
		age        int
		occupation string
	}

	type secretAgent struct {
		person
		ltk   bool
		alive bool
	}

	agent1 := secretAgent{
		person: person{
			name:       "ladade",
			lastName:   "okay",
			age:        23,
			occupation: "SDE",
		},
		ltk:   false,
		alive: true,
	}

	person2 := person{
		name:       "dark",
		lastName:   "knight",
		age:        100,
		occupation: "vigilante",
	}

	agent2 := secretAgent{
		person: person2,
		ltk:    true,
		alive:  true,
	}

	fmt.Println(agent1)
	fmt.Println(person2)
	fmt.Println(agent2.age, agent2.person, agent2.person.lastName)

}
