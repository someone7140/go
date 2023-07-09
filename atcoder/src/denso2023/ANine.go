package main

import (
	"fmt"
)

func ANineMain() {
	var a, b int
	fmt.Scan(&a, &b)
	result := "No"

	if a == 1 && b == 2 {
		result = "Yes"
	} else if a == 2 && b == 3 {
		result = "Yes"
	} else if a == 4 && b == 5 {
		result = "Yes"
	} else if a == 5 && b == 6 {
		result = "Yes"
	} else if a == 7 && b == 8 {
		result = "Yes"
	} else if a == 8 && b == 9 {
		result = "Yes"
	}

	fmt.Println(result)

}
