package main

import (
	"fmt"
)

func BSlimesMain() {
	var a, b, k int64
	fmt.Scan(&a, &b, &k)

	result := 0
	temp := a

	for {
		if temp >= b {
			break
		} else {
			temp = temp * k
			result = result + 1
		}
	}
	fmt.Println(result)
}
