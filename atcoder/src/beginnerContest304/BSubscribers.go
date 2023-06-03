package main

import (
	"fmt"
)

func BSubscribersMain() {
	var n int
	fmt.Scan(&n)

	if n <= 999 {
		fmt.Println(n)
	} else if n >= 1000 && n < 10000 {
		fmt.Println(n / 10 * 10)
	} else if n >= 10000 && n < 100000 {
		fmt.Println(n / 100 * 100)
	} else if n >= 100000 && n < 1000000 {
		fmt.Println(n / 1000 * 1000)
	} else if n >= 1000000 && n < 10000000 {
		fmt.Println(n / 10000 * 10000)
	} else if n >= 10000000 && n < 100000000 {
		fmt.Println(n / 100000 * 100000)
	} else if n >= 100000000 && n < 1000000000 {
		fmt.Println(n / 1000000 * 1000000)
	} else {
		fmt.Println(n)
	}
}
