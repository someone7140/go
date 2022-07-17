package main

import (
	"fmt"
	"math"
	"strconv"
)

func BCounterclockwiseRotationMain() {
	var a, b, d float64
	fmt.Scan(&a, &b, &d)

	if a == 0 && b == 0 {
		fmt.Println("0 0")
	} else if d == 360 {
		fmt.Println(strconv.FormatInt(int64(a), 10) + " " + strconv.FormatInt(int64(b), 10))
	} else {
		radian := d * math.Pi / 180
		cos := math.Cos(radian)
		sin := math.Sin(radian)

		aResult := a*cos - b*sin
		bResult := a*sin + b*cos
		fmt.Println(strconv.FormatFloat(aResult, 'f', -1, 64) + " " + strconv.FormatFloat(bResult, 'f', -1, 64))
	}

}
