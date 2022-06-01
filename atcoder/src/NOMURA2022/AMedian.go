package main

import (
	"fmt"
)

func AMedianMain() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if b >= a && b <= c {
		fmt.Println("Yes")
	} else if b >= c && b <= a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
