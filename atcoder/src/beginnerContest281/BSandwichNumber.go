package main

import (
	"fmt"
	"strconv"
)

func BSandwichNumberMain() {
	var s string
	fmt.Scan(&s)

	if len(s) != 8 {
		fmt.Println("No")
		return
	}

	upperA := rune('A')
	upperZ := rune('Z')
	numStr := ""
	for i, c := range s {
		if i == 0 {
			if c >= upperA && c <= upperZ {
				// 何もしない
			} else {
				fmt.Println("No")
				return
			}
		} else if i == (len(s) - 1) {
			if c >= upperA && c <= upperZ {
				// 何もしない
			} else {
				fmt.Println("No")
				return

			}
		} else {
			if c >= upperA && c <= upperZ {
				fmt.Println("No")
				return
			} else {
				numStr = numStr + string([]rune{c})
			}
		}
	}

	sNum, e := strconv.Atoi(numStr)
	if e != nil {
		fmt.Println("No")
	} else {
		if 100000 <= sNum && sNum <= 999999 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
