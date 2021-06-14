package main

import "fmt"

func main() {
	students := []map[string]string{
		{"Nama": "Dina", "Nilai": "B"},
		{"Nama": "Andien", "Nilai": "B+"},
		{"Nama": "Edna", "Nilai": "A-"},
	}

	for _, student := range students {
		fmt.Println("Nama :", student["Nama"], "memiliki nilai :", student["Nilai"])
	}
}
