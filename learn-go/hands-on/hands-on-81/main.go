package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type sortByAge []user
type sortByLastName []user

func (a sortByAge) Len() int {
	return len(a)
}

func (a sortByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a sortByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (b sortByLastName) Len() int {
	return len(b)
}

func (b sortByLastName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b sortByLastName) Less(i, j int) bool {
	return b[i].Last < b[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, val := range users {
		fmt.Println(val)
	}

	// your code goes here

	fmt.Println("----------------------------------")

	sort.Sort(sortByAge(users))

	for _, val := range users {
		fmt.Println(val.First, val.Last, val.Age)
		sort.Strings(val.Sayings)
		fmt.Println(val.Sayings)
	}

	fmt.Println("----------------------------------")

	sort.Sort(sortByLastName(users))

	for _, val := range users {
		fmt.Println(val.First, val.Last, val.Age)
		sort.Strings(val.Sayings)
		fmt.Println(val.Sayings)
	}
}
