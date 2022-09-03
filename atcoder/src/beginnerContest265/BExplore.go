package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BExploreMain() {
	var n, m, t int
	fmt.Scan(&n, &m, &t)

	var aArray = make([]int64, n)
	var bonusArray = make([]int64, n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := BExploreRdr(rdr)
	aStrArray := strings.Split(aLine, " ")
	for i, aMoji := range aStrArray {
		aNum, _ := strconv.Atoi(aMoji)
		aArray[i+1] = int64(aNum)
	}

	for i := 0; i < m; i++ {
		mLine := BExploreRdr(rdr)
		mStrArray := strings.Split(mLine, " ")

		x, _ := strconv.Atoi(mStrArray[0])
		y, _ := strconv.Atoi(mStrArray[1])
		bonusArray[x-1] = int64(y)
	}

	result := "Yes"
	var nokoriT int64
	nokoriT = int64(t)

	for i := 1; i < n; i++ {
		a := aArray[i]
		if a < nokoriT {
			nokoriT = nokoriT - a + bonusArray[i]
		} else {
			result = "No"
			break
		}
	}
	fmt.Println(result)
}

func BExploreRdr(rdr *bufio.Reader) string {
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
