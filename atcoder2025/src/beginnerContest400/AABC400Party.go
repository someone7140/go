package main

import (
	"fmt"
)

func AABC400PartyMain() {
	var a int
	fmt.Scan(&a)

	if 400%a == 0 {
		fmt.Println(400 / a)
	} else {
		fmt.Println(-1)
	}
}
