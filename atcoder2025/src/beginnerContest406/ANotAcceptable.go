package main

import (
	"fmt"
)

func ANotAcceptableMain() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	result := "Yes"

	if c > a {
		result = "No"
	} else if a == c {
		if b < d {
			result = "No"
		}
	}
	fmt.Println(result)
}
