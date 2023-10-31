package main

import (
	"fmt"
	"strconv"
	"strings"
)

func B326likeNumbersMain() {
	var n int
	fmt.Scan(&n)

	result := n
	for {
		resultStr := strconv.FormatInt(int64(result), 10)
		resultStrs := strings.Split(resultStr, "")

		hyaku, _ := strconv.Atoi(resultStrs[0])
		juu, _ := strconv.Atoi(resultStrs[1])
		ichi, _ := strconv.Atoi(resultStrs[2])

		if hyaku*juu == ichi {
			break
		} else {
			result = result + 1
		}
	}
	fmt.Println(result)
}
