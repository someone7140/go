package main

import (
	"fmt"
)

func AAdjacentSquaresMain() {
	var h, w int
	fmt.Scan(&h, &w)
	var r, c int
	fmt.Scan(&r, &c)

	if h == 1 && w == 1 {
		fmt.Println(0)
	} else if (r == 1 || r == h) && (c == 1 || c == w) {
		if h == 1 || w == 1 {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	} else if r == 1 || r == h || c == 1 || c == w {
		if h == 1 || w == 1 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	} else {
		fmt.Println(4)
	}

}
