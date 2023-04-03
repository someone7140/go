package main

import (
	"fmt"
)

func AAlternatelyMain() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	result := "Yes"
	before := ""
	for i, c := range s {
		sMoji := string([]rune{c})
		if i != 0 {
			if sMoji == before {
				result = "No"
				break
			}
		}
		before = sMoji
	}
	fmt.Println(result)
}
