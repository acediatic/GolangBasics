package main

import "fmt"

func maps() {
	// maps are a key value pair
	// maps are unordered
	// maps are reference types
	// maps are not thread safe

	// can use make to create empty map first
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}

	// Return order of elements is not guaranteed
	fmt.Println(statePopulations)

	// delete from map
	delete(statePopulations, "California")

	fmt.Println(statePopulations["California"]) // 0, not error

	// check if key exists
	population, ok := statePopulations["California"]
	fmt.Println(population, ok) // 0, false

	_, ok = statePopulations["Texas"]
	fmt.Println(ok) // true

}
