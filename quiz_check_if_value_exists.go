package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	dict := []string{
		"test",
		"tost",
		"foo",
	}

	is_exists, _ := in_array("foo", dict)
	fmt.Println(is_exists)
	fmt.Println(runtime.GOOS)

	today := time.Now()
	fmt.Println(today)
}

func in_array(val string, array []string) (ok bool, i int) {
	for i = range array {
		if ok = array[i] == val; ok {
			return
		}
	}
	return
}
