package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CMinMaxPairMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := CMinMaxPairRdr(rdr)
	aArray := make([]int, n)
	aStrArray := strings.Split(aLine, " ")
	for i, aMoji := range aStrArray {
		pNum, _ := strconv.Atoi(aMoji)
		aArray[i] = pNum
	}

	result := 0
	aMap := map[int]int{}
	var aSameSlice []int

	for i := 1; i <= n; i++ {
		a := aArray[i-1]
		if a != i {
			v, ok := aMap[a]
			if ok {
				if v == i {
					result = result + 1
				}
			}
			aMap[i] = a
		} else {
			aSameSlice = append(aSameSlice, a)
		}
	}

	sameLen := len(aSameSlice)
	if sameLen > 1 {
		result = result + sameLen*(sameLen-1)/2
	}

	fmt.Println(result)

}

func CMinMaxPairRdr(rdr *bufio.Reader) string {
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
