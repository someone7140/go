package main

import (
	"fmt"
)

func AwwwvvvvvvMain() {
	var s string
	fmt.Scan(&s)

	result := 0
	for _, c := range s {
		sMoji := string([]rune{c})
		if sMoji == "v" {
			result = result + 1
		} else {
			result = result + 2
		}
	}
	fmt.Println(result)
}
