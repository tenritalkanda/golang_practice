package main

import "fmt"

func main() {
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

	for index := 0; index < len(name); index++ {
		fmt.Println("index :", index, " and letter is :", string(name[index]))
	}

}
