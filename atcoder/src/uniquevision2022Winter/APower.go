package main

import (
	"fmt"
)

func APowerMain() {
	var a int64
	var b int
	fmt.Scan(&a, &b)

	var result int64
	result = 1
	for i := 0; i < b; i++ {
		result = result * a
	}
	fmt.Println(result)
}
