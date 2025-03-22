package main

import (
	"fmt"
	"math"
	"strconv"
)

func DCubesMain() {
	var n int64
	fmt.Scan(&n)

	x := int64(math.Cbrt(float64(n)))
	x3 := x * x * x

	if x3 == n {
		fmt.Println(-1)
		return
	}

	y := int64(0)
	y3 := y * y * y
	xChangeCount := 0
	for {
		if xChangeCount >= 9000000 {
			break
		}
		sabun := (x3 - y3)
		if sabun == n {
			break
		} else {
			if sabun < n {
				x = x + 1
				x3 = x * x * x
				if y < 1 && x3 > n {
					y = int64(math.Cbrt(float64(x3 - n)))
					y3 = y * y * y
				}
				xChangeCount = xChangeCount + 1
			} else {
				y = y + 1
				y3 = y * y * y
			}
		}
	}
	if xChangeCount >= 9000000 {
		fmt.Println(-1)
	} else {
		fmt.Println(strconv.FormatInt(x, 10) + " " + strconv.FormatInt(y, 10))

	}
}
