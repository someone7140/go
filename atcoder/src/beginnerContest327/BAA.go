package main

import (
	"fmt"
)

func BAAMain() {
	var b int64
	fmt.Scan(&b)

	bMap := map[int64]int{}
	bMap[1] = 1
	bMap[4] = 2
	bMap[27] = 3
	bMap[256] = 4
	bMap[3125] = 5
	bMap[46656] = 6
	bMap[823543] = 7
	bMap[16777216] = 8
	bMap[387420489] = 9
	bMap[10000000000] = 10
	bMap[285311670611] = 11
	bMap[8916100448256] = 12
	bMap[302875106592253] = 13
	bMap[11112006825558016] = 14
	bMap[437893890380859375] = 15

	result, ok := bMap[b]
	if !ok {
		result = result - 1
	}
	fmt.Println(result)

}
