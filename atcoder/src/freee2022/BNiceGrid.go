package main

import (
	"fmt"
)

func BNiceGridMain() {
	var r, c int
	fmt.Scan(&r, &c)

	var r2, c2 int
	if r > 9 {
		r2 = 16 - r
	} else {
		r2 = r
	}
	if c > 9 {
		c2 = 16 - c
	} else {
		c2 = c
	}

	var min int
	if r2 > c2 {
		min = c2
	} else {
		min = r2
	}

	result := ""
	if min%2 == 1 {
		result = "black"
	} else {
		result = "white"
	}

	fmt.Println(result)
}
