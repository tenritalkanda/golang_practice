package main

import "fmt"

func main() {
	var satu int32 = 100000
	var dua int64 = int64(satu)
	var tiga int8 = int8(satu)

	//tiga karena tipe data int8 tidak bisa memuat angka 100ribu. maka hasilnya akan loop kembali dari angka terendahnya
	fmt.Println(satu, dua, tiga)

	//konversi dari string ke byte, balik lagi ke string
	var nama = "Tenri"
	var test byte = nama[0]
	eString := string(test)

	fmt.Println(test, eString)

	//test range
	test_range := nama[0:3]
	fmt.Println(test_range)
}
