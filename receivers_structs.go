package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "marie",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	log.Println("myVar called from func receiver", myVar.printFirstName())
	log.Println("myVar2 called from func receiver", myVar2.printFirstName())
}
