package main

import (
	"fmt"
)

func AFirstContestoftheYearMain() {
	var d, f int
	fmt.Scan(&d, &f)
	result := 7 - (d-f)%7

	fmt.Println(result)
}
