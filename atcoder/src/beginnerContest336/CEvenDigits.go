package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int64
	fmt.Scan(&n)

	// まず何桁か判定する
	keta := 1
	startCount := int64(0)
	for {
		var temp int64
		if keta == 1 {
			temp = int64(5)
		} else {
			temp = 1
			for i := 0; i < keta-1; i++ {
				temp = temp * 5
			}
			temp = startCount + temp*4
		}

		if n <= temp {
			break
		} else {
			keta = keta + 1
			startCount = temp
		}
	}

	// 答えの配列
	resultArray := make([]int, keta)
	if keta > 1 {
		resultArray[0] = 2
	}
	numberCount := startCount

	// 一桁判定用
	judgeHitoketa := func(judgeNumber int64) int {
		return 2 * (int(judgeNumber) - 1)
	}
	// 複数桁判定用
	var judgeLoop func(judgeIndex int, ueNumber int) bool
	judgeLoop = func(judgeIndex int, ueNumber int) bool {
		if judgeIndex == 1 {
			sabun := n - numberCount - int64(ueNumber/2*5)
			if sabun <= 5 && sabun > 0 {
				resultArray[keta-1] = judgeHitoketa(sabun)
				return true
			} else {
				return false
			}
		} else {
			var tempCount int64
			tempCount = 1
			for i := 0; i < judgeIndex-1; i++ {
				tempCount = tempCount * 5
			}
			tempCount = numberCount + tempCount*4
			if tempCount >= n {
				// 数字を決めていく
				for i := 0; i <= 8; i = i + 2 {
					result := judgeLoop(judgeIndex-1, i)
					if result {
						resultArray[keta-judgeIndex] = i
						break
					}
				}
				return true
			} else {
				numberCount = tempCount
				return false
			}
		}
	}

	if keta == 1 {
		resultArray[0] = judgeHitoketa(n)
	} else {
		// 最大桁でループ
		for i := 2; i <= 8; i = i + 2 {
			endFlag := judgeLoop(keta-1, i)
			if endFlag {
				resultArray[0] = i
				break
			}
		}
	}

	resultStrArray := make([]string, keta)
	for i := 0; i < keta; i++ {
		resultStrArray[i] = strconv.FormatInt(int64(resultArray[i]), 10)
	}

	fmt.Println(strings.Join(resultStrArray, ""))
}
