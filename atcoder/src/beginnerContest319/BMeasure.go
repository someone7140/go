package main

import (
	"fmt"
	"strconv"
)

func BMeasureMain() {
	var n int
	fmt.Scan(&n)

	var yakusuuList []int
	for i := 1; i < 10; i++ {
		if n%i == 0 {
			yakusuuList = append(yakusuuList, i)
		}
	}

	result := ""
	for i := 0; i <= n; i++ {
		tempResult := -1
		for _, yakusuu := range yakusuuList {
			if i%(n/yakusuu) == 0 {
				tempResult = yakusuu
				break
			}
		}

		if tempResult < 1 {
			result = result + "-"
		} else {
			result = result + strconv.FormatInt(int64(tempResult), 10)
		}
	}

	fmt.Println(result)
}
