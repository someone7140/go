package main

import (
	"fmt"
)

func BEnlargedCheckerBoardMain() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	var resultSlice []string

	blackFlag1 := false
	for i := 0; i < n; i++ {
		blackFlag1 = i%2 != 0
		for i2 := 0; i2 < a; i2++ {
			blackFlag2 := false
			tempResult := ""
			for j := 0; j < n; j++ {
				if blackFlag1 {
					blackFlag2 = j%2 == 0
				} else {
					blackFlag2 = j%2 != 0
				}
				for j2 := 0; j2 < b; j2++ {
					if blackFlag2 {
						tempResult = tempResult + "#"
					} else {
						tempResult = tempResult + "."
					}
				}
			}
			resultSlice = append(resultSlice, tempResult)
		}

	}

	for _, v := range resultSlice {
		fmt.Println(v)
	}

}
