package main

import (
	"fmt"
)

func CJumpingTakahashiMain() {
	var n, x int
	fmt.Scan(&n, &x)

	var abSlice = make([]map[int]struct{}, n)
	result := "No"

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		set := make(map[int]struct{})

		if i == 0 {
			if n == 1 {
				if a == x || b == x {
					result = "Yes"
				}
			} else {
				set[a] = struct{}{}
				set[b] = struct{}{}
				abSlice[i] = set
			}

		} else if i == (n - 1) {
			beforeSet := abSlice[i-1]
			for k, _ := range beforeSet {
				if (a+k) == x || (b+k) == x {
					result = "Yes"
				}
			}
		} else {
			beforeSet := abSlice[i-1]
			for k, _ := range beforeSet {
				set[a+k] = struct{}{}
				set[b+k] = struct{}{}
			}
			abSlice[i] = set
		}

	}
	fmt.Println(result)

}
