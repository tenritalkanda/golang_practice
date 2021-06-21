package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := a
	fmt.Println(a)
	b[0] = 10

	fmt.Println(a, b)

	c := []int{1, 2}
	d := c
	fmt.Println(c)
	d[0] = 10

	fmt.Println(c, d)
}
