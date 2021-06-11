package main

import "fmt"

func main() {
	var nama string = "tenri"
	var nama2 string = "tenri2"

	nama3 := "tenri"
	nama4 := "tenri2"

	if nama == nama3 {
		fmt.Println("sama")
	} else {
		fmt.Println("beda")
	}

	samaGak := sama(nama2, nama4)
	fmt.Println(samaGak)

	var pastiFalse bool = nama == nama2
	fmt.Println(pastiFalse)
}

func sama(var1 string, var2 string) bool {
	return var1 == var2
}
