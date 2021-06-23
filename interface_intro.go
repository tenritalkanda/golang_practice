package main

import "fmt"

//interface type data yg isinya berupa method
type Absen interface {
	Hello() string
}

type Student struct {
	Name string
}

func (student Student) Hello() string {
	return "Hello " + student.Name
}

type Teacher struct {
	Name   string
	Gender string
}

func (teacher Teacher) Hello() string {
	if teacher.Gender == "male" {
		return "Terima kasih pak" + teacher.Name
	} else {
		return "Terima kasih bu " + teacher.Name
	}
}

//func dengan parameter interface
func Greeting(absen Absen) string {
	return absen.Hello()
}

func main() {
	guru := Teacher{
		Name:   "Sinta",
		Gender: "female",
	}

	//function greeting bisa pakai struct teacher
	greet := Greeting(guru)

	fmt.Println(greet)

	murid := Student{
		Name: "Tenri",
	}

	//function greeting bisa pakai struct murid
	greet2 := Greeting(murid)

	fmt.Println(greet2)
}
