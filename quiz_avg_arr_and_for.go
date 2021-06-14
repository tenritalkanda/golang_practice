package main

import "fmt"

func main() {
	nilai := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	rata2 := average(nilai)
	fmt.Println("rata2 = ", rata2)

	nilai2 := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	biggerThan := 90
	goodScores := valuesBiggerThan(nilai2, biggerThan)

	fmt.Println(goodScores)
}

func average(arr [8]int) int8 {
	var result int
	for _, i := range arr {
		result += i
	}

	return int8(result / 8)
}

func valuesBiggerThan(arr [8]int, pembanding int) []int {
	var result []int
	fmt.Println(pembanding)
	for _, v := range arr {
		if v >= pembanding {
			result = append(result, v)
		}
	}
	return result
}
