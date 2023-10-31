package main

import (
	"fmt"
)

func A2UP3DOWNMain() {
	var x, y int
	fmt.Scan(&x, &y)

	if y < x {
		if x-y <= 3 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else if x < y {
		if y-x <= 2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		fmt.Println("No")
	}
}
