package main

import (
	"fmt"
)

// a very simple function that we'll
// make asynchronous later on
func compute(value int) {
	for i := 0; i < value; i++ {
		// time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func sum(a []int, c chan int) {
	total := 0
	fmt.Println(a)
	for _, v := range a {
		total += v
	}
	fmt.Println(total)
	c <- total // send total to channel c
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// sequential execution of our compute function
	go compute(3)
	go compute(3)

	//communicate between different goroutines
	a := []int{7, 2, 8, -9, 4, 0}

	//Sending and receiving data in channels blocks by default, so it's much easier to use synchronous goroutines.
	//a goroutine will not continue when receiving data from an empty channel, i.e (value := <-ch), until other goroutines send data to this channel.
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c assign to x and y

	fmt.Println(x, y, x+y)

	// we scan fmt for input and print that to our console
	var input string
	fmt.Scanln(&input)
}
