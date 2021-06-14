package main

import "fmt"

var count int
var dogs []string

func main() {
	count, dogs = getDogs()

	fmt.Println(count)
	fmt.Println(dogs)
}

func getDogs() (count int, dogs []string) {
	count = 5
	dogs = []string{
		"Snoopy",
		"Skinny",
		"Scooby",
		"Santos L. Halper",
	}

	return
}
