package main

import (
	"fmt"
)

type Developer struct {
	Name  string
	Games []string
}

func (thedev *Developer) AddGame(game string) {
	thedev.Games = append(thedev.Games, game)
}

func main() {
	dev := Developer{"Detik", []string{"bola bekel", "bola sepak"}}

	fmt.Println(dev)

	dev.AddGame("Basket")
	dev.AddGame("Voli")

	fmt.Println(dev)

	for _, game := range dev.Games {
		fmt.Println(game)
	}

}
