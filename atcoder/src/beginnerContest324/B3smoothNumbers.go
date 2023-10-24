package main

import (
	"fmt"
)

func B3smoothNumbersMain() {
	var n int64
	fmt.Scan(&n)

	result := "No"
	tempResult := n
	for {
		if tempResult <= 1 {
			result = "Yes"
			break
		}
		if tempResult%2 == 0 {
			tempResult = tempResult / 2
		} else if tempResult%3 == 0 {
			tempResult = tempResult / 3
		} else {
			break
		}
	}

	fmt.Println(result)
}
