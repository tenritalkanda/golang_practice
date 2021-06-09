package main

import "log"

func main() {
	// fmt.Println("Hello World")
	var whatToSay string
	var say string
	var i int
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

}

func saySomething(s string) string {
	return s
}

func doubleReturn(s string) (string, string) {
	return s, "hello"
}
