package main

import (
	"fmt"
)

func ASwapOddandEvenMain() {
	var s string
	fmt.Scan(&s)
	result := ""
	len := len(s)
	for i := 0; i < len; i = i + 2 {
		sMoji := string(s[i])
		sMojiPlusOne := string(s[i+1])
		result = result + sMojiPlusOne + sMoji
	}
	fmt.Println(result)
}
