package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Caca", "Edo")
	mySlice = append(mySlice, "Budi", "Anto")

	// log.Println(mySlice)

	sort.Strings(mySlice)

	myNumbers := []int{19, 5, 8, 9, 10, 7, 6, 18}

	sort.Ints(myNumbers)

	log.Println(mySlice)

	log.Println(myNumbers)

	//slice from slice
	days := []string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	//string from slice
	daySlice := days[5]

	fmt.Println("days :", days)
	fmt.Println("string from slice :", daySlice)

	daySlice = "awaw"
	fmt.Printf("daySlice berubah dari '%s' ke : '%s'", days[5], daySlice)
	fmt.Println("\nparentnya tetep :", days)

	//slice from slice
	daySliceFR := days[0:2]
	fmt.Println("\nslice from slice 0:2 :", daySliceFR)

	//slice from slice kalo dirubah value, parentnya ikut berubah
	daySliceFR[0] = "Berubah"
	fmt.Println("\ndaySliceFR berubah :", daySliceFR)
	fmt.Println("Parentnya juga berubah :", days)

	//slice from slice tapi parentnya tidak ikut berubah
	daySliceTwo := days[0:2]
	fmt.Println("\ndaySliceTwo :", daySliceTwo)

	//karena daySliceTwo fix berkapasitas 2,
	//di append dari new declare, maka akan terbentuk array baru yg tidak berhubungan dengan parentnya saat berubah.
	//karena melebihi kapasitas si parent
	dayAppend := append(daySliceTwo, "LEBIH")
	fmt.Println("\ndayAppend :", dayAppend)

	fmt.Println("\ndaySliceTwo kapasitas 2 jadi tidak ada array 'LEBIH' :", daySliceTwo)
	fmt.Println("\ndays kapasitasnya 7 jadi bisa ikut berubah di index ke 3:", days)

}
