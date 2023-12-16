package main

import (
	"fmt"
	"strconv"
)

func AThreeThreesMain() {
	var n int
	fmt.Scan(&n)
	result := ""

	for i := 0; i < n; i++ {
		result = result + strconv.FormatInt(int64(n), 10)
	}

	fmt.Println(result)

}
