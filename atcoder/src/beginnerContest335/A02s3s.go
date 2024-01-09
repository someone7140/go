package main

import (
	"fmt"
)

func A02s3sMain() {
	var s string
	fmt.Scan(&s)

	lenS := len(s)
	fmt.Println(s[:lenS-1] + "4")
}
