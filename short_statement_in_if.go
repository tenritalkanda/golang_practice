package main

import "fmt"

func main() {
	myCars := make(map[string]string)
	myCars["satu"] = "Honda"
	myCars["dua"] = "Toyota"

	//short statement = bikin variable pakai tanda semicolon ; lalu di compare di if
	if length_cars := len(myCars); length_cars >= 1 {
		fmt.Println("Car exists")
	} else {
		fmt.Println("No Cars")
	}

	//short statement in switch
	switch many_cars := len(myCars); many_cars >= 1 {
	case true:
		fmt.Println("Many Cars")
	case false:
		fmt.Println("Not Many Cars")
	default:
		fmt.Println("default")

	}
}
