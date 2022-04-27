package main

import (
	"fmt"
)

func AJoggingMain() {
	var a, b, c, d, e, f, x int
	fmt.Scan(&a, &b, &c, &d, &e, &f, &x)

	axSyou := x / (a + c)
	axAmari := x % (a + c)

	aKyori := float64(axSyou * b * a)
	if axAmari >= a {
		aKyori = aKyori + float64(b*a)
	} else {
		aKyori = aKyori + float64(axAmari)*float64(b)
	}

	dxSyou := x / (d + f)
	dxAmari := x % (d + f)

	dKyori := float64(dxSyou * e * d)
	if dxAmari >= d {
		dKyori = dKyori + float64(e*d)
	} else {
		dKyori = dKyori + float64(dxAmari)*float64(e)
	}

	if aKyori == dKyori {
		fmt.Println("Draw")
	} else if aKyori > dKyori {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}

}
