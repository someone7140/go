package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	countMap := map[int]int{}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			val := i * j
			sum = sum + val
			count, ok := countMap[val]
			if ok {
				countMap[val] = count + 1
			} else {
				countMap[val] = 1
			}
		}
	}

	findCount, ok := countMap[n]
	if ok {
		fmt.Println(sum - n*findCount)
	} else {
		fmt.Println(sum)
	}

}
