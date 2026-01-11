package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BMagicSquareMain() {
	var n int
	fmt.Scan(&n)

	resultArrayArray := make([][]string, n)
	for i := 0; i < n; i++ {
		resultArrayArray[i] = make([]string, n)
	}

	nowColumn := 0
	nowRow := (n - 1) / 2
	resultArrayArray[nowColumn][nowRow] = "1"
	for i := 2; i <= n*n; i++ {
		tempNowColumn := (nowColumn - 1) % n
		if tempNowColumn < 0 {
			tempNowColumn = tempNowColumn + n
		}

		tempNowRow := (nowRow + 1) % n

		if resultArrayArray[tempNowColumn][tempNowRow] == "" {
			resultArrayArray[tempNowColumn][tempNowRow] = strconv.FormatInt(int64(i), 10)
			nowColumn = tempNowColumn
			nowRow = tempNowRow
		} else {
			tempNowColumn = (nowColumn + 1) % n
			resultArrayArray[tempNowColumn][nowRow] = strconv.FormatInt(int64(i), 10)
			nowColumn = tempNowColumn
		}
	}

	resultArray := make([]string, n)
	for i := 0; i < n; i++ {
		resultArray[i] = strings.Join(resultArrayArray[i], " ")
	}
	resultStr := strings.Join(resultArray, "\n")
	fmt.Println(resultStr)
}
