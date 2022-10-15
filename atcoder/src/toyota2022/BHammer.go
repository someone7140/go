package main

import (
	"fmt"
	"math"
)

func BHammerMain() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	if (x < 0 && y > 0) || (x > 0 && y < 0) {
		fmt.Println(math.Abs(float64(x)))
	} else {
		if x < 0 && y < x {
			fmt.Println(math.Abs(float64(x)))
		} else if x > 0 && y > x {
			fmt.Println(math.Abs(float64(x)))
		} else {
			if (y > 0 && y > z) || (y < 0 && y < z) {
				result := int((math.Abs(float64(z))))
				result = result + int((math.Abs(float64(x - z))))
				fmt.Println(result)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
