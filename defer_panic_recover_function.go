package main

import (
	"fmt"
)

func runApplication(b int) {
	//defer adalah function yg akan dicall setelah function yg sedang dieksekusi selesai, walaupun error. defer harus ditaro diawal line
	//panic function adalah function yg digunakan untuk menghentikan program, saat panic function dipanggil, defer tetap akan dieksekusi
	defer loggingUser()
	for i := 10; i >= 1; i-- {
		fmt.Println("i skrg :", i)
		if i == 1 {
			a := i / b
			fmt.Println(a)
		}
	}
}

func loggingUser() {
	fmt.Println("run App")
}

func main() {
	runApplication(0)
}
