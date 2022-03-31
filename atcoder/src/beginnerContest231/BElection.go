package main

import "fmt"

func BElectionMain() {
	var n int
	fmt.Scan(&n)

	sMap := map[string]int{}

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		v, ok := sMap[s]
		if ok {
			sMap[s] = 1 + v
		} else {
			sMap[s] = 1
		}

	}
	maxCount := -1
	result := ""
	for k, v := range sMap {
		if v > maxCount {
			maxCount = v
			result = k
		}
	}
	fmt.Println(result)
}
