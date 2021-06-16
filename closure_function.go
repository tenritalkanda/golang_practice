package main

import (
	"fmt"
)

//closure adalah kemampuan function untuk mengubah variable diluar function itu
//ex: func mengubah value counter, diluar func, counter ikut berubah
func main() {
	counter := 0
	multiply := func(a int, b int) {
		for a > b {
			a -= b
			counter++
			fmt.Println("nilai sekarang :", a, "counter dalam anon func", counter)
		}
		return
	}

	multiply(100, 5)
	fmt.Println("counter diluar anon func ", counter)
}
