package main

import "fmt"

func main() {
	var ini_array = [...]int{1, 2, 3}
	var ini_slice = []int{4, 5, 6}

	fmt.Println("ini array : ", ini_array)
	fmt.Println("ini slice : ", ini_slice)

	//declare slice pakai make([]type, len, cap)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "satu"
	newSlice[1] = "dua"

	fmt.Println(newSlice)

	//kode newSlice[2] akan error, karena len array cuma 2 (index 0 dan index 1)
	//newSlice[2] = "tiga"
	//fmt.Println(newSlice)

	//kode dibawah tidak akan error, karena append bisa menambahkan selama masih dalam kapasitas 5.
	newSlice = append(newSlice, "test", "tost")
	fmt.Println(newSlice)
	fmt.Println("length nya :", len(newSlice))
}
