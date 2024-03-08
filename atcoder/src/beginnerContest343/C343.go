package main

import (
	"fmt"
)

func C343Main() {
	var n int64
	fmt.Scan(&n)

	var i int64
	var result int64

	var isPalindrome func(nInput int64) bool
	isPalindrome = func(nInput int64) bool {
		nTemp := nInput
		var rev int64
		for {
			if nTemp < 1 {
				break
			}
			r := nTemp % 10
			rev = rev*10 + r
			nTemp = nTemp / 10
		}
		return (rev == nInput)
	}

	for i = 1000000; i >= 1; i-- {
		temp := i * i * i
		if temp <= n && isPalindrome(temp) {
			result = temp
			break
		}
	}
	fmt.Println(result)
}
