package main

import (
	"fmt"
)

// Maps in Go:
/*
* Go's map is a hashmap in computer science terms.
 */

func main() {

	// creating a nil map
	var nilMap = map[string]int{}

	fmt.Println(nilMap)      // map[]
	fmt.Println(len(nilMap)) // 0

	// read and write to nilMap
	nilMap["one"] = 1
	nilMap["two"] = 2
	nilMap["three"] = 3

	fmt.Println(nilMap["three"]) // 3

	// Non empty map
	teams := map[string][]string{
		"Orcas": {
			"Fred", "Ralph", "Bijou",
		},
		"Lions": {
			"Sarah", "Peter", "Billie",
		},
		"Kittens": {
			"Waldo", "Raul", "Ze",
		},
	}

	fmt.Println(teams)

	// If we know how many key-value pairs we intend to put in the map but don't know the exact values, we can use make to create a map with a default size:
	newTeams := make(map[string][]string, 4)

	// Add entries
	newTeams["Orcas"] = []string{"Fred", "Ralph", "Bijou"}
	newTeams["Lions"] = []string{"Sarah", "Peter", "Billie"}
	newTeams["Kittens"] = []string{"Waldo", "Raul", "Ze"}

	fmt.Println(newTeams)

	getArgumentsFromTerminalAndPrintUser()

	fmt.Println(slicingTheWeek())

	fmt.Println(deleteUser("305"))

}
