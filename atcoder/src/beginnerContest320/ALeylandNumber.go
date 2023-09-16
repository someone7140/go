package main

import (
	"fmt"
	"math"
)

func ALeylandNumberMain() {
	var a, b float64
	fmt.Scan(&a, &b)

	result := int64(math.Pow(a, b) + math.Pow(b, a))
	fmt.Println(result)
}
