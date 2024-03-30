package main

import (
	"fmt"
)

func BPianoMain() {
	var w, b int
	fmt.Scan(&w, &b)

	baseStr := "wbwbwwbwbwbw"
	str := "wbwbwwbwbwbw"
	lenStr := 12

	for {
		if lenStr > (w+b)*2+30 {
			break
		} else {
			str = str + baseStr
			lenStr = lenStr + 12
		}
	}

	result := "No"
	tempW := 0
	tempB := 0
	// まずは頭から
	lenNow := 0
	for _, c := range str {
		sMoji := string([]rune{c})
		if sMoji == "w" {
			if tempW == w {
				break
			} else {
				tempW = tempW + 1
				lenNow = lenNow + 1
			}
		} else {
			if tempB == b {
				break
			} else {
				tempB = tempB + 1
				lenNow = lenNow + 1
			}
		}
	}

	if tempW == w && tempB == b {
		result = "Yes"
	} else {
		// 一周1個ずつずらしていく
		for i := 0; i < (w+b)+5; i++ {
			delete := string(str[i])
			if delete == "w" {
				tempW = tempW - 1
			} else {
				tempB = tempB - 1
			}
			addCount := 0
			for {
				add := string(str[lenNow+i+addCount])

				if add == "w" {
					if tempW == w {
						break
					} else {
						tempW = tempW + 1
						addCount = addCount + 1
					}

				} else {
					if tempB == b {
						break
					} else {
						tempB = tempB + 1
						addCount = addCount + 1
					}
				}
			}

			lenNow = lenNow - 1 + addCount
			if tempW == w && tempB == b {
				result = "Yes"
				break
			}
		}
	}

	fmt.Println(result)
}
