package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CAdjacentSwapsMain() {
	var n, q int
	fmt.Scan(&n, &q)

	var xSlice = make([]int, n)
	xMap := map[int]int{}
	for i := 0; i < n; i++ {
		xSlice[i] = i
		xMap[i] = i
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Scan(&x)
		target := x - 1

		indexValue, _ := xMap[target]

		if indexValue == n-1 {
			left := xSlice[indexValue-1]
			xSlice[indexValue] = left
			xSlice[indexValue-1] = target
			xMap[target] = indexValue - 1
			xMap[left] = indexValue
		} else {
			right := xSlice[indexValue+1]
			xSlice[indexValue] = right
			xSlice[indexValue+1] = target
			xMap[target] = indexValue + 1
			xMap[right] = indexValue
		}
	}

	var resultSlice = make([]string, n)
	for i := 0; i < n; i++ {
		resultSlice[i] = strconv.FormatInt(int64(xSlice[i]+1), 10)
	}
	fmt.Println(strings.Join(resultSlice, " "))
}
