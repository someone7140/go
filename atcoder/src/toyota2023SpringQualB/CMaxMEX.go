package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aArr := make([]int, n)
	aMap := map[int]int{}
	for i, aStr := range strings.Split(CMaxMEXReadLine(rdr), " ") {
		a, _ := strconv.Atoi(aStr)
		aArr[i] = a
		aMapV, ok := aMap[a]

		if ok {
			aMap[a] = aMapV + 1
		} else {
			aMap[a] = 1
		}
	}
	result := 0
	for i := 0; i < k; i++ {
		_, ok := aMap[i]
		if i == 0 {
			if !ok {
				break
			} else {
				result = i + 1
			}
		} else {
			if ok {
				result = i + 1
			} else {
				break
			}
		}
	}

	fmt.Println(result)
}

func CMaxMEXReadLine(rdr *bufio.Reader) string {
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
