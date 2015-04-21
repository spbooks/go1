package main

import "fmt"

func main() {
	starWarsYears := map[string]int{
		"A New Hope":              1977,
		"The Empire Strikes Back": 1980,
		"Return of the Jedi":      1983,
		"Attack of the Clones":    2002,
		"Revenge of the Sith":     2005,
	}
	starWarsYears["The Force Awakens"] = 2015

	for title, year := range starWarsYears {
		fmt.Println(title, "was released in", year)
	}
}
