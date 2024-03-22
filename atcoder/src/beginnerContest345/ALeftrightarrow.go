package main

import (
	"fmt"
)

func ALeftrightarrowMain() {
	var s string
	fmt.Scan(&s)

	result := "Yes"
	sLen := len(s)
	for i, c := range s {
		sMoji := string([]rune{c})
		if i == 0 {
			if sMoji != "<" {
				result = "No"
				break
			}
		} else if i == sLen-1 {
			if sMoji != ">" {
				result = "No"
				break
			}
		} else {
			if sMoji != "=" {
				result = "No"
				break
			}
		}
	}

	fmt.Println(result)
}
