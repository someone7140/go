package main

import (
	"fmt"
)

func AIntersectionMain() {
	var l1, r1, l2, r2 int
	fmt.Scan(&l1, &r1, &l2, &r2)

	if r1 > l2 {
		result := 0
		if l1 > l2 && r1 < r2 {
			result = r1 - l1
		} else if l1 > l2 && r1 >= r2 {
			result = r2 - l1
		} else if l1 <= l2 && r1 < r2 {
			result = r1 - l2
		} else {
			result = r2 - l2
		}

		if result < 0 {
			fmt.Println(0)
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println(0)
	}
}
