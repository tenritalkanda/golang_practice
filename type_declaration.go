package main

import "fmt"

func main() {
	type Tenri string
	type Ten int8

	var varIniJadiString Tenri = "coba"
	var varIniJadiInt Ten = 30

	fmt.Println("stringnya = ", varIniJadiString)
	fmt.Println("angkanya = ", varIniJadiInt)

}
