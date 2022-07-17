package main

import (
	"fmt"
)

func AAUniqueLetterMain() {
	var s string
	fmt.Scan(&s)

	first := ""
	second := ""
	third := ""

	for i, c := range s {
		sMoji := string([]rune{c})
		if i == 0 {
			first = sMoji
		} else if i == 1 {
			second = sMoji
		} else {
			third = sMoji
		}
	}

	if first != second && second != third && first != third {
		fmt.Println(first)
	} else if first == second && second != third {
		fmt.Println(third)
	} else if first != second && second == third {
		fmt.Println(first)
	} else if first == third && second != third {
		fmt.Println(second)
	} else {
		fmt.Println(-1)
	}

}
