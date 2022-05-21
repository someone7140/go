package main

import (
	"fmt"
	"strconv"
)

func CSlotStrategyMain() {
	var n int
	fmt.Scan(&n)

	sArrayArray := make([][]int, 10)
	for i := 0; i < 10; i++ {
		sArrayArray[i] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}

	var result int64 = -1

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		for index, c := range s {
			sInt, _ := strconv.Atoi(string(c))
			sArrayArray[sInt][index] = sArrayArray[sInt][index] + 1
		}
	}

	for i := 0; i < 10; i++ {
		var tempResult int64 = 0
		sArray := sArrayArray[i]
		max := -1
		maxIndex := -1
		for j := 0; j < 10; j++ {
			if max <= sArray[j] {
				max = sArray[j]
				maxIndex = j
			}
		}
		tempResult = int64(maxIndex + (max-1)*10)
		if result == -1 || result > tempResult {
			result = tempResult
		}
	}

	fmt.Println(result)
}
