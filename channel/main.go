package main

import (
	"fmt"
	"time"
)

/*for i := range d will not stop reading data from channel until the channel is closed.
We use the keyword close to close the channel in example.
It's impossible to send or receive data on a closed channel;
you can use v, ok := <-ch to test if a channel is closed.
If ok returns false, it means the there is no data in that channel and it was closed.
Remember to always close channels in producers and not in consumers, or it's very easy to get into panic status.

Another thing you need to remember is that channels are not like files.
You don't have to close them frequently unless you are sure the channel is completely useless,
or you want to exit range loops.
*/
func fibonacci(n int, d chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		if i == 5 {
			time.Sleep(time.Second * 5)
		}
		d <- x
		x, y = y, x+y
	}
	close(d)
}

// buffered channels that can store more than a single element.
func main() {
	c := make(chan int, 2) // change 2 to 1 will have runtime error, but 3 is fine
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)

	d := make(chan int, 10)
	go fibonacci(cap(d), d)
	for i := range d {
		fmt.Println(i)
	}
}
