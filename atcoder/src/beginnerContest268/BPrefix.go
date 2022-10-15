package main

import (
	"fmt"
	"strings"
)

func BPrefixMain() {
	var s string
	fmt.Scan(&s)
	var t string
	fmt.Scan(&t)

	if strings.HasPrefix(t, s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
