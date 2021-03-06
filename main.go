package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World")
	//multiple declare with one var
	var (
		whatToSay string
		say       string
		i         int
	)

	var doubleD string
	var secondD string

	whatToSay = saySomething("Hello")
	say = saySomething("Guys")

	log.Println(whatToSay)
	log.Println(say)
	log.Println(saySomething("There"))

	log.Println(i)

	i = 7

	log.Println(i)

	doubleD, _ = doubleReturn("Hello")
	log.Println(doubleD)

	doubleD, secondD = doubleReturn("Wow")
	log.Println(doubleD, secondD)

	//switch case
	myCat := "red"

	switch myCat {
	case "black":
		log.Println("i got black cat")
	case "blue":
		log.Println("there is no blue cat")
	default:
		log.Println("default")
	}

	//length of string
	log.Println(len("satu_dua"))
}

func saySomething(s string) string {
	return s
}

func doubleReturn(s string) (string, string) {
	return s, "hello"
}
