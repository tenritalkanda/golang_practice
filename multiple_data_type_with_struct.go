package main

import (
	"fmt"
	"strings"
)

type dog struct {
	name  string
	age   int
	likes []string
}

func main() {
	goodBoy := dog{"Skinny", 4, []string{"Food", "Sleeping"}}
	fmt.Printf("This is %s. It's age is %d and it's favorite things are %s", goodBoy.name, goodBoy.age, strings.Join(goodBoy.likes, ", "))
}
