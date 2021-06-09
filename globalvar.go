package main

import "log"

var s = "one"

func main() {
	log.Println("s is", s)

	s := "two"
	log.Println("now s is", s)

	doubleReturn("hello")

}

func doubleReturn(s2 string) (string, string) {
	log.Println("s inside doublereturn is", s)
	log.Println("s2 is", s2)
	return s2, "hello"
}
