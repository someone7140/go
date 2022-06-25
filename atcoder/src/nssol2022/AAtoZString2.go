package main

import (
	"fmt"
)

func AAtoZString2Main() {
	var n, x int
	fmt.Scan(&n, &x)

	amari := x % n
	ascii := 0
	if amari > 0 {
		ascii = x/n + 65
	} else {
		ascii = x/n + 64
	}

	fmt.Println(string(ascii))
}
