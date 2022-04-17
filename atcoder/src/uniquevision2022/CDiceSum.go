package main

import (
	"fmt"
)

func CDiceSumMain() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	nMap := map[int]int64{}

	var result int64
	result = 0

	for i := 0; i < n; i++ {
		if i == 0 {
			for j := 1; j <= m; j++ {
				if j <= k {
					nMap[j] = 1
				}
			}
		} else {
			nMapNext := map[int]int64{}
			for key, value := range nMap {
				for j := 1; j <= m; j++ {
					updateKey := key + j
					if updateKey <= k {
						updateNext, ok := nMapNext[updateKey]
						if ok {
							nMapNext[updateKey] = (updateNext + value) % 998244353
						} else {
							nMapNext[updateKey] = value
						}

					}
				}
			}
			nMap = nMapNext
		}
	}

	for _, value := range nMap {
		result = (result + value) % 998244353
	}

	fmt.Println(result)

}
