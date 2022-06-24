package main

import "time"

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			//Sometimes a goroutine becomes blocked. How can we avoid this to prevent the whole program from blocking? It's simple, we can set a timeout in the select.
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
