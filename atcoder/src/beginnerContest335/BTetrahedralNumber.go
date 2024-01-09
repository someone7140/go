package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BTetrahedralNumberMain() {
	var n int
	fmt.Scan(&n)

	var resultSlice []string

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= n; k++ {
				sum := i + j + k
				if sum <= n {
					result := strconv.FormatInt(int64(i), 10) + " " + strconv.FormatInt(int64(j), 10) + " " + strconv.FormatInt(int64(k), 10)
					resultSlice = append(resultSlice, result)
				}

			}
		}
	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}
