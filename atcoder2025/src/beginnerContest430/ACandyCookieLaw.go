package main

import (
	"fmt"
)

func ACandyCookieLawMain() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if c < a {
		fmt.Println("No")
	} else {
		if d < b {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
