package main

import (
	"fmt"
)

type HitungScore func(int, int, int) int

func main() {
	helloName("Tenri", spamFilter)
	helloName("bodoh", spamFilter)

	congrats("tenri", calculateScore)
	congrats("dani", calculateScore)
}

func helloName(a string, b func(string) string) {
	nameFiltered := b(a)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(a string) string {
	if a == "bodoh" {
		return ""
	} else {
		return a
	}
}

//func use type declaration untuk shorten nama function di parameter
func congrats(name string, total HitungScore) {
	var score_akhir int
	if name == "tenri" {
		score_akhir = total(60, 70, 80)
	} else {
		score_akhir = total(50, 60, 70)
	}

	fmt.Println("Halo", name, "Score akhir anda", score_akhir)
}

func calculateScore(ujian int, pr int, kuis int) int {
	return (ujian + pr + kuis) / 3
}
