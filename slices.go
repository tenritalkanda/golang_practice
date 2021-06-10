package main

import (
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

	log.Println(myNumbers[0:5])

	log.Println(myNumbers[1:1])

	log.Println(myNumbers[1:2])

	log.Println(myNumbers[6:8])

	log.Println(myNumbers[7])
}
