package main

import (
	"fmt"
	"math"
)

func BPXorYMain() {
	var x, y int
	fmt.Scan(&x, &y)

	var count float64 = 0

	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			if i+j >= x {
				count = count + 1
				continue
			}
			if math.Abs(float64(i-j)) >= float64(y) {
				count = count + 1
				continue
			}
		}
	}
	fmt.Println(count / float64(36))
}
