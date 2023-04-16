package main

import (
	"fmt"
)

func AJobInterviewMain() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	yesCount := 0
	noCount := 0
	for _, c := range s {
		sMoji := string([]rune{c})
		if sMoji == "o" {
			yesCount = yesCount + 1
		}
		if sMoji == "x" {
			noCount = noCount + 1
		}
	}
	if noCount > 0 {
		fmt.Println("No")
	} else if yesCount > 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
