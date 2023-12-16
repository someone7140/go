package main

import (
	"fmt"
	"strconv"
)

func CRepunitTrioMain() {
	var n int
	fmt.Scan(&n)

	count := 1
	var result int64
	result = 3

	startCount := 1
	secondCount := 1
	thirdCount := 1

	countCalc := func(count int) int64 {
		calcStr := ""
		for i := 0; i < count; i++ {
			calcStr = calcStr + "1"
		}
		calcInt64, _ := strconv.ParseInt(calcStr, 10, 64)
		return calcInt64
	}

	for {
		if count == n {
			if count > 1 {
				result = countCalc(startCount) + countCalc(secondCount) + countCalc(thirdCount)
			}
			break
		}

		if startCount == secondCount && secondCount == thirdCount {
			startCount = startCount + 1
			secondCount = 1
			thirdCount = 1
		} else {
			if thirdCount == secondCount {
				secondCount = secondCount + 1
				thirdCount = 1
			} else {
				thirdCount = thirdCount + 1
			}
		}
		count = count + 1
	}
	fmt.Println(result)
}
