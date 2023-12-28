package main

import (
	"fmt"
	"math"
)

func BChristmasTreesMain() {
	var a, m, l, r int64
	fmt.Scan(&a, &m, &l, &r)

	absM := int64(math.Abs(float64(m)))

	if l == a && r == a {
		fmt.Println(1)
	} else if l >= a {
		result := (r - a) / absM
		minus := (l - 1 - a) / absM

		if minus > 0 {
			result = result - minus
		}
		if l == a {
			result = result + 1
		}
		fmt.Println(result)
	} else if r <= a {
		result := (a - l) / absM
		minus := (a - (r + 1)) / absM

		if minus > 0 {
			result = result - minus
		}
		if r == a {
			result = result + 1
		}
		fmt.Println(result)
	} else {
		result1 := (r - a) / absM
		result2 := (a - l) / absM
		fmt.Println(result1 + result2 + 1)
	}
}
