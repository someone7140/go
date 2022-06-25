package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CKSwapMain() {
	var n, k int
	fmt.Scan(&n, &k)

	result := "Yes"
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := CKSwapRdr(rdr)
	aStrArray := strings.Split(aLine, " ")

	var aArray = make([]int, n)

	for i := 0; i < n; i++ {
		aArray[i], _ = strconv.Atoi(aStrArray[i])
	}
	sortedArray := make([]int, n)
	copy(sortedArray, aArray)
	sort.Ints(sortedArray)

	var confirmArray = make([]int, n)

	for i := 0; i < n; i++ {
		if confirmArray[i] == 0 {
			var aTempArray []int
			var indexArray []int
			j := i
			for {
				indexArray = append(indexArray, j)
				aTempArray = append(aTempArray, aArray[j])
				j = j + k
				if j > n-1 {
					break
				}
			}
			sort.Ints(aTempArray)
			for l := 0; l < len(aTempArray); l++ {
				if aTempArray[l] != sortedArray[indexArray[l]] {
					result = "No"
					break
				} else {
					confirmArray[indexArray[l]] = 1
				}
			}
			if result == "No" {
				break
			}
		}
	}

	fmt.Println(result)
}

func CKSwapRdr(rdr *bufio.Reader) string {
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
