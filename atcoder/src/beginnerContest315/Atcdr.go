package main

import (
	"fmt"
	"strings"
)

func AtcdrMain() {
	var s string
	fmt.Scan(&s)

	result := strings.Replace(s, "a", "", -1)
	result = strings.Replace(result, "e", "", -1)
	result = strings.Replace(result, "i", "", -1)
	result = strings.Replace(result, "o", "", -1)
	result = strings.Replace(result, "u", "", -1)

	fmt.Println(result)
}
