package main

import (
	"fmt"
)

func AcamelCaseMain() {
	var s string
	fmt.Scan(&s)

	upperA := rune('A')
	upperZ := rune('Z')

	var index = 0
	for i, c := range s {
		if c >= upperA && c <= upperZ {
			index = i
		}
	}

	fmt.Println(index + 1)
}
