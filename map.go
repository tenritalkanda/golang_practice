package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Blackie"
	myMap["dog2"] = "Cassie"

	myMap2 := make(map[string]int)

	myMap2["age"] = 1

	log.Println(myMap["dog"])
	log.Println(myMap["dog2"])

	log.Println(myMap2["age"])

	mainUser()
}

func mainUser() {
	userMap := make(map[string]User)

	me := User{
		FirstName: "Tenri",
		LastName:  "Tend",
	}

	userMap["who"] = me

	log.Println(userMap["who"].FirstName)
}