package main

import (
	"fmt"
)

func AatcodersubstrMain() {
	var l, r int
	fmt.Scan(&l, &r)

	result := ""
	atcoderStr := "atcoder"

	for i, c := range atcoderStr {
		if i+1 >= l && i+1 <= r {
			nMoji := string([]rune{c})
			result = result + nMoji
		}
	}
	fmt.Println(result)
}
