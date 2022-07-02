package main

import (
	"fmt"
)

func AWhenMain() {
	var k int
	fmt.Scan(&k)

	hour := 21
	minute := k

	if k >= 60 {
		hour = 22
		minute = k - 60
	}

	result := fmt.Sprintf("%02d:%02d", hour, minute)
	fmt.Println(result)
}
