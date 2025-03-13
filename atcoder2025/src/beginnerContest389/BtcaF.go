package main

import (
	"fmt"
)

func BtcaFMain() {
	var n int64
	fmt.Scan(&n)

	resultCount := int64(1)
	temp := int64(1)

	for {
		if temp >= n {
			break
		}
		resultCount = resultCount + 1
		temp = temp * resultCount
	}
	fmt.Println(resultCount)

}
