package main

import (
	"fmt"
	"math"
	"strconv"
)

func DFreefallMain() {
	var DFreefallA, DFreefallB int64
	var DFreefallMin float64
	fmt.Scan(&DFreefallA, &DFreefallB)
	DFreefallMin = float64(DFreefallA)

	start := int64(0)
	end := DFreefallA
	half := DFreefallA / 2
	startFlag := true

	for {
		if start >= half || end <= half {
			break
		}

		halfResult := float64(half*DFreefallB) + float64(DFreefallA)/math.Sqrt(float64(half+1))

		if halfResult < 0 {
			end = half
			half = (start + end) / 2
		} else if halfResult < DFreefallMin {
			DFreefallMin = halfResult
			if startFlag {
				end = half
				startFlag = false
				half = (start + end) / 2
			} else {
				half1 := (half + end) / 2
				half2 := (start + half) / 2
				halfResult1 := float64(half*DFreefallB) + float64(DFreefallA)/math.Sqrt(float64(half+1))
				halfResult2 := float64(half*DFreefallB) + float64(DFreefallA)/math.Sqrt(float64(half+1))
				if halfResult1 < halfResult2 {
					half = half1
				} else {
					half = half2
				}
			}
		} else {
			end = half
			half = (start + end) / 2
		}
	}
	fmt.Println(strconv.FormatFloat(DFreefallMin, 'f', -1, 64))
}
