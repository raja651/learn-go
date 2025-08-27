package main

import "fmt"

func main() {
	names := map[string][]string{
		`bond_james`:       {`sdasd`, `shaken`, `not stirred`, `martinis`},
		`moneypenny_jenny`: {`intelligence`, `gut`, `bold`, `sdasv`},
		`mugilan_krithu`:   {`baby`, `cute`, `love`},
	}

	names["daddf"] = []string{`sdasd`, `shaken`, `not stirred`, `drgm`}
	names["gofj"] = []string{`dsd`, `moermj`, `mfsoo`}

	for key, value := range names {
		fmt.Printf("%#v  \t %v \n", key, value)
		for _, val := range value {
			fmt.Printf("%v \n", val)
		}
	}

	// delete
	delete(names, "gofj")

	for key, value := range names {
		fmt.Printf("%#v  \t %v \n", key, value)
		for _, val := range value {
			fmt.Printf("%v \n", val)
		}
	}

}
