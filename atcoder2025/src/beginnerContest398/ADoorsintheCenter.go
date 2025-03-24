package main

import (
	"fmt"
)

func ADoorsintheCenterMain() {
	var n int
	fmt.Scan(&n)

	result := ""

	if n%2 == 0 {
		for i := 0; i < n; i++ {
			if i == (n/2) || i == (n/2-1) {
				result = result + "="
			} else {
				result = result + "-"
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if i == (n / 2) {
				result = result + "="
			} else {
				result = result + "-"
			}
		}
	}

	fmt.Println(result)
}
