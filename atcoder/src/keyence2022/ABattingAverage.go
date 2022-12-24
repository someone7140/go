package main

import (
	"fmt"
)

func ABattingAverageMain() {
	var a, b float64
	fmt.Scan(&a, &b)

	fmt.Printf("%.3f", b/a)
}
