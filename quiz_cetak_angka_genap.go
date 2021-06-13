package main

import (
	"fmt"
	"strings"
)

func main() {
	cetak_huruf("tenri", "Genap")
}

func cetak_huruf(kata string, mode string) {
	var divider int

	if strings.ToLower(mode) == "ganjil" {
		divider = 1
	}

	for index, huruf := range kata {
		if index%2 == divider {
			fmt.Println(string(huruf))
		}
	}

	return
}
