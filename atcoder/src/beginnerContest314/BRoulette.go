package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BRouletteCount struct {
	index int
	count int
}

func BRouletteMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aArrayArray := make([][]int, n)
	for i := 0; i < n; i++ {
		BRouletteReadLine(rdr)
		var aArray []int
		for _, aMoji := range strings.Split(BRouletteReadLine(rdr), " ") {
			a, _ := strconv.Atoi(aMoji)
			aArray = append(aArray, a)
		}
		aArrayArray[i] = aArray
	}
	x, _ := strconv.Atoi(BRouletteReadLine(rdr))
	minCount := 100000000000
	var xTargetArray []BRouletteCount

	for i := 0; i < n; i++ {
		aArray := aArrayArray[i]
		findFlag := false
		for _, a := range aArray {
			if a == x {
				findFlag = true
				break
			}
		}
		if findFlag {
			tempCount := len(aArray)
			xTargetArray = append(xTargetArray, BRouletteCount{
				index: i,
				count: tempCount,
			})
			if tempCount < minCount {
				minCount = tempCount
			}
		}
	}

	if len(xTargetArray) == 0 {
		fmt.Println(0)
	} else {
		resultCount := 0
		var resultArray []string
		for _, xTarget := range xTargetArray {
			if xTarget.count == minCount {
				resultCount = resultCount + 1
				resultArray = append(resultArray, strconv.FormatInt(int64(xTarget.index+1), 10))
			}
		}
		fmt.Println(resultCount)
		fmt.Println(strings.Join(resultArray, " "))
	}
}

func BRouletteReadLine(rdr *bufio.Reader) string {
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
