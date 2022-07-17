package main

import (
	"fmt"
)

func CChangingJewelsMain() {
	var n, x, y int64
	fmt.Scan(&n, &x, &y)

	redCount := int64(1)
	blueCount := int64(0)

	for i := n; i > 1; i-- {
		// 赤い処理
		blueCount = blueCount + redCount*x
		// 青い処理
		redCount = redCount + blueCount
		blueCount = blueCount * y
	}

	fmt.Println(blueCount)
}
