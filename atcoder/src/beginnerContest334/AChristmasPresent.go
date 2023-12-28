package main

import (
	"fmt"
)

func AChristmasPresentMain() {
	var b, g int
	fmt.Scan(&b, &g)

	if b > g {
		fmt.Println("Bat")
	} else {
		fmt.Println("Glove")
	}

}
