package main

import (
	"fmt"
)

func AJoggingMain() {
	var s string
	fmt.Scan(&s)

	if len(s) == 3 {
		fmt.Println(s + s)
	} else if len(s) == 2 {
		fmt.Println(s + s + s)
	} else {
		fmt.Println(s + s + s + s + s + s)
	}

}
