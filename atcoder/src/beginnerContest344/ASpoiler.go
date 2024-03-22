package main

import (
	"fmt"
	"strings"
)

func ASpoilerMain() {
	var s string
	fmt.Scan(&s)

	sArray := strings.Split(s, "|")
	fmt.Println(sArray[0] + sArray[2])
}
