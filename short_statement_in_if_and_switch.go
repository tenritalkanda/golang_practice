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
	switch many_cars := len(myCars); many_cars {
	case 2:
		fmt.Println("Many Cars")
	case 1:
		fmt.Println("Not Many Cars")
	default:
		fmt.Println("default")

	}

	switch truefalse := len(myCars); truefalse > 1 {
	case true:
		fmt.Println("Many Cars")
	case false:
		fmt.Println("Not Many Cars")
	default:
		fmt.Println("default")

	}

	name := "Tenri6"
	//switch tanpa expression
	switch {
	case len(name) > 5:
		fmt.Println("nama lebih dari 5 huruf")
	case len(name) > 1 && len(name) <= 5:
		fmt.Println("nama terlalu sedikit")
	default:
		fmt.Println("default nama")
	}

	for index, letter := range name {
		fmt.Println("index :", index, " and letter is :", string(letter))
	}

	for _, huruf := range name {
		fmt.Println("hurufnya :", string(huruf))
	}
}
