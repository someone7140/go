package main

import (
	"fmt"
	"strings"
)

func BCTZMain() {
	var n int
	fmt.Scan(&n)

	nishin := fmt.Sprintf("%b", n)
	nishinSlice := strings.Split(nishin, "")
	nishinLen := len(nishin)

	result := 0
	for i := 1; i <= nishinLen; i++ {
		if nishinSlice[nishinLen-i] == "0" {
			result = result + 1
		} else {
			break
		}
	}
	fmt.Println(result)
}
