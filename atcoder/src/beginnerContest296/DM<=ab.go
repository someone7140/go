package main

import (
	"fmt"
)

func main() {
	var n, m int64
	fmt.Scan(&n, &m)

	var result int64
	result = -1

	var i int64
	for i = 1; i <= n; i++ {
		i2 := i * i
		if m%i == 0 {
			yaku := m / i
			if yaku <= n {
				result = m
				break
			}
		} else {
			yaku := m/i + 1
			if yaku <= n {
				tempResult := yaku * i
				if result == -1 || tempResult < result {
					result = tempResult
				}
				if i2 > m+1 {
					break
				}
			}
		}
	}

	fmt.Println(result)
}
