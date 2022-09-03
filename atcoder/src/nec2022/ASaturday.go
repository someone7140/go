package main

import (
	"fmt"
)

func ASaturdayMain() {
	var s string
	fmt.Scan(&s)

	result := 0
	if s == "Monday" {
		result = 5
	} else if s == "Tuesday" {
		result = 4
	} else if s == "Wednesday" {
		result = 3
	} else if s == "Thursday" {
		result = 2
	} else {
		result = 1
	}
	fmt.Println(result)
}
