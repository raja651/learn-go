package main

import "fmt"

func main() {
	a := []string{"James", "bond", "shaken, not stirred"}
	b := []string{"Miss", "moneypenny", "I'm, 008"}

	xss := [][]string{a, b}

	for _, v := range xss {
		for _, val := range v {
			fmt.Println(val)
		}
	}

	//map

	m := map[string]int{
		"Henry": 32,
		"Jason": 34,
		"Msd":   23,
	}

	fmt.Printf("%v - \t %#v \n", m, m)
	fmt.Println(m["Henry"])

	ma := make(map[string]int)
	ma["lucas"] = 28
	ma["steph"] = 22
	ma["George"] = 234

	fmt.Printf("%v - \t %#v \n", ma, ma)
	fmt.Println(len(ma))

	for k, v := range ma {
		fmt.Println(k, v)
	}
}
