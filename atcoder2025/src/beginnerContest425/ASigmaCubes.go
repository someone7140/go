package main

import (
	"fmt"
)

func ASigmaCubesMain() {
	var n int
	fmt.Scan(&n)

	result := 0
	for i := 1; i <= n; i++ {
		temp1 := -1
		temp2 := i * i * i
		for j := 1; j < i; j++ {
			temp1 = temp1 * -1
		}
		result = result + (temp1 * temp2)
	}

	fmt.Println(result)
}
