package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CRotationMain() {
	var n, q int
	fmt.Scan(&n, &q)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := CRotationRdr(rdr)
	sSlice := make([]string, n)
	for i, c := range s {
		sMoji := string([]rune{c})
		sSlice[i] = sMoji
	}

	var resultSlice []string
	xSum := 0
	for i := 0; i < q; i++ {
		query := CRotationRdr(rdr)
		querySlice := strings.Split(query, " ")
		t, _ := strconv.Atoi(querySlice[0])
		x, _ := strconv.Atoi(querySlice[1])
		if t == 2 {
			xSumTemp := x - xSum
			if xSumTemp < 0 {
				xSumTemp = n + xSumTemp
			} else if xSumTemp == 0 {
				xSumTemp = 1
			}

			resultSlice = append(resultSlice, sSlice[xSumTemp-1])
		}
		if t == 1 {
			xSumTemp := xSum + x
			if xSumTemp >= n {
				xSum = xSumTemp - n
			} else {
				xSum = xSumTemp
			}
		}
	}

	for _, v := range resultSlice {
		fmt.Println(v)
	}

}

func CRotationRdr(rdr *bufio.Reader) string {
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
