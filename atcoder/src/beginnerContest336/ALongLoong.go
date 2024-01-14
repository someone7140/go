package main

import (
	"fmt"
)

func ALongLoongMain() {
	var n int
	fmt.Scan(&n)

	result := "L"
	for i := 0; i < n; i++ {
		result = result + "o"
	}
	result = result + "ng"
	fmt.Println(result)
}
