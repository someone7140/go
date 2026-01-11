package main

import (
	"fmt"
)

func ATriangularNumberMain() {
	var n int
	fmt.Scan(&n)

	fmt.Println((n * (n + 1)) / 2)
}
