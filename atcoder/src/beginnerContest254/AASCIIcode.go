package main

import (
	"fmt"
)

func AASCIIcodeMain() {
	var n string
	fmt.Scan(&n)

	result := ""
	for i, c := range n {
		if i > 0 {
			nMoji := string([]rune{c})
			result = result + nMoji
		}
	}
	fmt.Println(result)
}
