package main

import (
	"fmt"
)

func AFiveIntegersMain() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	set := make(map[int]struct{})
	set[a] = struct{}{}
	set[b] = struct{}{}
	set[c] = struct{}{}
	set[d] = struct{}{}
	set[e] = struct{}{}

	fmt.Println(len(set))
}
