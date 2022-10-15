package main

import (
	"fmt"
)

func A484558Main() {
	var n int
	fmt.Scan(&n)

	result := fmt.Sprintf("%X", n)
	if len(result) == 1 {
		result = "0" + result
	}
	fmt.Println(result)
}
