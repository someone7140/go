package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CBlueSpringMain() {
	var n, d, p int
	fmt.Scan(&n, &d, &p)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	fList := make([]int, n)
	fStrList := strings.Split(CBlueSpringReadLine(rdr), " ")

	var normalSum int64
	for i, v := range fStrList {
		f, _ := strconv.Atoi(v)
		fList[i] = f
		normalSum = normalSum + int64(f)
	}

	result := normalSum
	sort.Sort(sort.Reverse(sort.IntSlice(fList)))
	bubunwaList := make([]int64, n)
	bubunwaList[0] = normalSum
	for i := 1; i < n; i++ {
		bubunwaList[i] = bubunwaList[i-1] - int64(fList[i-1])
	}

	maisuu := 0
	for {
		var tempResult int64
		maisuu = maisuu + d
		if maisuu > n-1 {
			tempResult = int64(maisuu/d) * int64(p)
		} else {
			tempResult = int64(maisuu/d)*int64(p) + bubunwaList[maisuu]
		}
		if tempResult >= result {
			break
		} else {
			result = tempResult
		}
	}
	fmt.Println(result)
}

func CBlueSpringReadLine(rdr *bufio.Reader) string {
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
