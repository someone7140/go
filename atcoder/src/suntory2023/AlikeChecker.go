package main

import (
	"fmt"
	"strconv"
)

func AlikeCheckerMain() {
	var n string
	fmt.Scan(&n)

	result := "Yes"
	before := -1
	for _, c := range n {
		cMoji := string([]rune{c})
		num, _ := strconv.Atoi(cMoji)
		if num >= before && before >= 0 {
			result = "No"
			break
		}
		before = num
	}

	fmt.Println(result)
}
