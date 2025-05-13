package main

import (
	"fmt"
)

func AIsitratedMain() {
	var r, x int
	fmt.Scan(&r, &x)

	result := ""

	if x == 1 {
		if r >= 1600 && r <= 2999 {
			result = "Yes"
		} else {
			result = "No"
		}
	} else {
		if r >= 1200 && r <= 2399 {
			result = "Yes"
		} else {
			result = "No"
		}
	}

	fmt.Println(result)
}
