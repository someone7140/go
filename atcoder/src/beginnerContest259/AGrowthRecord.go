package main

import (
	"fmt"
)

func AGrowthRecordMain() {
	var n, m, x, t, d int
	fmt.Scan(&n, &m, &x, &t, &d)

	if m <= n && m >= x {
		fmt.Println(t)
	} else {
		if m < x {
			fmt.Println(t - (x-m)*d)
		} else {
			fmt.Println(t + (n-x)*d)
		}
	}

}
