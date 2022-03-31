package main

import "fmt"

func AEdgeCheckerMain() {
	var a, b int
	fmt.Scan(&a, &b)

	if b == 10 && a == 1 {
		fmt.Println("Yes")
	} else if (b - a) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
