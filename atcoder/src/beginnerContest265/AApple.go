package main

import (
	"fmt"
)

func AAppleMain() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)

	xPriorityFlag := true
	ySyou := y / 3

	if x > ySyou {
		xPriorityFlag = false
	}

	result := 0

	if xPriorityFlag || n < 3 {
		result = x * n
	} else {
		n3Syou := n / 3
		result = result + n3Syou*y
		result = result + x*(n-n3Syou*3)
	}
	fmt.Println(result)
}
