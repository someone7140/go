package main

import (
	"fmt"
	"math"
	"strconv"
)

func BBrokenRoundingMain() {
	var x, k int64
	fmt.Scan(&x, &k)

	jousuu := int64(1)
	result := float64(x)
	for i := int64(1); i <= k; i++ {
		jousuu = jousuu * 10
		jousuuF := float64(jousuu)
		result = math.Round(result/jousuuF) * jousuuF
	}
	fmt.Println(strconv.FormatFloat(result, 'f', -1, 64))
}
