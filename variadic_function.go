package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	max, _ := variadic("max", numbers...)
	fmt.Println(max)

	min, _ := variadic("min", 15, 25, 35, 45, 55)
	fmt.Println(min)
}

//variadic function : func yg memiliki varargs, varargs adalah parameter yg bisa menerima lebih dari 1 data,
// tidak perlu ditentukan jumlah indexnya seperti array [8] saat memanggil function
//numbers ...int adalah varargs harus ditaro dipaling kanan (final argument)
func variadic(operand string, numbers ...int) (int, error) {
	switch operand {
	case "max":
		return max(numbers...), nil
	case "min":
		return min(numbers...), nil
	default:
		return 0, errors.New("Error, operand not available")
	}
}

func max(numbers ...int) (maximum int) {
	for i, number := range numbers {
		if i == 0 {
			maximum = number
		}
		if maximum <= number {
			maximum = number
		}
	}

	return
}

func min(numbers ...int) (minimum int) {
	for i, number := range numbers {
		if i == 0 {
			minimum = number
		}
		if minimum >= number {
			minimum = number
		}
	}

	return
}
