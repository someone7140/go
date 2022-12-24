package main

import (
	"fmt"
)

func ARightmostMain() {
	var s string
	fmt.Scan(&s)

	result := -1
	runes := []rune(s)
	for j := len(runes) - 1; 0 <= j; j = j - 1 {
		sMoji := string([]rune{runes[j]})
		if sMoji == "a" {
			result = j + 1
			break
		}
	}
	fmt.Println(result)

}
