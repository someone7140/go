package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func CCandyTribulationMain() {
	var n, x, y int64
	fmt.Scan(&n, &x, &y)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CCandyTribulationRdr(rdr), " ")
	aArray := make([]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		a, _ := strconv.ParseInt(aStrArray[i], 10, 64)
		aArray[i] = a
	}
	slices.Sort(aArray)

	minIndexYCount := aArray[0]
	targetVal := y * aArray[0]
	var xCount int64
	var result int64 = -1
	for xCount = 0; i < aArray[0]; xCount++ {
		if result > 0 {
			break
		}
		if xCount > 0 {
			targetVal = targetVal - y + x
			minIndexYCount = minIndexYCount - 1
		}
		if minIndexYCount < 0 {
			break
		}

		tempResult := minIndexYCount
		minusFlag := false
		amariFlag := false
		for i = 1; i < n; i++ {
			count := aArray[i]
			tempVal := x * count
			sabun := targetVal - tempVal
			if sabun < 0 {
				minusFlag = true
				break
			}

			yCountSyou := sabun / (y - x)
			amari := sabun % (y - x)
			if amari > 0 {
				amariFlag = true
				break
			}
			tempResult = tempResult + yCountSyou
		}
		if minusFlag {
			break
		}
		if !amariFlag {
			result = tempResult
		}
	}

	fmt.Println(result)
}

func CCandyTribulationRdr(rdr *bufio.Reader) string {
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
