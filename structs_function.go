package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(greetType string) {
	if greetType == "morning" {
		fmt.Println("Hello", customer.Name, "do you want breakfast?")
	} else {
		fmt.Println("Hello", customer.Name, "what can i do for you?")
	}

}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	myCustomer := Customer{
		Name:    "Tenri",
		Address: "Indonesia",
		Age:     30,
	}

	myCustomer.sayHi("morning")
	myCustomer.sayHi("else")
	myCustomer.sayHuuu()

}
