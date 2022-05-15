package main

import (
	"fmt"
	"math"
	"strconv"
)

func D250likeNumberMain() {
	var nStr string
	fmt.Scan(&nStr)

	n, _ := strconv.ParseInt(nStr, 10, 64)

	sosuuList := getSosuuListABC250(int64(math.Cbrt(float64(n))) + 100)
	sosuuListSize := len(sosuuList)
	var result int64
	result = 0

	if n > 50 {
		for i := 0; i < (sosuuListSize - 1); i++ {
			leftvalue := sosuuList[i]
			rightValue := n / leftvalue
			rightValue3 := math.Cbrt(float64(rightValue))
			index := nibunABC250(sosuuList, i, sosuuListSize-1, int64(rightValue3))
			if index != -1 {
				result = result + int64(index-i)
			}
		}
	}

	fmt.Println(result)
}

func getSosuuListABC250(target int64) []int64 {
	var sosuuSlice []int64

	if target >= 2 {
		sosuuSlice = append(sosuuSlice, 2)
	}
	if target >= 3 {
		sosuuSlice = append(sosuuSlice, 3)
	}

	if target >= 5 {
		var n int64
	L:
		for n = 5; n <= target; n += 2 {
			i := 1                                 //素数のスライスの中を走査する添字。prime[1]の3から
			for sosuuSlice[i]*sosuuSlice[i] <= n { //prime[i]の2乗がn以下であるか。つまり、nの平方根以下であるかの確認。乗算を利用して判定する。
				if n%sosuuSlice[i] == 0 { //割り切れると素数では無い
					continue L //この素数で割るループを抜ける
				}
				i++
			}
			//最後まで割り切れなかったら
			sosuuSlice = append(sosuuSlice, n)
		}
	}

	return sosuuSlice
}

func nibunABC250(arrayInput []int64, startInput int, endInput int, x int64) int {

	if x < arrayInput[startInput] {
		return -1
	} else if x > arrayInput[endInput] {
		return -1
	} else {
		start := startInput
		end := endInput
		half := (start + end) / 2
		for {
			if start >= half || end <= half {
				break
			}
			if arrayInput[half] <= x {
				start = half
				half = (start + end) / 2
			} else if arrayInput[half] >= x {
				end = half
				half = (start + end) / 2
			}
		}
		return half
	}
}
