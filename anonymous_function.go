package main

import (
	"fmt"
)

// function tanpa nama, seperti di js dll, function langsung di declare di sebuah variable
func main() {
	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Println(multiply(100, 5))
}
