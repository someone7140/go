package main

import (
	"fmt"
)

func AWrongAnswerMain() {
	var a, b int
	fmt.Scan(&a, &b)

	var result int
	for i := 0; i < 10; i++ {
		if i != (a + b) {
			result = i
			break
		}

	}
	fmt.Println(result)
}
