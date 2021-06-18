package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

//panic function adalah function yg digunakan untuk menghentikan program saat ada error muncul, saat panic function dipanggil, defer tetap akan dieksekusi
//kode dibawah panic tidak akan dilanjutkan, karena panic menghentikan eksekusi
func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Eko")
}
