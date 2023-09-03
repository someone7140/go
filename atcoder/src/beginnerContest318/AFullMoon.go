package main

import (
	"fmt"
)

func AFullMoonMain() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)

	result := 0
	day := m

	for {
		if n >= day {
			result = result + 1
			day = day + p
		} else {
			break
		}
	}

	fmt.Println(result)
}
