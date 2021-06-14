package main

import (
	"fmt"
	"log"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
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

	map2 := map[string]string{
		"key":   "value",
		"kunci": "nilai",
	}

	fmt.Println(map2)

	//map mengembalikan 2 nilai : 1 yg _ itu valuenya dari key, trueOrFalse itu bool dari ada atau tidaknya index dengan key itu
	_, trueOrFalse := map2["ada_gak_key_ini"]
	fmt.Println("ada gak key ini :", trueOrFalse)

	mainUser()
}

func mainUser() {
	userMap := make(map[string]User)

	me := User{
		FirstName: "Tenri",
		LastName:  "Tend",
		Age:       30,
	}

	userMap["who"] = me

	log.Println(userMap["who"].FirstName, userMap["who"].Age)
}
