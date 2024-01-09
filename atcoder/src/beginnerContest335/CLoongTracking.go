package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CLoongTrackingMain() {
	var n, q int
	fmt.Scan(&n, &q)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var resultSlice []string

	type CLoongTrackingXy struct {
		x int
		y int
	}
	var xyReverseSlice = make([]CLoongTrackingXy, n)

	for i := n; i >= 1; i-- {
		xyReverseSlice[n-i] = CLoongTrackingXy{
			x: i,
			y: 0,
		}
	}

	for i := 0; i < q; i++ {
		qStrs := strings.Split(CLoongTrackingReadLine(rdr), " ")
		if qStrs[0] == "1" {
			topXy := CLoongTrackingXy{
				x: xyReverseSlice[n-1].x, y: xyReverseSlice[n-1].y,
			}
			if qStrs[1] == "R" {
				topXy.x = topXy.x + 1
			}
			if qStrs[1] == "L" {
				topXy.x = topXy.x - 1
			}
			if qStrs[1] == "U" {
				topXy.y = topXy.y + 1
			}
			if qStrs[1] == "D" {
				topXy.y = topXy.y - 1
			}

			xyReverseSlice = append(xyReverseSlice[1:], topXy)
		} else {
			index, _ := strconv.Atoi(qStrs[1])
			index = n - index
			result := strconv.FormatInt(int64(xyReverseSlice[index].x), 10) + " " + strconv.FormatInt(int64(xyReverseSlice[index].y), 10)
			resultSlice = append(resultSlice, result)
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CLoongTrackingReadLine(rdr *bufio.Reader) string {
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
