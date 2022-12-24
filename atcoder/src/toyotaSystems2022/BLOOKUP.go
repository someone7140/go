package main

import (
	"fmt"
	"strings"
)

func BLOOKUPMain() {
	var s string
	fmt.Scan(&s)
	var t string
	fmt.Scan(&t)

	resultInt := strings.Index(s, t)
	if resultInt < 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
