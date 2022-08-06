package main

import (
	"fmt"
	"strconv"
	"strings"
)

var CMonotonicallyIncreasingN, CMonotonicallyIncreasingM int
var CMonotonicallyIncreasingResultSlice []string

func CMonotonicallyIncreasingMain() {

	fmt.Scan(&CMonotonicallyIncreasingN, &CMonotonicallyIncreasingM)

	if CMonotonicallyIncreasingN == 1 {
		for i := 1; i <= CMonotonicallyIncreasingM; i++ {
			CMonotonicallyIncreasingResultSlice = append(CMonotonicallyIncreasingResultSlice, strconv.FormatInt(int64(i), 10))
		}
	} else {
		CMonotonicallyIncreasing([]string{}, -1, 1)
	}
	fmt.Println(strings.Join(CMonotonicallyIncreasingResultSlice, "\n"))

}

func CMonotonicallyIncreasing(inputArray []string, lastValue int, nowIndex int) {
	if nowIndex == 1 {
		for i := 1; i <= (CMonotonicallyIncreasingM - CMonotonicallyIncreasingN + 1); i++ {
			tempArray := append(inputArray, strconv.FormatInt(int64(i), 10))
			CMonotonicallyIncreasing(tempArray, i, 2)
		}
	} else {
		nextFirst := lastValue + 1
		if nowIndex == CMonotonicallyIncreasingN {
			for i := nextFirst; i <= CMonotonicallyIncreasingM; i++ {
				tempArray := append(inputArray, strconv.FormatInt(int64(i), 10))
				CMonotonicallyIncreasingResultSlice = append(CMonotonicallyIncreasingResultSlice, strings.Join(tempArray, " "))
			}
		} else {
			max := nextFirst + (CMonotonicallyIncreasingM - (CMonotonicallyIncreasingN - nowIndex) - nextFirst)
			for i := nextFirst; i <= max; i++ {
				tempArray := append(inputArray, strconv.FormatInt(int64(i), 10))
				CMonotonicallyIncreasing(tempArray, i, nowIndex+1)
			}
		}
	}
}
