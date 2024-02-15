package main

import (
	"fmt"
)

func CDivideandDivideMain() {
	var n int64
	fmt.Scan(&n)

	tennsuuMap := map[int64]int64{}
	// まずは2から200000まで計算結果を格納する
	tennsuuMap[2] = 2
	for i := int64(3); i <= 200000; i++ {
		half1 := i / 2
		var half2 int64
		if i%2 == 0 {
			half2 = i / 2
		} else {
			half2 = half1 + 1
		}

		tensuu := i
		if half1 > 1 {
			tensuu = tensuu + tennsuuMap[half1]
		}
		if half2 > 1 {
			tensuu = tensuu + tennsuuMap[half2]
		}
		tennsuuMap[i] = tensuu
	}

	v, ok := tennsuuMap[n]
	if ok {
		fmt.Println(v)
	} else {
		// 対象の数を割っていく
		var wariFunc func(target int64) int64
		wariFunc = func(target int64) int64 {
			half1 := target / 2
			var half2 int64
			if target%2 == 0 {
				half2 = target / 2
			} else {
				half2 = half1 + 1
			}

			var half1Res int64
			v1, ok1 := tennsuuMap[half1]
			if !ok1 {
				half1Res = wariFunc(half1)
				tennsuuMap[half1] = half1Res
			} else {
				half1Res = v1
			}

			var half2Res int64
			v2, ok2 := tennsuuMap[half2]
			if !ok2 {
				half2Res = wariFunc(half2)
				tennsuuMap[half2] = half2Res
			} else {
				half2Res = v2
			}
			return half1Res + half2Res + target
		}

		fmt.Println(wariFunc(n))
	}

}
