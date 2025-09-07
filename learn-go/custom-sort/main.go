package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name     string
	Age      int
	Location string
}

type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

type ByName []Person

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func main() {

	people := []Person{
		{Name: "Darkhorse", Age: 34, Location: "Chennai"},
		{Name: "WhiteHorse", Age: 23, Location: "Vellore"},
		{Name: "RedHorse", Age: 56, Location: "Banglr"},
		{Name: "GreenHorse", Age: 12, Location: "Villupuran"},
	}

	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
