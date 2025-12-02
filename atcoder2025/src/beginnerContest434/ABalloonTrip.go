package main

import (
	"fmt"
)

func ABalloonTripMain() {
	var w, b int
	fmt.Scan(&w, &b)

	fmt.Println(w*1000/b + 1)
}
