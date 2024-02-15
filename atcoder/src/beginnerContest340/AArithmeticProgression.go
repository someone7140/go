package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AArithmeticProgressionMain() {
	var a, b, d int
	fmt.Scan(&a, &b, &d)

	temp := a
	var resultStrArray []string
	resultStrArray = append(resultStrArray, strconv.FormatInt(int64(temp), 10))
	for {
		if temp >= b {
			break
		}
		temp = temp + d
		resultStrArray = append(resultStrArray, strconv.FormatInt(int64(temp), 10))
	}

	fmt.Println(strings.Join(resultStrArray, " "))

}
