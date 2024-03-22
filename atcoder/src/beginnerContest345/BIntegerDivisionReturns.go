package main

import (
	"fmt"
)

func BIntegerDivisionReturnsMain() {
	var x int64
	fmt.Scan(&x)

	wari10 := x / 10
	amari10 := x % 10

	result := wari10
	if amari10 > 0 {
		result = result + 1
	}

	fmt.Println(result)
}
