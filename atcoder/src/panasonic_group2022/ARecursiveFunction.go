package main

import (
	"fmt"
)

func ARecursiveFunctionMain() {
	var n int
	fmt.Scan(&n)

	var result = n
	if n == 0 {
		result = 1
	}

	for i := 2; i <= n-1; i++ {
		result = result * i
	}
	fmt.Println(result)

}
