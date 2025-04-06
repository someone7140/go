package main

import (
	"fmt"
	"strconv"
)

func BSumofGeometricSeriesMain() {
	var n, m int
	fmt.Scan(&n, &m)

	max := int64(1000000000)
	total := int64(1)
	temp := int64(1)

	result := ""

	for i := 0; i < m; i++ {
		temp = temp * int64(n)
		if temp > max {
			result = "inf"
			break
		}
		total = total + temp
		if total > max {
			result = "inf"
			break
		}
	}

	if result == "" {
		result = strconv.FormatInt(int64(total), 10)
	}
	fmt.Println(result)
}
