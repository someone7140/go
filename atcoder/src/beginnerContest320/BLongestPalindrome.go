package main

import (
	"fmt"
)

func BLongestPalindromeMain() {
	var s string
	fmt.Scan(&s)

	var reverseFuncStr func(str string) (result string)
	reverseFuncStr = func(str string) (result string) {
		for _, v := range str {
			result = string(v) + result
		}
		return
	}
	sReverse := reverseFuncStr(s)

	result := -1
	lenS := len(s)
	setStr := make(map[string]struct{})

	var judgeResult func(str string, strReverse string, index int, reverse bool)
	judgeResult = func(str string, strReverse string, index int, reverse bool) {
		if index == lenS {
			return
		}

		if !reverse {
			tempS1 := s[:lenS-index]
			tempS1Reverse := strReverse[1:]
			_, ok1 := setStr[tempS1]
			_, ok2 := setStr[tempS1Reverse]
			if ok1 || ok2 {
				return
			}
			setStr[tempS1] = struct{}{}
			setStr[tempS1Reverse] = struct{}{}

			if tempS1 == tempS1Reverse {
				tempResult := lenS - index
				if result < tempResult {
					result = tempResult
				}
			} else {
				judgeResult(tempS1, tempS1Reverse, index+1, false)
				judgeResult(tempS1, tempS1Reverse, index+1, true)
			}
		} else {
			tempS2 := str[1:]
			tempS2Reverse := strReverse[:lenS-index]
			_, ok1 := setStr[tempS2]
			_, ok2 := setStr[tempS2Reverse]
			if ok1 || ok2 {
				return
			}
			setStr[tempS2] = struct{}{}
			setStr[tempS2Reverse] = struct{}{}

			if tempS2 == tempS2Reverse {
				tempResult := lenS - index
				if result < tempResult {
					result = tempResult
				}
			} else {
				judgeResult(tempS2, tempS2Reverse, index+1, true)
				judgeResult(tempS2, tempS2Reverse, index+1, false)
			}
		}

	}

	if s == sReverse {
		result = lenS
	} else {
		judgeResult(s, sReverse, 1, false)
		judgeResult(s, sReverse, 1, true)
	}

	if result == -1 {
		fmt.Println(1)
	} else {
		fmt.Println(result)
	}

}
