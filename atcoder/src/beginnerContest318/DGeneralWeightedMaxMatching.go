package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DGeneralWeightedMaxMatchingMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	dListList := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var dList []int
		dListStr := strings.Split(DGeneralWeightedMaxMatching(rdr), " ")
		for _, v := range dListStr {
			d, _ := strconv.Atoi(v)
			dList = append(dList, d)
		}
		dListList[i] = dList
	}

	var result int64
	result = -1
	maxLength := n
	if n%2 != 0 {
		maxLength = n - 1
	}

	var funcCalc func(nextIndex int, tempSum int64, alreadySet map[int]struct{}) int
	funcCalc = func(nextIndex int, tempSum int64, alreadySet map[int]struct{}) int {
		dList := dListList[nextIndex]
		for i, dValue := range dList {
			counterIndex := nextIndex + i + 1
			_, ok := alreadySet[counterIndex]
			if !ok {
				tempSum2 := tempSum + int64(dValue)
				newAlreadySet := map[int]struct{}{}
				for k, _ := range alreadySet {
					newAlreadySet[k] = struct{}{}
				}
				newAlreadySet[nextIndex] = struct{}{}
				newAlreadySet[counterIndex] = struct{}{}

				if maxLength <= len(newAlreadySet) {
					if result < tempSum2 {
						result = tempSum2
						return 0
					}
				} else {
					for j := 0; j < n-1; j++ {
						_, ok2 := newAlreadySet[j]
						if !ok2 {
							funcCalc(j, tempSum2, newAlreadySet)
						}
					}
				}
			}
		}
		return 0
	}
	for i := 0; i < n-1; i++ {
		funcCalc(i, int64(0), map[int]struct{}{})
	}
	fmt.Println(result)
}

func DGeneralWeightedMaxMatching(rdr *bufio.Reader) string {
	buf := make([]byte, 0, 10000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
