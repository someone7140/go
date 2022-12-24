package main

import (
	"fmt"
)

func CCashRegisterMain() {
	var s string
	fmt.Scan(&s)

	sLen := len(s)
	sArray := make([]string, sLen)

	for i, c := range s {
		sMoji := string([]rune{c})
		sArray[i] = sMoji
	}

	var result int64
	maeZero := false
	for i := 0; i < sLen; i++ {
		if sArray[i] != "0" {
			result = result + 1
			maeZero = false
		} else {
			if maeZero {
				maeZero = false
			} else {
				result = result + 1
				maeZero = true
			}
		}
	}

	fmt.Println(result)
}
