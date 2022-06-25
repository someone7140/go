package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BPracticalComputingMain() {
	var n int
	fmt.Scan(&n)

	resultArrayArray := make([][]int, n)

	for i := 0; i < n; i++ {
		var resultSlice []int
		if i == 0 {
			resultSlice = append(resultSlice, 1)
		} else if i == 1 {
			resultSlice = append(resultSlice, 1)
			resultSlice = append(resultSlice, 1)
		} else {
			before := resultArrayArray[i-1]
			tempI := i + 1
			for j := 0; j < tempI; j++ {
				if j == 0 || j == tempI-1 {
					resultSlice = append(resultSlice, 1)
				} else {
					resultSlice = append(resultSlice, before[j-1]+before[j])
				}
			}
		}
		resultArrayArray[i] = resultSlice
	}

	for i := 0; i < n; i++ {
		resultInts := resultArrayArray[i]
		var resultSlice []string
		for _, resultInt := range resultInts {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(resultInt), 10))
		}

		fmt.Println(strings.Join(resultSlice, " "))
	}
}
