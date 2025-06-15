package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BReverseProxyMain() {
	var n, q int
	fmt.Scan(&n, &q)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	xStrList := strings.Split(BReverseProxyRdr(rdr), " ")
	boxCountMap := map[int]int{}
	for i := 0; i < n; i++ {
		boxCountMap[i] = 0
	}

	var resultSlice []string
	for _, xStr := range xStrList {
		x, _ := strconv.Atoi(xStr)
		if x == 0 {
			minCount := -1
			minBox := -1
			for k, v := range boxCountMap {
				if minCount == -1 {
					minCount = v
					minBox = k
				} else if minCount > v {
					minCount = v
					minBox = k
				} else if minCount == v {
					if minBox > k {
						minBox = k
					}
				}
			}
			boxCountMap[minBox] = boxCountMap[minBox] + 1
			resultSlice = append(resultSlice, strconv.FormatInt(int64(minBox+1), 10))
		} else {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(x), 10))
			x = x - 1
			boxCountMap[x] = boxCountMap[x] + 1
		}
	}

	fmt.Println(strings.Join(resultSlice, " "))
}

func BReverseProxyRdr(rdr *bufio.Reader) string {
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
