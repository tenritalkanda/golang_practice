package main

import "fmt"

func main() {
	dict := []string{
		"test",
		"tost",
		"foo",
	}

	is_exists, _ := in_array("foo", dict)
	fmt.Println(is_exists)
}

func in_array(val string, array []string) (ok bool, i int) {
	for i = range array {
		if ok = array[i] == val; ok {
			return
		}
	}
	return
}
