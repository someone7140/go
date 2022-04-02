package main

import (
	"fmt"
	"math"
	"strconv"
)

func BGetCloserMain() {
	var a, b float64
	fmt.Scan(&a, &b)

	katamuki := b / a

	x2Jou := 1 / (1 + katamuki*katamuki)
	y2Jou := 1 - x2Jou

	fmt.Println(
		strconv.FormatFloat(math.Sqrt(x2Jou), 'f', -1, 64) + " " + strconv.FormatFloat(math.Sqrt(y2Jou), 'f', -1, 64),
	)
}
