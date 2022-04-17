package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ALackedNumberMain() {
	var s string
	fmt.Scan(&s)
	result := ""

	for i := 0; i < 10; i++ {
		iStr := strconv.FormatInt(int64(i), 10)
		find := strings.Index(s, iStr)
		if find == -1 {
			result = iStr
			break
		}
	}
	fmt.Println(result)
}
