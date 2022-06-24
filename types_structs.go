package main

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

var a int

type B struct {
	Test string
}

func main() {
	user := User{
		FirstName: "test",
	}

	bb := B{Test: "waw"}
	log.Println(user.FirstName, user.LastName, user.PhoneNumber, user.Age, user.Birthdate)
	log.Println(a)
	log.Println(bb.Test)

	fmt.Println(user.FirstName, user.LastName, user.PhoneNumber, user.Age, user.Birthdate)
	// test := test()
	// log.Println(test)
}
