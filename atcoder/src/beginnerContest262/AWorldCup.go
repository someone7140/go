package main

import (
	"fmt"
)

func AWorldCupMain() {
	var y int
	fmt.Scan(&y)

	amari := y % 4

	if amari == 2 {
		fmt.Println(y)
	} else if amari == 0 {
		fmt.Println(y + 2)
	} else if amari == 1 {
		fmt.Println(y + 1)
	} else {
		fmt.Println(y + 3)
	}

}
