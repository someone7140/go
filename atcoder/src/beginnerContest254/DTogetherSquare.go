package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int64
	fmt.Scan(&n)

	var result int64 = 0
	var i int64
	for i = 1; i <= n; i++ {
		jousuu := i * i
		yakusuuList := DTogetherSquaregetYakusuu(jousuu, n)
		yakusuuLen := len(yakusuuList)
		j := yakusuuLen/2 + 1

		result = result + 1

		for k := j; k < yakusuuLen; k++ {
			if yakusuuList[k] > n {
				break
			} else {
				result = result + 2
			}
		}
	}

	fmt.Println(result)
}

func DTogetherSquaregetYakusuu(n int64, maxLimit int64) []int64 {
	var results []int64
	var i int64
	for i = 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}

		b := n / i
		if i <= maxLimit && b <= maxLimit {
			if i == b {
				results = append(results, i)
			} else {
				results = append(results, i)
				results = append(results, b)
			}
		}

	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	return results
}
